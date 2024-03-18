package config

type Config struct {
	HttpPort string
	PostgresHost string
	PostgresPort int
	PostgresUser string
	PostgresPassword string
	PostgresDatabase string
}

func Load() Config {
	var config Config
    config.HttpPort = ":7070"
    config.PostgresHost = "localhost"
    config.PostgresPort = 5432
    config.PostgresUser = "postgres"
    config.PostgresPassword = "0"
    config.PostgresDatabase = "doctordb"
	
	return config
}