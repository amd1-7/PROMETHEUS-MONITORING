package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

func CpuAverage(metric prometheus.Gauge) {
	for {
		l, err := load.Avg()
		if err != nil {
			log.Fatalf("Erreur de récuperartion du nombre de processus en cours: %v \n", err)
		}

		coeurs, err := cpu.Counts(true)
		if err != nil {
			log.Fatalf("Erreur de récuperartion du nombre de coeurs du processeur (CPU): %v \n", err)
		}

		charge_cpu := l.Load1 / float64(coeurs) /* Recuperère le nombre de processus sur 1 minute */
		metric.Set(charge_cpu)

		time.Sleep(time.Second * time.Duration(15))
	}
}
