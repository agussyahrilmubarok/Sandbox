package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var (
	// CPU % per core
	CPUUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage_percent",
			Help: "CPU usage percentage per core",
		},
		[]string{"core"},
	)

	// Memory used (bytes)
	MemoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_memory_usage_bytes",
			Help: "Current memory usage in bytes",
		},
		[]string{"type"},
	)

	// Memory total (bytes)
	TotalMemory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_memory_total_bytes",
			Help: "Total memory available in bytes",
		},
		[]string{"type"},
	)

	// Disk usage (bytes)
	DiskUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_disk_usage_bytes",
			Help: "Disk used in bytes",
		},
		[]string{"disk"},
	)

	// Disk total (bytes) → untuk hitung %
	DiskTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_disk_total_bytes",
			Help: "Total disk capacity in bytes",
		},
		[]string{"disk"},
	)
)

func init() {
	prometheus.MustRegister(CPUUsage)
	prometheus.MustRegister(MemoryUsage)
	prometheus.MustRegister(TotalMemory)
	prometheus.MustRegister(DiskUsage)
	prometheus.MustRegister(DiskTotal)
}

func UpdateCPUUsage() {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		return
	}

	for i, usage := range percent {
		CPUUsage.WithLabelValues(fmt.Sprintf("%d", i)).Set(usage)
	}
}

func UpdateMemoryUsage() {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	MemoryUsage.WithLabelValues("ram").Set(float64(vmStat.Used))
	TotalMemory.WithLabelValues("ram").Set(float64(vmStat.Total))
}

func UpdateDiskUsage() {
	usage, err := disk.Usage("/")
	if err != nil {
		return
	}

	DiskUsage.WithLabelValues("/").Set(float64(usage.Used))
	DiskTotal.WithLabelValues("/").Set(float64(usage.Total))
}
