package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	CpuLoadAverageGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mac_cpu_charge",
		Help: "Nombre de processus en cours ou executable par rapport au nombre de coeur",
	})

	prometheus.MustRegister(CpuLoadAverageGauge)

	go CpuAverage(CpuLoadAverageGauge)

	StartServer()
}
