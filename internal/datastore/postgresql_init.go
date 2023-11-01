package datastore

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"httpproject/internal/config"
	lr "httpproject/util/logger"
)

type PostgreDB struct {
	db     *gorm.DB
	dbInfo config.PostgresqlInfo
}

func NewDBInit(dbConfig config.PostgresqlInfo) (DBConnector, error) {

	postgresConnName := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbUsername, dbConfig.DbName, dbConfig.DbPassword)
	lr.Log.Info().Msgf("postgres connect Info : %s", postgresConnName)

	dbConn, err := gorm.Open(postgres.Open(postgresConnName), &gorm.Config{
		SkipDefaultTransaction: true, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		lr.Log.Error().Msgf("gorm Open Error : %v", err.Error())
		return nil, err
	}

	db := &PostgreDB{
		db:     dbConn,
		dbInfo: dbConfig,
	}

	return db, nil
}
