// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { css } from '@emotion/css';
import cx from 'classnames';
import { PropsWithChildren } from 'react';

import { GrafanaTheme2, PluginExtensionLink, TraceKeyValuePair } from '@grafana/data';
import { Icon, useStyles2 } from '@grafana/ui';

import { autoColor } from '../../Theme';
import CopyIcon from '../../common/CopyIcon';
import TNil from '../../types/TNil';

import jsonMarkup from './jsonMarkup';

const copyIconClassName = 'copyIcon';

export const getStyles = (theme: GrafanaTheme2) => {
  return {
    KeyValueTable: css({
      label: 'KeyValueTable',
      background: autoColor(theme, '#fff'),
      border: `1px solid ${autoColor(theme, '#ddd')}`,
      marginBottom: '0.5rem',
      maxHeight: '450px',
      overflow: 'auto',
    }),
    table: css({
      width: '100%',
    }),
    body: css({
      label: 'body',
      verticalAlign: 'baseline',
    }),
    row: css({
      label: 'row',
      '& > td': {
        padding: '0 0.5rem',
        height: '30px',
      },
      '&:nth-child(2n) > td': {
        background: autoColor(theme, '#f5f5f5'),
      },
      [`&:not(:hover) .${copyIconClassName}`]: {
        visibility: 'hidden',
      },
      'a span': {
        color: `${theme.colors.text.link} !important`,
      },
      'a:hover span': {
        textDecoration: 'underline',
      },
    }),
    keyColumn: css({
      label: 'keyColumn',
      color: autoColor(theme, '#888'),
      whiteSpace: 'pre',
      width: '125px',
    }),
    copyColumn: css({
      label: 'copyColumn',
      textAlign: 'right',
    }),
    linkIcon: css({
      label: 'linkIcon',
      verticalAlign: 'middle',
      fontWeight: 'bold',
    }),
    jsonTable: css({
      display: 'inline-block',
    }),
  };
};

const jsonObjectOrArrayStartRegex = /^(\[|\{)/;

function parseIfComplexJson(value: unknown) {
  // if the value is a string representing actual json object or array, then use json-markup
  if (typeof value === 'string' && jsonObjectOrArrayStartRegex.test(value)) {
    // otherwise just return as is
    try {
      return JSON.parse(value);
      // eslint-disable-next-line no-empty
    } catch (_) {}
  }
  return value;
}

export type KeyValuesTableLink = Pick<PluginExtensionLink, 'path' | 'title' | 'onClick' | 'icon'>;

interface LinkValueProps {
  link: KeyValuesTableLink;
}

export const LinkValue = ({ link, children }: PropsWithChildren<LinkValueProps>) => {
  const { path, title = '', onClick, icon = 'external-link-alt' } = link;

  return (
    <a href={path} title={title} onClick={onClick} target="_blank" rel="noopener noreferrer">
      {children} <Icon name={icon} />
    </a>
  );
};

export type KeyValuesTableProps = {
  data: TraceKeyValuePair[];
  linksGetter?: ((pairs: TraceKeyValuePair[], index: number) => KeyValuesTableLink[]) | TNil;
  onlyValues?: boolean;
};

export default function KeyValuesTable(props: KeyValuesTableProps) {
  const { data, linksGetter, onlyValues } = props;
  const styles = useStyles2(getStyles);
  return (
    <div className={cx(styles.KeyValueTable)} data-testid="KeyValueTable">
      <table className={styles.table}>
        <tbody className={styles.body}>
          {data.map((row, i) => {
            let markup = { __html: '' };
            if (row.type === 'code') {
              markup = {
                __html: `<pre style="border: none; background: none">${row.value}</pre>`,
              };
            } else if (row.type === 'text') {
              markup = {
                __html: `<span style="white-space: pre-wrap;">${row.value}</span>`,
              };
            } else {
              markup = {
                __html: jsonMarkup(parseIfComplexJson(row.value)),
              };
            }

            const jsonTable = <div className={styles.jsonTable} dangerouslySetInnerHTML={markup} />;
            const links = linksGetter?.(data, i);
            let valueMarkup;
            if (links && links.length) {
              // TODO: handle multiple items
              valueMarkup = (
                <div>
                  <LinkValue link={links[0]}>{jsonTable}</LinkValue>
                </div>
              );
            } else {
              valueMarkup = jsonTable;
            }
            return (
              // `i` is necessary in the key because row.key can repeat
              <tr className={styles.row} key={`${row.key}-${i}`}>
                {!onlyValues && (
                  <td className={styles.keyColumn} data-testid="KeyValueTable--keyColumn">
                    {row.key}
                  </td>
                )}
                <td>{valueMarkup}</td>
                <td className={styles.copyColumn}>
                  <CopyIcon
                    className={copyIconClassName}
                    copyText={row.type === 'code' || row.type === 'text' ? row.value : JSON.stringify(row, null, 2)}
                    tooltipTitle="Copy"
                  />
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
