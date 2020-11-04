package config

type configuration struct {
	LogFormat      string
	LogLevel       string
	SleepPeriod    int
	Database       DatabaseConfiguration
	ConnectionPool ConnectionPool
	Sage           Sage
}
