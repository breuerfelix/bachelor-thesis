package main

import (
	"github.com/prometheus/common/expfmt"
	"net/http"
	"time"
	"fmt"
)

var INTERVAL = 5 * time.Second

func fetchMetrics(url string) {
	for {
		resp, _ := http.Get(url)
		defer resp.Body.Close()

		var parser expfmt.TextParser
		mf, _ := parser.TextToMetricFamilies(resp.Body)

		metric := "com_hivemq_system_os_global_cpu_total_total"
		value := mf[metric].GetMetric()[0].Gauge.GetValue()

		fmt.Printf("CPU load: %.1f\n", value)

		time.Sleep(INTERVAL)
	}
}

func main() {
	go fetchMetrics("http://localhost:9399/metrics")

	// ...
}

// Output:
// CPU load: 5.0
// CPU load: 10.3
// CPU load: 7.5
