package main

import (
	"collector/internal/collect"
	"collector/internal/exporter"
	"collector/internal/log"
)

var port = 8080

func main() {
	log.InitLogPath()
	exporter.InitMetrics()

	go exporter.StartExporter(port)
	collect.StartCollecting()
}
