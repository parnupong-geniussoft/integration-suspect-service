package configs

import (
	"os"
	"strconv"
	"time"
)

type Configs struct {
	App        Fiber
	PostgreSQL PostgreSQL
	Suspect    Suspect
	Retry      Retry
}

type Fiber struct {
	Host string
	Port string
}

// Database
type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}

// Suspect List
type Suspect struct {
	KtbSuspectListHost string
}

// Retry List
type Retry struct {
	RetryCount             int
	RetryMinWaitTimeSecond time.Duration
	RetryMaxWaitTimeSecond time.Duration
	RetryTimeoutSecond     time.Duration
}

func LoadEnv() Configs {
	cfg := Configs{}

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// Suspect Configs
	cfg.Suspect.KtbSuspectListHost = os.Getenv("KTB_SUSPECT_LIST_HOST")

	// Retry Configs
	cfg.Retry.RetryCount, _ = strconv.Atoi(os.Getenv("RETRY_COUNT"))
	cfg.Retry.RetryMinWaitTimeSecond, _ = time.ParseDuration(os.Getenv("RETRY_MIN_WAIT_TIME_SECOND"))
	cfg.Retry.RetryMaxWaitTimeSecond, _ = time.ParseDuration(os.Getenv("RETRY_MAX_WAIT_TIME_SECOND"))
	cfg.Retry.RetryTimeoutSecond, _ = time.ParseDuration(os.Getenv("RETRY_TIMEOUT_SECOND"))

	return cfg
}
