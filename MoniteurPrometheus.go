package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	CpuLoadAverageGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mac_cpu_charge",
		Help: "Nombre de processus en cours ou executable par rapport au nombre de coeur ||| Interprétation pour CPU avec 8coeurs: ≤6 Optimal | 8 normal | +9 Critique",
	})

	RamSPressureGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mac_ram_surcharge",
		Help: "Pourcentage de pression de la memoire vive (RAM) ||| Interprétation: ≤60 Optimal | ≤85 Normal | ≤95 Anormal | +95 Critique",
	})

	prometheus.MustRegister(CpuLoadAverageGauge)
	prometheus.MustRegister(RamSPressureGauge)

	go CpuAverage(CpuLoadAverageGauge)
	go RamSPressure(RamSPressureGauge)

	StartServer()
}
