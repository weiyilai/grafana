import { locationUtil } from '@grafana/data';
import { t } from '@grafana/i18n';
import { Spec as DashboardV2Spec } from '@grafana/schema/dist/esm/schema/dashboard/v2';
import { Status } from '@grafana/schema/src/schema/dashboard/v2';
import { backendSrv } from 'app/core/services/backend_srv';
import { getMessageFromError, getStatusFromError } from 'app/core/utils/errors';
import kbn from 'app/core/utils/kbn';
import { ScopedResourceClient } from 'app/features/apiserver/client';
import {
  AnnoKeyFolder,
  AnnoKeyFolderTitle,
  AnnoKeyFolderUrl,
  AnnoKeyMessage,
  AnnoKeyGrantPermissions,
  DeprecatedInternalId,
  Resource,
  ResourceClient,
  ResourceForCreate,
} from 'app/features/apiserver/types';
import { getDashboardUrl } from 'app/features/dashboard-scene/utils/getDashboardUrl';
import { DeleteDashboardResponse } from 'app/features/manage-dashboards/types';
import { DashboardDTO, SaveDashboardResponseDTO } from 'app/types/dashboard';

import { SaveDashboardCommand } from '../components/SaveDashboard/types';

import { DashboardAPI, DashboardVersionError, DashboardWithAccessInfo, ListDeletedDashboardsOptions } from './types';
import { isDashboardV2Spec } from './utils';

export const K8S_V2_DASHBOARD_API_CONFIG = {
  group: 'dashboard.grafana.app',
  version: 'v2beta1',
  resource: 'dashboards',
};

export class K8sDashboardV2API
  implements DashboardAPI<DashboardWithAccessInfo<DashboardV2Spec> | DashboardDTO, DashboardV2Spec>
{
  private client: ResourceClient<DashboardV2Spec, Status>;

  constructor() {
    this.client = new ScopedResourceClient<DashboardV2Spec>(K8S_V2_DASHBOARD_API_CONFIG);
  }

  async getDashboardDTO(uid: string) {
    try {
      const dashboard = await this.client.subresource<DashboardWithAccessInfo<DashboardV2Spec>>(uid, 'dto');

      // FOR /dto calls returning v2 spec we are ignoring the conversion status to avoid runtime errors caused by the status
      // being saved for v2 resources that's been client-side converted to v2 and then PUT to the API server.
      if (
        !isDashboardV2Spec(dashboard.spec) &&
        dashboard.status?.conversion?.failed &&
        (dashboard.status.conversion.storedVersion === 'v1alpha1' ||
          dashboard.status.conversion.storedVersion === 'v1beta1' ||
          dashboard.status.conversion.storedVersion === 'v0alpha1')
      ) {
        throw new DashboardVersionError(dashboard.status.conversion.storedVersion, dashboard.status.conversion.error);
      }

      // load folder info if available
      if (dashboard.metadata.annotations && dashboard.metadata.annotations[AnnoKeyFolder]) {
        try {
          const folder = await backendSrv.getFolderByUid(dashboard.metadata.annotations[AnnoKeyFolder]);
          dashboard.metadata.annotations[AnnoKeyFolderTitle] = folder.title;
          dashboard.metadata.annotations[AnnoKeyFolderUrl] = folder.url;
        } catch (e) {
          throw new Error('Failed to load folder');
        }
      } else if (dashboard.metadata.annotations && !dashboard.metadata.annotations[AnnoKeyFolder]) {
        // Set AnnoKeyFolder to empty string for top-level dashboards
        // This ensures NestedFolderPicker correctly identifies it as being in the "Dashboard" root folder
        // AnnoKeyFolder undefined -> top-level dashboard -> empty string
        dashboard.metadata.annotations[AnnoKeyFolder] = '';
      }

      return dashboard;
    } catch (e) {
      const status = getStatusFromError(e);
      const message = getMessageFromError(e);
      // Hacking around a bug in k8s api server that returns 500 for not found resources
      if (message.includes('not found') && status !== 404) {
        // @ts-expect-error
        e.status = 404;
        // @ts-expect-error
        e.data.message = 'Dashboard not found';
      }

      throw e;
    }
  }

  deleteDashboard(uid: string, showSuccessAlert: boolean): Promise<DeleteDashboardResponse> {
    return this.client.delete(uid, showSuccessAlert).then((v) => ({
      id: 0,
      message: v.message,
      title: t('dashboard.k8s-dashboard-v2api.title.deleted', 'deleted'),
    }));
  }

  async saveDashboard(options: SaveDashboardCommand<DashboardV2Spec>): Promise<SaveDashboardResponseDTO> {
    const dashboard = options.dashboard;

    const obj: ResourceForCreate<DashboardV2Spec> = {
      // the metadata will have the name that's the uid
      metadata: {
        ...options?.k8s,
      },
      spec: {
        ...dashboard,
      },
    };

    // add annotations
    if (options.message) {
      obj.metadata.annotations = {
        ...obj.metadata.annotations,
        [AnnoKeyMessage]: options.message,
      };
    } else if (obj.metadata.annotations) {
      delete obj.metadata.annotations[AnnoKeyMessage];
    }

    // add folder annotation
    if (options.folderUid) {
      // remove frontend folder annotations
      delete obj.metadata.annotations?.[AnnoKeyFolderTitle];
      delete obj.metadata.annotations?.[AnnoKeyFolderUrl];

      obj.metadata.annotations = {
        ...obj.metadata.annotations,
        [AnnoKeyFolder]: options.folderUid,
      };
    }

    if (obj.metadata.name) {
      // remove resource version when updating
      delete obj.metadata.resourceVersion;
      return this.client.update(obj).then((v) => this.asSaveDashboardResponseDTO(v));
    }
    obj.metadata.annotations = {
      ...obj.metadata.annotations,
      [AnnoKeyGrantPermissions]: 'default',
    };
    return await this.client.create(obj).then((v) => this.asSaveDashboardResponseDTO(v));
  }

  asSaveDashboardResponseDTO(v: Resource<DashboardV2Spec>): SaveDashboardResponseDTO {
    const url = locationUtil.assureBaseUrl(
      getDashboardUrl({
        uid: v.metadata.name,
        currentQueryParams: '',
        slug: kbn.slugifyForUrl(v.spec.title.trim()),
      })
    );

    let dashId = 0;
    if (v.metadata.labels?.[DeprecatedInternalId]) {
      dashId = parseInt(v.metadata.labels[DeprecatedInternalId], 10);
    }

    return {
      uid: v.metadata.name,
      version: v.metadata.generation ?? 0,
      id: dashId,
      status: 'success',
      url,
      slug: '',
    };
  }

  listDeletedDashboards(options: ListDeletedDashboardsOptions) {
    return this.client.list({ ...options, labelSelector: 'grafana.app/get-trash=true' });
  }

  restoreDashboard(dashboard: Resource<DashboardV2Spec>) {
    // reset the resource version to create a new resource
    dashboard.metadata.resourceVersion = '';
    return this.client.create(dashboard);
  }
}
