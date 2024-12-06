import { BehaviorSubject, Observable, pairwise, Subscription } from 'rxjs';

import { config } from '@grafana/runtime';

export abstract class ScopesServiceBase<T> {
  private _state: BehaviorSubject<T>;
  protected _fetchSub: Subscription | undefined;
  protected _apiGroup = 'scope.grafana.app';
  protected _apiVersion = 'v0alpha1';
  protected _apiNamespace = config.namespace ?? 'default';

  protected constructor(initialState: T) {
    this._state = new BehaviorSubject<T>(Object.freeze(initialState));
  }

  public get state(): T {
    return this._state.getValue();
  }

  public get stateObservable(): Observable<T> {
    return this._state.asObservable();
  }

  public subscribeToState = (cb: (newState: T, prevState: T) => void): Subscription => {
    return this._state.pipe(pairwise()).subscribe(([prevState, newState]) => cb(newState, prevState));
  };

  protected updateState = (newState: Partial<T>) => {
    this._state.next(Object.freeze({ ...this.state, ...newState }));
  };
}
