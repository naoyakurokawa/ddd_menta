package config

type contextKey string

const (
	DBKey       string     = "DB_KEY"
	AppErrorKey contextKey = "APP_ERROR_KEY"
)
