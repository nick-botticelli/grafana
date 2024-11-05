import {
  SceneDataLayerProviderState,
  SceneDataLayerProvider,
  SceneDataLayerSetBase,
  SceneComponentProps,
} from '@grafana/scenes';

import { AlertStatesDataLayer } from './AlertStatesDataLayer';
import { CorrelationsDataLayer } from './CorrelationsDataLayer';

export interface DashboardDataLayerSetState extends SceneDataLayerProviderState {
  alertStatesLayer?: AlertStatesDataLayer;
  annotationLayers: SceneDataLayerProvider[];
  correlationsLayer?: CorrelationsDataLayer;
}

export class DashboardDataLayerSet
  extends SceneDataLayerSetBase<DashboardDataLayerSetState>
  implements SceneDataLayerProvider
{
  public constructor(state: Partial<DashboardDataLayerSetState>) {
    super({
      ...state,
      name: state.name ?? 'Data layers',
      annotationLayers: state.annotationLayers ?? [],
    });

    this.addActivationHandler(() => this._onActivate());
  }

  private _onActivate() {
    this._subs.add(
      this.subscribeToState((newState, oldState) => {
        if (newState.annotationLayers !== oldState.annotationLayers) {
          this.querySub?.unsubscribe();
          this.subscribeToAllLayers(this.getAllLayers());
        }
      })
    );

    this.subscribeToAllLayers(this.getAllLayers());

    return () => {
      this.querySub?.unsubscribe();
    };
  }

  public addAnnotationLayer(layer: SceneDataLayerProvider) {
    this.setState({ annotationLayers: [...this.state.annotationLayers, layer] });
  }

  private getAllLayers() {
    const layers = [...this.state.annotationLayers];

    if (this.state.alertStatesLayer) {
      layers.push(this.state.alertStatesLayer);
    }

    if (this.state.correlationsLayer) {
      layers.push(this.state.correlationsLayer);
    }

    return layers;
  }

  public static Component = ({ model }: SceneComponentProps<DashboardDataLayerSet>) => {
    const { annotationLayers } = model.useState();

    return (
      <>
        {annotationLayers.map((layer) => (
          <layer.Component model={layer} key={layer.state.key} />
        ))}
      </>
    );
  };
}
