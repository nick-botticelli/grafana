import { config } from '@grafana/runtime';
import { setDashboardAPI } from 'app/features/dashboard/api/dashboard_api';
import { getDashboardScenePageStateManager } from 'app/features/dashboard-scene/pages/DashboardScenePageStateManager';

import { enterEditMode, updateMyVar, updateScopes, updateTimeRange } from './utils/actions';
import { getDatasource, getInstanceSettings } from './utils/mocks';
import { renderDashboard, resetScenes } from './utils/render';

jest.mock('@grafana/runtime', () => ({
  __esModule: true,
  ...jest.requireActual('@grafana/runtime'),
  useChromeHeaderHeight: jest.fn(),
  getDataSourceSrv: () => ({ get: getDatasource, getInstanceSettings }),
  usePluginLinks: jest.fn().mockReturnValue({ links: [] }),
}));

describe('Dashboard reload', () => {
  beforeAll(() => {
    config.featureToggles.scopeFilters = true;
    config.featureToggles.groupByVariable = true;
  });

  it.each([
    [false, false, false, false],
    [false, false, true, false],
    [false, true, false, false],
    [false, true, true, false],
    [true, false, false, false],
    [true, false, true, false],
    [true, true, false, true],
    [true, true, true, true],
    [true, true, false, false],
    [true, true, true, false],
  ])(
    `reloadDashboardsOnParamsChange: %s, reloadOnParamsChange: %s, withUid: %s, editMode: %s`,
    async (reloadDashboardsOnParamsChange, reloadOnParamsChange, withUid, editMode) => {
      config.featureToggles.reloadDashboardsOnParamsChange = reloadDashboardsOnParamsChange;
      setDashboardAPI(undefined);

      const dashboardScene = renderDashboard({ uid: withUid ? 'dash-1' : undefined }, { reloadOnParamsChange });
      const dashboardReloadSpy = jest.spyOn(getDashboardScenePageStateManager(), 'reloadDashboard');

      if (editMode) {
        await enterEditMode(dashboardScene);
      }

      const shouldReload = reloadDashboardsOnParamsChange && reloadOnParamsChange && withUid && !editMode;

      await updateTimeRange(dashboardScene);
      await jest.advanceTimersToNextTimerAsync();
      if (!shouldReload) {
        expect(dashboardReloadSpy).not.toHaveBeenCalled();
      } else {
        expect(dashboardReloadSpy).toHaveBeenCalled();
      }

      await updateMyVar(dashboardScene, '2');
      await jest.advanceTimersToNextTimerAsync();
      if (!shouldReload) {
        expect(dashboardReloadSpy).not.toHaveBeenCalled();
      } else {
        expect(dashboardReloadSpy).toHaveBeenCalled();
      }

      await updateScopes(['grafana']);
      await jest.advanceTimersToNextTimerAsync();
      if (!shouldReload) {
        expect(dashboardReloadSpy).not.toHaveBeenCalled();
      } else {
        expect(dashboardReloadSpy).toHaveBeenCalled();
      }

      getDashboardScenePageStateManager().clearDashboardCache();
      getDashboardScenePageStateManager().clearSceneCache();
      setDashboardAPI(undefined);
      await resetScenes();
    }
  );
});
