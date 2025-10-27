package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// TotalBookings tracks the total number of bookings
	TotalBookings = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_bookings_count",
			Help: "Total number of bookings made",
		},
	)

	// SuccessfulBookings tracks successful bookings
	SuccessfulBookings = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "successful_bookings_count",
			Help: "Total number of successful bookings",
		},
	)

	// FailedBookings tracks failed bookings
	FailedBookings = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "failed_bookings_count",
			Help: "Total number of failed bookings",
		},
	)

	// BookingStatus tracks bookings by their status
	BookingStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "booking_status_count",
			Help: "Count of bookings by their status (pending, confirmed, cancelled)",
		},
		[]string{"status"},
	)
)

func init() {
	// Register metrics with Prometheus
	prometheus.MustRegister(TotalBookings)
	prometheus.MustRegister(SuccessfulBookings)
	prometheus.MustRegister(FailedBookings)
	prometheus.MustRegister(BookingStatus)
}
