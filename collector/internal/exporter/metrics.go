package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var MyCounter prometheus.Counter
var MyGauge prometheus.Gauge

func InitMetrics() {
	MyCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_counter",
		Help: "My Counter",
	})
	MyGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "my_gauge",
		Help: "My Gauge",
	})
}
