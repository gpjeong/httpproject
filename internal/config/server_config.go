package config

import (
	"httpproject/util"
	"httpproject/util/logger"
	"sync"
)

// 환경 변수 관련 정보
var (
	config       Server
	configOnce   sync.Once
	ServerName   string
	StartService = false
	//IsStandbyServerRunning bool
)

// PostgresqlInfo Postgresql 접속 정보
type PostgresqlInfo struct {
	DbName     string
	DbPort     string
	DbHost     string
	DbUsername string
	DbPassword string
}

// ApiInfo api 정보
type ApiInfo struct {
	ApiPort string
	ApiHost string
}

type Server struct {
	PostgresqlInfo PostgresqlInfo
	ApiInfo        ApiInfo
	Logger         logger.LoggerServer
}

var ServerConfig Server

func DefaultServiceConfigFromEnv() Server {
	configOnce.Do(func() {
		config = Server{
			PostgresqlInfo{
				util.GetEnv("DB_NAME", "rule_engine"),
				util.GetEnv("DB_PORT", "5432"),
				util.GetEnv("DB_HOST", "localhost"),
				util.GetEnv("DB_USERNAME", "rule"),
				util.GetEnv("DB_PASSWORD", "okestro2018"),
			},
			ApiInfo{
				util.GetEnv("ApiPort", "8080"),
				util.GetEnv("ApiHost", "localhost"),
			},
			logger.LoggerServer{
				Level:                 logger.GetLogLevelFromString(util.GetEnv("RULE_SERVER_LOGGER_LEVEL", "TRACE")),
				Directory:             util.GetEnv("RULE_LOG_DIRECTORY", "./data/logs"),
				Filename:              util.GetEnv("RULE_FILE_NAME", "rule_collector.log"),
				ConsoleLoggingEnabled: util.GetEnvAsBool("LOG_CONSOLE_LOGGING_ENABLED", true),
				EncodeLogsAsJson:      util.GetEnvAsBool("LOG_ENCODE_LOGS_AS_JSON", true),
				FileLoggingEnabled:    util.GetEnvAsBool("LOG_FILE_LOGGING_ENABLED", true),
				MaxSize:               util.GetEnvAsInt("LOG_MAX_SIZE", 200),
				MaxBackups:            util.GetEnvAsInt("LOG_MAX_BACKUPS", 30),
				MaxAge:                util.GetEnvAsInt("LOG_MAX_AGE", 90),
			},
		}
	})
	ServerConfig = config

	return config
}
