package datastore

import (
	"httpproject/internal/config"
	"httpproject/util/logger"
)

type TaskManager struct {
	sql *DataStore
}

var TaskManger *TaskManager

func DBService() *TaskManager {
	TaskManger = TaskManger.CheckDB()
	return TaskManger
}

func NewTaskManager() *TaskManager {
	data, err := NewFromConfig()
	if err != nil {
		logger.Log.Error().Msgf("TaskManager error : %v", err.Error())
	}
	return &TaskManager{
		sql: data,
	}
}

func (t TaskManager) CheckDB() *TaskManager {
	if !t.sql.CheckDBConnection() {
		t = *NewTaskManager()
		logger.Log.Debug().Msgf("CONNECTION")
	}
	return &t
}

type DataStore struct {
	db     DBConnector
	dbInfo config.PostgresqlInfo
}

func NewFromConfig() (*DataStore, error) {
	var db DBConnector
	var err error = nil
	serverConfig := config.DefaultServiceConfigFromEnv()
	dbInfo := serverConfig.PostgresqlInfo

	db, err = NewDBInit(dbInfo)
	if err != nil {
		return nil, err
	}

	return &DataStore{db: db, dbInfo: dbInfo}, nil
}
