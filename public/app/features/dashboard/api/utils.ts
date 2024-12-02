import { config, locationService } from '@grafana/runtime';
import { DashboardV2Spec } from '@grafana/schema/dist/esm/schema/dashboard/v2alpha0/dashboard.gen';
import { DashboardDataDTO, DashboardDTO } from 'app/types';

import { DashboardWithAccessInfo } from './types';

export function getDashboardsApiVersion() {
  const forcingOldDashboardArch = locationService.getSearch().get('scenes') === 'false';

  // if dashboard scene is disabled, use legacy API response for the old architecture
  if (!config.featureToggles.dashboardScene || forcingOldDashboardArch) {
    // for old architecture, use v0 API for k8s dashboards
    if (config.featureToggles.kubernetesDashboards) {
      return 'v0';
    }

    return 'legacy';
  }

  if (config.featureToggles.useV2DashboardsAPI) {
    return 'v2';
  }

  if (config.featureToggles.kubernetesDashboards) {
    return 'v0';
  }

  return 'legacy';
}

export function isDashboardResource(
  obj?: DashboardDTO | DashboardWithAccessInfo<DashboardV2Spec | DashboardDataDTO> | null
): obj is DashboardWithAccessInfo<DashboardV2Spec> {
  if (!obj) {
    return false;
  }

  return 'kind' in obj && obj.kind === 'DashboardWithAccessInfo';
}

export function isDashboardV2Spec(obj: DashboardDataDTO | DashboardV2Spec): obj is DashboardV2Spec {
  return 'elements' in obj;
}
