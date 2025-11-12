package collectors

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"

	"go.opentelemetry.io/otel/metric"
)

func init() {
	key := "healthCheck"
	registerModule(key, NewHealth())

}

type ModuleHealth struct {
	name     string
	defaults bool
	desc     []*MetricDescriptor
}

func NewHealth() *ModuleHealth {
	return &ModuleHealth{
		defaults: true,
	}
}

func (_m *ModuleHealth) SetConfig(inf interface{}) Module {
	data, _ := json.Marshal(inf)
	json.NewDecoder(bytes.NewReader(data)).Decode(&_m)
	return _m
}

func (_m *ModuleHealth) Init(key string) {
	_m.name = key
	_m.desc = []*MetricDescriptor{
		{
			Key:      "up",
			Name:     "unisphere_up",
			Desc:     "check to scrape data is success",
			Unit:     "",
			TypeName: "gauge",
		},
	}
}

func (_m *ModuleHealth) Run(logger *slog.Logger, col *Collector) {
	meter := col.MeterProvider.Meter(_m.name)

	// Register Metrics...
	var observableMap map[string]metric.Float64Observable
	observableMap = CreateMapMetricDescriptor(meter, _m.desc, logger)

	// Register Metrics for Observables...
	var observableArray []metric.Observable
	for _, obserable := range observableMap {
		observableArray = append(observableArray, obserable)
	}

	meter.RegisterCallback(func(ctx context.Context, observer metric.Observer) error {
		clientAttrs := metric.WithAttributes(append(col.customLabels, col.detectLabels...)...)

		var health float64
		if col.success == true {
			health = 1
		} else {
			health = 0
		}
		observer.ObserveFloat64(observableMap["up"], health, clientAttrs)
		return nil
	}, observableArray...)
}
