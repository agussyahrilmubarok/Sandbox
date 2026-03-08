package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// total_courses_count{job="course-app"}
	TotalCourses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "course_total_count",
			Help: "Total number of courses",
		},
	)

	// available_courses_count{job="course-app"}
	AvailableCourses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "course_available_count",
			Help: "Number of available courses (with seats available)",
		},
	)

	// reserved_courses_count{job="course-app"}
	ReservedCourses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "course_reserved_count",
			Help: "Number of reserved courses (seats taken)",
		},
	)
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(TotalCourses)
	prometheus.MustRegister(AvailableCourses)
	prometheus.MustRegister(ReservedCourses)
}
