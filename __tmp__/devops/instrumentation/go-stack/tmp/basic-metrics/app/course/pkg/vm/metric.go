package vm

import (
	"fmt"
	"log"

	"example.com/course/internal/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func UpdateCPUUsage() {
	percent, err := cpu.Percent(0, true)
	if err != nil {
		log.Println("Error fetching CPU usage:", err)
		return
	}

	for i, usage := range percent {
		metrics.CPUUsage.With(prometheus.Labels{"core": fmt.Sprintf("%d", i)}).Set(usage)
	}
}

func UpdateMemoryUsage() {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error fetching memory usage:", err)
		return
	}

	metrics.MemoryUsage.With(prometheus.Labels{"type": "ram"}).Set(float64(vmStat.Used))
}

func UpdateDiskUsage() {
	usage, err := disk.Usage("/")
	if err != nil {
		log.Println("Error fetching disk usage:", err)
		return
	}

	metrics.DiskUsage.With(prometheus.Labels{"disk": "/"}).Set(float64(usage.Used))
}
