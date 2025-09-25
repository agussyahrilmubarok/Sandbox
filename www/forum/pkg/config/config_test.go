package config

import (
	"testing"
)

func TestLoadViper(t *testing.T) {
	// Load configuration using LoadViper
	viper := LoadViper()

	appName := viper.GetString("APPLICATION.NAME")
	appPort := viper.GetInt("APPLICATION.PORT")

	host := viper.GetString("POSTGRES.HOST")
	port := viper.GetInt("POSTGRES.PORT")
	user := viper.GetString("POSTGRES.USER")
	password := viper.GetString("POSTGRES.PASSWORD")
	name := viper.GetString("POSTGRES.NAME")
	sslMode := viper.GetString("POSTGRES.SSL_MODE")
	maxOpenConns := viper.GetInt("POSTGRES.MAX_OPEN_CONNS")
	maxIdleConns := viper.GetInt("POSTGRES.MAX_IDLE_CONNS")
	maxIdleTime := viper.GetString("POSTGRES.MAX_IDLE_TIME")

	// Check APPLICATION values
	if appName != "forum-api" {
		t.Errorf("Expected APPLICATION.NAME to be 'forum-api', got %s", appName)
	}
	if appPort != 8080 {
		t.Errorf("Expected APPLICATION.PORT to be 8080, got %d", appPort)
	}

	// Check POSTGRES values
	if host != "localhost" {
		t.Errorf("Expected POSTGRES.HOST to be 'localhost', got %s", host)
	}
	if port != 5432 {
		t.Errorf("Expected POSTGRES.PORT to be 5432, got %d", port)
	}
	if user != "forum" {
		t.Errorf("Expected POSTGRES.USER to be 'forum', got %s", user)
	}
	if password != "my-secret-pw" {
		t.Errorf("Expected POSTGRES.PASSWORD to be 'my-secret-pw', got %s", password)
	}
	if name != "forum" {
		t.Errorf("Expected POSTGRES.NAME to be 'forum', got %s", name)
	}
	if sslMode != "disable" {
		t.Errorf("Expected POSTGRES.SSL_MODE to be 'disable', got %s", sslMode)
	}
	if maxOpenConns != 25 {
		t.Errorf("Expected POSTGRES.MAX_OPEN_CONNS to be 25, got %d", maxOpenConns)
	}
	if maxIdleConns != 25 {
		t.Errorf("Expected POSTGRES.MAX_IDLE_CONNS to be 25, got %d", maxIdleConns)
	}
	if maxIdleTime != "15m" {
		t.Errorf("Expected POSTGRES.MAX_IDLE_TIME to be '15m', got %s", maxIdleTime)
	}
}
