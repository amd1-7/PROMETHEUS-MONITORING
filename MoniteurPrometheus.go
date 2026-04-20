package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	CpuLoadAverageGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mac_cpu_charge",
		Help: "Nombre de processus en cours ou executable par rapport au nombre de coeur",
	}, []string{"seuil_critique"})

	RamSPressureGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mac_ram_surcharge",
		Help: "Pourcentage de pression de la memoire vive (RAM)",
	}, []string{"seuil_critique"})

	prometheus.MustRegister(CpuLoadAverageGauge)
	prometheus.MustRegister(RamSPressureGauge)

	go CpuAverage(CpuLoadAverageGauge)
	go RamSPressure(RamSPressureGauge)

	StartServer()
}

/* ./prometheus --config.file=Prometheus.yml */
