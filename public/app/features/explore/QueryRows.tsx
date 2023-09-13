import { createSelector } from '@reduxjs/toolkit';
import React, { useCallback, useMemo } from 'react';

import { CoreApp } from '@grafana/data';
import { reportInteraction } from '@grafana/runtime';
import { DataQuery } from '@grafana/schema';
import { getNextRefIdChar } from 'app/core/utils/query';
import { useDispatch, useSelector } from 'app/types';

import { getDatasourceSrv } from '../plugins/datasource_srv';
import { QueryEditorRows } from '../query/components/QueryEditorRows';

import { changeCorrelationEditorDetails } from './state/main';
import { changeQueries, runQueries } from './state/query';
import { getExploreItemSelector, selectCorrelationDetails, isLeftPaneSelector } from './state/selectors';

interface Props {
  exploreId: string;
}

const makeSelectors = (exploreId: string) => {
  const exploreItemSelector = getExploreItemSelector(exploreId);
  return {
    getQueries: createSelector(exploreItemSelector, (s) => s!.queries),
    getQueryResponse: createSelector(exploreItemSelector, (s) => s!.queryResponse),
    getHistory: createSelector(exploreItemSelector, (s) => s!.history),
    getEventBridge: createSelector(exploreItemSelector, (s) => s!.eventBridge),
    getDatasourceInstanceSettings: createSelector(
      exploreItemSelector,
      (s) => getDatasourceSrv().getInstanceSettings(s!.datasourceInstance?.uid)!
    ),
  };
};

export const QueryRows = ({ exploreId }: Props) => {
  const dispatch = useDispatch();
  const { getQueries, getDatasourceInstanceSettings, getQueryResponse, getHistory, getEventBridge } = useMemo(
    () => makeSelectors(exploreId),
    [exploreId]
  );

  const queries = useSelector(getQueries);
  const dsSettings = useSelector(getDatasourceInstanceSettings);
  const queryResponse = useSelector(getQueryResponse);
  const history = useSelector(getHistory);
  const eventBridge = useSelector(getEventBridge);
  const correlationDetails = useSelector(selectCorrelationDetails);
  const isCorrelationsEditorMode = correlationDetails?.editorMode || false;
  const isLeftPane = useSelector(isLeftPaneSelector(exploreId));

  const onRunQueries = useCallback(() => {
    dispatch(runQueries({ exploreId }));
  }, [dispatch, exploreId]);

  const onChange = useCallback(
    (newQueries: DataQuery[]) => {
      dispatch(changeQueries({ exploreId, queries: newQueries }));

      if (isCorrelationsEditorMode && !correlationDetails?.dirty && !isLeftPane) {
        dispatch(changeCorrelationEditorDetails({ dirty: true }));
      }
    },
    [correlationDetails?.dirty, dispatch, exploreId, isCorrelationsEditorMode, isLeftPane]
  );

  const onAddQuery = useCallback(
    (query: DataQuery) => {
      onChange([...queries, { ...query, refId: getNextRefIdChar(queries) }]);
    },
    [onChange, queries]
  );

  const onQueryCopied = () => {
    reportInteraction('grafana_explore_query_row_copy');
  };

  const onQueryRemoved = () => {
    reportInteraction('grafana_explore_query_row_remove');
  };

  const onQueryToggled = (queryStatus?: boolean) => {
    reportInteraction('grafana_query_row_toggle', queryStatus === undefined ? {} : { queryEnabled: queryStatus });
  };

  return (
    <QueryEditorRows
      dsSettings={dsSettings}
      queries={queries}
      onQueriesChange={onChange}
      onAddQuery={onAddQuery}
      onRunQueries={onRunQueries}
      onQueryCopied={onQueryCopied}
      onQueryRemoved={onQueryRemoved}
      onQueryToggled={onQueryToggled}
      data={queryResponse}
      app={CoreApp.Explore}
      history={history}
      eventBus={eventBridge}
    />
  );
};
