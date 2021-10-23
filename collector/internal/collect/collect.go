package collect

import (
	"fmt"
	"math/rand"
	"time"

	"collector/internal/exporter"
	"collector/internal/log"
)

var cycleInterval = 10

func StartCollecting() {
	for range time.Tick(time.Second * time.Duration(cycleInterval)) {
		log.Log("Starting collect cycle")
		cycle()
	}
}

func cycle() {
	exporter.MyCounter.Add(1)
	log.Log("Counter increased")

	random := rand.Float64()
	exporter.MyGauge.Set(random)
	log.Log("Gauge updated " + fmt.Sprintf("%f", random))

	log.Log("Cycle Finished")
}
