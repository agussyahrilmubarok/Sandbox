package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// member_total_count{job="member-app"}
	TotalMembers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "member_total_count",
			Help: "Total number of members",
		},
	)
)

func init() {
	prometheus.MustRegister(TotalMembers)
}
