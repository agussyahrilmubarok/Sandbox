package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
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

func UpdateCPUUsage() {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		return
	}

	for i, usage := range percent {
		CPUUsage.With(prometheus.Labels{"core": fmt.Sprintf("%d", i)}).Set(usage)
	}
}

func UpdateMemoryUsage() {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	MemoryUsage.With(prometheus.Labels{"type": "ram"}).Set(float64(vmStat.Used))
}

func UpdateDiskUsage() {
	usage, err := disk.Usage("/")
	if err != nil {
		return
	}

	DiskUsage.With(prometheus.Labels{"disk": "/"}).Set(float64(usage.Used))
}
