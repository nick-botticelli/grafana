import { cloneDeep } from 'lodash';
import { of } from 'rxjs';

import { config } from '../config';
import { BackendSrvRequest, FetchError, FetchResponse, BackendSrv } from '../services';

import { getItem, setItem } from './userStorage';

const request = jest.fn<Promise<FetchResponse | FetchError>, BackendSrvRequest[]>();

const backendSrv = {
  fetch: (options: BackendSrvRequest) => {
    return of(request(options));
  },
} as unknown as BackendSrv;

jest.mock('../services', () => ({
  ...jest.requireActual('../services'),
  getBackendSrv: () => backendSrv,
}));

describe('userStorage', () => {
  const originalGetItem = Storage.prototype.getItem;
  const originalSetItem = Storage.prototype.setItem;
  const originalConfig = cloneDeep(config);

  beforeEach(() => {
    config.featureToggles.userStorageAPI = true;
    config.bootData.user.isSignedIn = true;
    config.bootData.user.uid = 'abc';
    request.mockReset();
    Storage.prototype.setItem = jest.fn();
    Storage.prototype.getItem = jest.fn();
  });

  afterEach(() => {
    Storage.prototype.setItem = originalSetItem;
    Storage.prototype.getItem = originalGetItem;
    config.featureToggles = originalConfig.featureToggles;
    config.bootData = originalConfig.bootData;
  });

  describe('getItem', () => {
    it('use localStorage if the feature flag is disabled', async () => {
      config.featureToggles.userStorageAPI = false;
      getItem('service', 'key');
      expect(localStorage.getItem).toHaveBeenCalled();
    });

    it('use localStorage if the user is not logged in', async () => {
      config.bootData.user.isSignedIn = false;
      getItem('service', 'key');
      expect(localStorage.getItem).toHaveBeenCalled();
    });

    it('use localStorage if the user storage is not found', async () => {
      request.mockReturnValue(Promise.reject({ status: 404 } as FetchError));
      await getItem('service', 'key');
      expect(localStorage.getItem).toHaveBeenCalled();
    });

    it('returns the value from the user storage', async () => {
      request.mockReturnValue(
        Promise.resolve({ status: 200, data: { spec: { data: { key: 'value' } } } } as FetchResponse)
      );
      const value = await getItem('service', 'key');
      expect(value).toBe('value');
    });
  });

  describe('setItem', () => {
    it('use localStorage if the feature flag is disabled', async () => {
      config.featureToggles.userStorageAPI = false;
      setItem('service', 'key', 'value');
      expect(localStorage.setItem).toHaveBeenCalled();
    });

    it('use localStorage if the user is not logged in', async () => {
      config.bootData.user.isSignedIn = false;
      setItem('service', 'key', 'value');
      expect(localStorage.setItem).toHaveBeenCalled();
    });

    it('creates a new user storage if it does not exist', async () => {
      request.mockReturnValueOnce(Promise.reject({ status: 404 } as FetchError));
      await setItem('service', 'key', 'value');
      expect(request).toHaveBeenCalledWith({
        url: '/apis/userstorage.grafana.app/v0alpha1/namespaces/default/user-storage/service:abc',
        method: 'GET',
        showErrorAlert: false,
      });
      expect(request).toHaveBeenCalledWith(
        expect.objectContaining({
          url: '/apis/userstorage.grafana.app/v0alpha1/namespaces/default/user-storage/',
          method: 'POST',
          data: {
            metadata: { labels: { service: 'service', user: 'abc' }, name: 'service:abc' },
            spec: {
              data: { key: 'value' },
            },
          },
        })
      );
    });

    it('updates the user storage if it exists', async () => {
      request.mockReturnValueOnce(
        Promise.resolve({
          status: 200,
          data: { metadata: { name: 'service:abc' }, spec: { data: { key: 'value' } } },
        } as FetchResponse)
      );
      await setItem('service', 'key', 'new-value');
      expect(request).toHaveBeenCalledWith({
        url: '/apis/userstorage.grafana.app/v0alpha1/namespaces/default/user-storage/service:abc',
        method: 'GET',
        showErrorAlert: false,
      });
      expect(request).toHaveBeenCalledWith(
        expect.objectContaining({
          url: '/apis/userstorage.grafana.app/v0alpha1/namespaces/default/user-storage/service:abc',
          method: 'PUT',
          data: {
            metadata: { name: 'service:abc' },
            spec: {
              data: { key: 'new-value' },
            },
          },
        })
      );
    });
  });
});
