package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// system_cpu_usage_percent{job="member-app"}
	// system_cpu_usage_percent{job="member-app", core="0"}
	// avg(rate(system_cpu_usage_percent{job="member-app"}[5m])) by (core)
	CPUUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage_percent",
			Help: "Percentage of CPU usage",
		},
		[]string{"core"},
	)

	// system_memory_usage_bytes{job="member-app", type="ram"}
	MemoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_memory_usage_bytes",
			Help: "Amount of memory used in bytes",
		},
		[]string{"type"},
	)

	// system_memory_total_bytes{job="member-app", type="ram"}
	TotalMemory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_memory_total_bytes",
			Help: "Total amount of memory in bytes",
		},
		[]string{"type"},
	)

	// system_disk_usage_bytes{job="member-app", disk="/"}
	DiskUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_disk_usage_bytes",
			Help: "Amount of disk space used in bytes",
		},
		[]string{"disk"},
	)
)

func init() {
	prometheus.MustRegister(CPUUsage)
	prometheus.MustRegister(MemoryUsage)
	prometheus.MustRegister(DiskUsage)
}
