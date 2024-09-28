package config

var (
	AWS_REGION       = GetEnv("DB_HOST", "db")
	AWS_USER_POOL_ID = GetEnv("DB_PORT", "5432")
)
