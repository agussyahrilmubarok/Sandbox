package main

import (
	"sync"

	"greenlight.id/internal/data"
	"greenlight.id/internal/json_log"
	"greenlight.id/internal/mailer"
)

type application struct {
	config config
	logger *json_log.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}
