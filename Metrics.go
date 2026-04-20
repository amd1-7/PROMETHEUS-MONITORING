package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

func CpuAverage(metric *prometheus.GaugeVec) {
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
		metric.WithLabelValues("+9").Set(charge_cpu)

		time.Sleep(time.Second * time.Duration(15))
	}
}

func RamSPressure(metric *prometheus.GaugeVec) {
	for {
		ram, err := mem.VirtualMemory()
		if err != nil {
			log.Fatalf("erreur de recupération des information de la RAM: %v\n", err)
		}

		pourcentage := float64(ram.Total-ram.Available) / float64(ram.Total) * 100

		metric.WithLabelValues("90").Set(pourcentage)

		time.Sleep(time.Second * time.Duration(15))
	}
}
