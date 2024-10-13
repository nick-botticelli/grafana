import { renderHook, screen, within } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { first, noop } from 'lodash';
import { Routes, Route } from 'react-router-dom-v5-compat';
import { render } from 'test/test-utils';

import { config } from '@grafana/runtime';
import { contextSrv } from 'app/core/core';
import {
  AlertmanagerGroup,
  MatcherOperator,
  ObjectMatcher,
  RouteWithID,
} from 'app/plugins/datasource/alertmanager/types';
import { ReceiversState } from 'app/types/alerting';

import { useAlertmanagerAbilities } from '../../hooks/useAbilities';
import { mockAlertGroup, mockAlertmanagerAlert, mockReceiversState } from '../../mocks';
import { AlertmanagerProvider } from '../../state/AlertmanagerContext';
import { GRAFANA_RULES_SOURCE_NAME } from '../../utils/datasource';

import {
  AUTOGENERATED_ROOT_LABEL_NAME,
  Policy,
  TimingOptionsMeta,
  isAutoGeneratedRootAndSimplifiedEnabled,
  useCreateDropdownMenuActions,
} from './Policy';

jest.mock('../../hooks/useAbilities', () => ({
  ...jest.requireActual('../../hooks/useAbilities'),
  useAlertmanagerAbilities: jest.fn(),
}));

const useAlertmanagerAbilitiesMock = jest.mocked(useAlertmanagerAbilities);

describe('Policy', () => {
  beforeAll(() => {
    jest.spyOn(contextSrv, 'hasPermission').mockReturnValue(true);
    useAlertmanagerAbilitiesMock.mockReturnValue([
      [true, true],
      [true, true],
      [true, true],
    ]);
  });

  it('should render a policy tree', async () => {
    const onEditPolicy = jest.fn();
    const onAddPolicy = jest.fn();
    const onDeletePolicy = jest.fn();
    const onShowAlertInstances = jest.fn(
      (alertGroups: AlertmanagerGroup[], matchers?: ObjectMatcher[] | undefined) => {}
    );

    const routeTree = mockRoutes;
    const user = userEvent.setup();

    renderPolicy(
      <Policy
        routeTree={routeTree}
        currentRoute={routeTree}
        alertManagerSourceName={GRAFANA_RULES_SOURCE_NAME}
        onEditPolicy={onEditPolicy}
        onAddPolicy={onAddPolicy}
        onDeletePolicy={onDeletePolicy}
        onShowAlertInstances={onShowAlertInstances}
      />
    );

    // should have default policy
    const defaultPolicy = screen.getByTestId('am-root-route-container');
    expect(defaultPolicy).toBeInTheDocument();
    expect(within(defaultPolicy).getByText('Default policy')).toBeVisible();

    // click "more actions" and check if we can edit and delete
    expect(within(defaultPolicy).getByTestId('more-actions')).toBeInTheDocument();
    await user.click(within(defaultPolicy).getByTestId('more-actions'));

    // should be editable
    const editDefaultPolicy = screen.getByRole('menuitem', { name: 'Edit' });
    expect(editDefaultPolicy).toBeInTheDocument();
    expect(editDefaultPolicy).toBeEnabled();
    await user.click(editDefaultPolicy);
    expect(onEditPolicy).toHaveBeenCalledWith(routeTree, true);

    // should not be deletable
    expect(screen.queryByRole('menuitem', { name: 'Delete' })).not.toBeInTheDocument();

    // default policy should show the metadata

    // no continue matching
    expect(within(defaultPolicy).queryByTestId('continue-matching')).not.toBeInTheDocument();

    // for matching instances
    // expect(within(defaultPolicy).getByTestId('matching-instances')).toHaveTextContent('0instances');

    // for contact point
    expect(within(defaultPolicy).getByTestId('contact-point')).toHaveTextContent('grafana-default-email');
    expect(within(defaultPolicy).getByRole('link', { name: 'grafana-default-email' })).toBeInTheDocument();

    // for grouping
    expect(within(defaultPolicy).getByTestId('grouping')).toHaveTextContent('grafana_folder, alertname');

    // no timings
    expect(within(defaultPolicy).queryByTestId('mute-timings')).not.toBeInTheDocument();
    expect(within(defaultPolicy).queryByTestId('active-timings')).not.toBeInTheDocument();

    // for timing options
    expect(within(defaultPolicy).getByTestId('timing-options')).toHaveTextContent(
      'Wait 30s to group instances · Wait 5m before sending updates · Repeated every 4h'
    );

    // should have custom policies
    const customPolicies = screen.getAllByTestId('am-route-container');
    expect(customPolicies).toHaveLength(3);

    // all policies should be editable and deletable
    for (const container of customPolicies) {
      const policy = within(container);

      // click "more actions" and check if we can delete
      await user.click(policy.getByTestId('more-actions'));
      expect(screen.queryByRole('menuitem', { name: 'Edit' })).toBeEnabled();
      expect(screen.queryByRole('menuitem', { name: 'Delete' })).toBeEnabled();

      await user.click(screen.getByRole('menuitem', { name: 'Delete' }));
      expect(onDeletePolicy).toHaveBeenCalled();
    }

    // first custom policy should have the correct information
    const firstPolicy = customPolicies[0];
    expect(within(firstPolicy).getByTestId('label-matchers')).toHaveTextContent(/^team \= operations$/);
    expect(within(firstPolicy).getByTestId('continue-matching')).toBeInTheDocument();
    // expect(within(firstPolicy).getByTestId('matching-instances')).toHaveTextContent('0instances');
    expect(within(firstPolicy).getByTestId('contact-point')).toHaveTextContent('provisioned-contact-point');
    expect(within(firstPolicy).getByTestId('mute-timings')).toHaveTextContent('Muted whenmt-1');
    expect(within(firstPolicy).getByTestId('active-timings')).toHaveTextContent('Active whenmt-2');
    expect(within(firstPolicy).getByTestId('inherited-properties')).toHaveTextContent('Inherited2 properties');

    // second custom policy should be correct
    const secondPolicy = customPolicies[1];
    expect(within(secondPolicy).getByTestId('label-matchers')).toHaveTextContent(/^region \= EMEA$/);
    expect(within(secondPolicy).queryByTestId('continue-matching')).not.toBeInTheDocument();
    expect(within(secondPolicy).queryByTestId('mute-timings')).not.toBeInTheDocument();
    expect(within(secondPolicy).queryByTestId('active-timings')).not.toBeInTheDocument();
    expect(within(secondPolicy).getByTestId('inherited-properties')).toHaveTextContent('Inherited3 properties');

    // third custom policy should be correct
    const thirdPolicy = customPolicies[2];
    expect(within(thirdPolicy).getByTestId('label-matchers')).toHaveTextContent(
      /^foo = barbar = bazbaz = quxasdf = asdftype = diskand 1 more$/
    );
  });

  it('should show export option when export is allowed and supported returns true', async () => {
    const onEditPolicy = jest.fn();
    const onAddPolicy = jest.fn();
    const onDeletePolicy = jest.fn();
    const onShowAlertInstances = jest.fn(
      (alertGroups: AlertmanagerGroup[], matchers?: ObjectMatcher[] | undefined) => {}
    );

    const routeTree = mockRoutes;
    const user = userEvent.setup();

    renderPolicy(
      <Policy
        routeTree={routeTree}
        currentRoute={routeTree}
        alertManagerSourceName={GRAFANA_RULES_SOURCE_NAME}
        onEditPolicy={onEditPolicy}
        onAddPolicy={onAddPolicy}
        onDeletePolicy={onDeletePolicy}
        onShowAlertInstances={onShowAlertInstances}
        isAutoGenerated={false}
      />
    );
    // should have default policy
    const defaultPolicy = screen.getByTestId('am-root-route-container');
    // click "more actions"
    expect(within(defaultPolicy).getByTestId('more-actions')).toBeInTheDocument();
    await user.click(within(defaultPolicy).getByTestId('more-actions'));
    expect(screen.getByRole('menuitem', { name: 'Export' })).toBeInTheDocument();
  });

  it('should not show export option when is not supported', async () => {
    const onEditPolicy = jest.fn();
    const onAddPolicy = jest.fn();
    const onDeletePolicy = jest.fn();
    const onShowAlertInstances = jest.fn(
      (alertGroups: AlertmanagerGroup[], matchers?: ObjectMatcher[] | undefined) => {}
    );

    const routeTree = mockRoutes;

    useAlertmanagerAbilitiesMock.mockReturnValue([
      [true, true],
      [true, true],
      [false, true],
    ]);

    const user = userEvent.setup();

    renderPolicy(
      <Policy
        routeTree={routeTree}
        currentRoute={routeTree}
        alertManagerSourceName={GRAFANA_RULES_SOURCE_NAME}
        onEditPolicy={onEditPolicy}
        onAddPolicy={onAddPolicy}
        onDeletePolicy={onDeletePolicy}
        onShowAlertInstances={onShowAlertInstances}
      />
    );
    // should have default policy
    const defaultPolicy = screen.getByTestId('am-root-route-container');
    // click "more actions"
    expect(within(defaultPolicy).getByTestId('more-actions')).toBeInTheDocument();
    await user.click(within(defaultPolicy).getByTestId('more-actions'));
    expect(screen.queryByRole('menuitem', { name: 'Export' })).not.toBeInTheDocument();
  });

  it('should not show export option when is not allowed', async () => {
    const onEditPolicy = jest.fn();
    const onAddPolicy = jest.fn();
    const onDeletePolicy = jest.fn();
    const onShowAlertInstances = jest.fn(
      (alertGroups: AlertmanagerGroup[], matchers?: ObjectMatcher[] | undefined) => {}
    );

    const routeTree = mockRoutes;

    useAlertmanagerAbilitiesMock.mockReturnValue([
      [true, true],
      [true, true],
      [true, false],
    ]);

    const user = userEvent.setup();

    renderPolicy(
      <Policy
        routeTree={routeTree}
        currentRoute={routeTree}
        alertManagerSourceName={GRAFANA_RULES_SOURCE_NAME}
        onEditPolicy={onEditPolicy}
        onAddPolicy={onAddPolicy}
        onDeletePolicy={onDeletePolicy}
        onShowAlertInstances={onShowAlertInstances}
      />
    );
    // should have default policy
    const defaultPolicy = screen.getByTestId('am-root-route-container');
    // click "more actions"
    expect(within(defaultPolicy).getByTestId('more-actions')).toBeInTheDocument();
    await user.click(within(defaultPolicy).getByTestId('more-actions'));
    expect(screen.queryByRole('menuitem', { name: 'Export' })).not.toBeInTheDocument();
  });

  it('should not allow editing readOnly policy tree', () => {
    const routeTree: RouteWithID = { id: '0', routes: [{ id: '1' }] };

    renderPolicy(
      <Policy
        readOnly
        routeTree={routeTree}
        currentRoute={routeTree}
        alertManagerSourceName={GRAFANA_RULES_SOURCE_NAME}
        onEditPolicy={noop}
        onAddPolicy={noop}
        onDeletePolicy={noop}
        onShowAlertInstances={noop}
      />
    );

    expect(screen.queryByRole('button', { name: 'Edit' })).not.toBeInTheDocument();
  });

  it.skip('should show matching instances', () => {
    const routeTree: RouteWithID = {
      id: '0',
      routes: [{ id: '1', object_matchers: [['foo', eq, 'bar']] }],
    };

    const matchingGroups: AlertmanagerGroup[] = [
      mockAlertGroup({
        labels: {},
        alerts: [mockAlertmanagerAlert({ labels: { foo: 'bar' } }), mockAlertmanagerAlert({ labels: { foo: 'bar' } })],
      }),
      mockAlertGroup({
        labels: {},
        alerts: [mockAlertmanagerAlert({ labels: { bar: 'baz' } })],
      }),
    ];

    renderPolicy(
      <Policy
        readOnly
        alertGroups={matchingGroups}
        routeTree={routeTree}
        currentRoute={routeTree}
        alertManagerSourceName={GRAFANA_RULES_SOURCE_NAME}
        onEditPolicy={noop}
        onAddPolicy={noop}
        onDeletePolicy={noop}
        onShowAlertInstances={noop}
      />
    );

    const defaultPolicy = screen.getByTestId('am-root-route-container');
    expect(within(defaultPolicy).getByTestId('matching-instances')).toHaveTextContent('1instance');
    const customPolicy = screen.getByTestId('am-route-container');
    expect(within(customPolicy).getByTestId('matching-instances')).toHaveTextContent('2instances');
  });

  it('should show warnings and errors', () => {
    const routeTree: RouteWithID = {
      id: '0', // this one should show an error
      receiver: 'broken-receiver',
      routes: [{ id: '1', object_matchers: [] }], // this one should show a warning
    };

    const receiversState: ReceiversState = mockReceiversState();

    renderPolicy(
      <Policy
        readOnly
        routeTree={routeTree}
        currentRoute={routeTree}
        contactPointsState={receiversState}
        alertManagerSourceName={GRAFANA_RULES_SOURCE_NAME}
        onEditPolicy={noop}
        onAddPolicy={noop}
        onDeletePolicy={noop}
        onShowAlertInstances={noop}
      />
    );

    const defaultPolicy = screen.getByTestId('am-root-route-container');
    expect(within(defaultPolicy).queryByTestId('matches-all')).not.toBeInTheDocument();
    expect(within(defaultPolicy).getByText('1 error')).toBeInTheDocument();

    const customPolicy = screen.getByTestId('am-route-container');
    expect(within(customPolicy).getByTestId('matches-all')).toBeInTheDocument();
  });
});

// Doesn't matter which path the routes use, it just needs to match the initialEntries history entry to render the element
const renderPolicy = (element: JSX.Element) =>
  render(
    <Routes>
      <Route path={'/'} element={<AlertmanagerProvider accessType="notification">{element}</AlertmanagerProvider>} />
    </Routes>,
    {
      historyOptions: {
        initialEntries: ['/'],
      },
    }
  );

const eq = MatcherOperator.equal;

const mockRoutes: RouteWithID = {
  id: '0',
  receiver: 'grafana-default-email',
  group_by: ['grafana_folder', 'alertname'],
  routes: [
    {
      id: '1',
      receiver: 'provisioned-contact-point',
      object_matchers: [['team', eq, 'operations']],
      mute_time_intervals: ['mt-1'],
      active_time_intervals: ['mt-2'],
      continue: true,
      routes: [
        {
          id: '2',
          object_matchers: [['region', eq, 'EMEA']],
        },
        {
          id: '3',
          receiver: 'grafana-default-email',
          object_matchers: [
            ['foo', eq, 'bar'],
            ['bar', eq, 'baz'],
            ['baz', eq, 'qux'],
            ['asdf', eq, 'asdf'],
            ['type', eq, 'disk'],
            ['severity', eq, 'critical'],
          ],
        },
      ],
    },
  ],
  group_wait: '30s',
  group_interval: undefined,
  repeat_interval: undefined,
};

describe('isAutoGeneratedRootAndSimplifiedEnabled', () => {
  it('returns false when simplified routing is not enabled', () => {
    const route: RouteWithID = {
      id: '1',
      object_matchers: [['label', MatcherOperator.equal, 'true']],
    };
    config.featureToggles.alertingSimplifiedRouting = false;
    expect(isAutoGeneratedRootAndSimplifiedEnabled(route)).toBe(false);
  });

  it('returns false when object_matchers is not defined', () => {
    const route: RouteWithID = {
      id: '1',
    };
    config.featureToggles.alertingSimplifiedRouting = true;
    expect(isAutoGeneratedRootAndSimplifiedEnabled(route)).toBe(false);
  });

  it('returns true when object_matchers contains AUTOGENERATED_ROOT_LABEL_NAME, and simplified routing is enabled', () => {
    const route: RouteWithID = {
      id: '1',
      object_matchers: [[AUTOGENERATED_ROOT_LABEL_NAME, MatcherOperator.equal, 'true']],
    };
    config.featureToggles.alertingSimplifiedRouting = true;
    expect(isAutoGeneratedRootAndSimplifiedEnabled(route)).toBe(true);
  });

  it('returns false when object_matchers does not contain AUTOGENERATED_ROOT_LABEL_NAME, and simplified routing is enabled', () => {
    const route: RouteWithID = {
      id: '1',
      object_matchers: [['label', MatcherOperator.equal, 'true']],
    };
    config.featureToggles.alertingSimplifiedRouting = true;
    expect(isAutoGeneratedRootAndSimplifiedEnabled(route)).toBe(false);
  });
});

describe('useCreateDropdownMenuActions', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });
  const openDetailModal = jest.fn();
  const currentRoute: RouteWithID = { id: '0', routes: [{ id: '1' }] };
  const toggleShowExportDrawer = jest.fn();
  const onDeletePolicy = jest.fn();
  const testCases = [
    {
      isAutoGenerated: false,
      isDefaultPolicy: true,
      provisioned: false,
      expectedMenu: ['edit-policy', 'export-policy'],
    },
    {
      isAutoGenerated: false,
      isDefaultPolicy: true,
      provisioned: true,
      expectedMenu: ['edit-policy', 'export-policy'],
    },
    {
      isAutoGenerated: false,
      isDefaultPolicy: false,
      provisioned: false,
      expectedMenu: ['edit-policy', 'delete-policy'],
    },
    {
      isAutoGenerated: false,
      isDefaultPolicy: false,
      provisioned: true,
      expectedMenu: ['edit-policy', 'delete-policy'],
    },
    { isAutoGenerated: true, isDefaultPolicy: true, provisioned: true, expectedMenu: ['edit-policy'] },
    { isAutoGenerated: true, isDefaultPolicy: false, provisioned: false, expectedMenu: ['edit-policy'] },
    { isAutoGenerated: true, isDefaultPolicy: true, provisioned: false, expectedMenu: ['edit-policy'] },
    { isAutoGenerated: true, isDefaultPolicy: false, provisioned: true, expectedMenu: ['edit-policy'] },
  ];

  testCases.forEach(({ isAutoGenerated, isDefaultPolicy, provisioned, expectedMenu }) => {
    it(`Having all the permissions returns ${expectedMenu.length} menu items for isAutoGenerated=${isAutoGenerated}, isDefaultPolicy=${isDefaultPolicy}, provisioned=${provisioned}`, () => {
      useAlertmanagerAbilitiesMock.mockReturnValue([
        [true, true],
        [true, true],
        [true, true],
      ]);
      const { result } = renderHook(() =>
        useCreateDropdownMenuActions(
          isAutoGenerated,
          isDefaultPolicy,
          provisioned,
          openDetailModal,
          currentRoute,
          toggleShowExportDrawer,
          onDeletePolicy
        )
      );

      const menuItemsKeys = result.current.map((item) => item.key ?? '');
      expect(menuItemsKeys).toEqual(expectedMenu);
    });
  });
});

describe('TimingOptionsMeta', () => {
  it('should render nothing without options', () => {
    render(<TimingOptionsMeta timingOptions={{}} />);
    expect(screen.queryByText(/wait/i)).not.toBeInTheDocument();
  });

  it('should render only repeat interval', () => {
    render(<TimingOptionsMeta timingOptions={{ repeat_interval: '5h' }} />);
    expect(screen.getByText(/repeated every/i)).toBeInTheDocument();
    expect(screen.getByText('5h')).toBeInTheDocument();
  });

  it('should render all options', () => {
    render(<TimingOptionsMeta timingOptions={{ group_wait: '30s', group_interval: '5m', repeat_interval: '4h' }} />);
    expect(
      first(
        screen.getAllByText(
          (_, element) =>
            element?.textContent === 'Wait 30s to group instances · Wait 5m before sending updates · Repeated every 4h',
          { collapseWhitespace: false, trim: false, exact: true }
        )
      )
    ).toBeInTheDocument();
  });
});
