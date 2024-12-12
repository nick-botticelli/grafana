package traceql

import (
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/tempo/pkg/tempopb"
)

const bucketFieldName = "__bucket"

func transformExemplarToFrame(name string, series *tempopb.TimeSeries, isHistogram bool) *data.Frame {
	exemplars := series.Exemplars

	// Setup fields for basic data
	fields := []*data.Field{
		data.NewField("Time", nil, []time.Time{}),
		data.NewField("Value", nil, []float64{}),
		data.NewField("traceId", nil, []string{}),
	}

	fields[2].Config = &data.FieldConfig{
		DisplayName: "Trace ID",
	}

	// Add fields for each label to be able to link exemplars to the series
	for _, label := range series.Labels {
		fields = append(fields, data.NewField(label.GetKey(), nil, []string{}))
	}

	// Special case for histograms, requires a bucket field
	if isHistogram {
		fields = append(fields, data.NewField(bucketFieldName, nil, []float64{}))
	}

	frame := &data.Frame{
		RefID:  name,
		Name:   "exemplar",
		Fields: fields,
		Meta: &data.FrameMeta{
			DataTopic: "annotations",
		},
	}

	for _, exemplar := range exemplars {
		_, labels := transformLabelsAndGetName(exemplar.GetLabels())
		traceId := labels["trace:id"]
		if traceId != "" {
			traceId = strings.ReplaceAll(traceId, "\"", "")
		}

		// Add basic data
		frame.AppendRow(time.UnixMilli(exemplar.GetTimestampMs()), exemplar.GetValue(), traceId)

		// Add bucket for histograms
		if isHistogram {
			bucketField, _ := frame.FieldByName(bucketFieldName)
			if bucketField != nil {
				bucketField.Append(exemplar.GetValue())
			}
		}

		// Add labels
		for _, label := range series.Labels {
			field, _ := frame.FieldByName(label.GetKey())
			if field != nil {
				val, _ := metricsValueToString(label.GetValue())
				field.Append(val)
			}
		}
	}
	return frame
}
