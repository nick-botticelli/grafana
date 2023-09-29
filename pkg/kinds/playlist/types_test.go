package playlist

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPlaylistClone(t *testing.T) {
	src := Playlist{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Playlist",
			APIVersion: APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              "k8sname",
			ResourceVersion:   "12345",
			CreationTimestamp: metav1.NewTime(time.Now()),
			Annotations: map[string]string{
				"grafana.com/updatedTime": time.Now().Format(time.RFC3339),
			},
		},
		Spec: &Spec{
			Uid:      "TheUID",
			Name:     "A title",
			Interval: "20s",
			Items: []Item{
				{Type: ItemTypeDashboardByTag, Value: "graph-ng"},
			},
		},
	}
	copy := src.DeepCopyObject()

	json0, err := json.Marshal(src)
	require.NoError(t, err)
	json1, err := json.Marshal(copy)
	require.NoError(t, err)

	require.JSONEq(t, string(json0), string(json1))
}
