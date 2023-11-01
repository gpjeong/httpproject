package datastore

import (
	"httpproject/util/logger"
)

func (t *TaskManager) CreateData(data interface{}) (int64, error) {
	logger.Log.Debug().Msgf("Create Data :", data)
	//var rowsAffectedCount int64
	rowsAffectedCount, err := t.sql.CreateData(data)
	if err != nil {
		logger.Log.Error().Msgf("Create Data error: %s", err.Error())
		return rowsAffectedCount, err
	}

	return rowsAffectedCount, nil
}

func (t *TaskManager) UpsertData(data interface{}) (int64, error) {
	logger.Log.Debug().Msgf("Upsert Data :", data)

	rowsAffectedCount, err := t.sql.UpsertData(data)
	if err != nil {
		logger.Log.Error().Msgf("Upsert Data error: %s", err.Error())
		return rowsAffectedCount, err
	}

	return rowsAffectedCount, nil
}

func (t *TaskManager) GetCountCustomQuery(table interface{}, query map[string]interface{}) (int64, error) {
	count, err := t.sql.GetCountCustomQuery(table, query)
	if err != nil {
		logger.Log.Error().Msgf("Get Count Error: %s", err.Error())
		return count, err
	}
	return count, nil
}

func (t *TaskManager) GetDataList(input interface{}) (interface{}, error) {
	list, err := t.sql.GetDataList(input)
	if err != nil {
		logger.Log.Error().Msgf("Get Data List Error: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (t *TaskManager) GetData(name string, info interface{}) (interface{}, error) {
	info, err := t.sql.GetData(name, info)
	if err != nil {
		logger.Log.Error().Msgf("Get Data Error: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (t *TaskManager) DeleteCustomQuery(query map[string]interface{}, data interface{}) (int64, error) {
	logger.Log.Debug().Msgf("Delete Custom Data :", data)

	rowsAffected, err := t.sql.DeleteCustomQuery(query, data)
	if err != nil {
		logger.Log.Error().Msgf("Delete Custom Data error: %s", err.Error())
		return rowsAffected, err
	}

	return rowsAffected, nil
}

func (t *TaskManager) GetIndexFromTable(tableName string, idList, providerList []string) ([]int, error) {
	ids, err := t.sql.GetIndexFromTable(tableName, idList, providerList)
	if err != nil {
		logger.Log.Error().Msgf("Get provider From Table error: %s", err.Error())
		return ids, err
	}
	return ids, err
}

func (t *TaskManager) GetCountIdProviderQuery(tableName string, id, provider string) (int64, error) {
	dataCount, err := t.sql.GetCountIdProviderQuery(tableName, id, provider)
	if err != nil {
		logger.Log.Error().Msgf("Get provider, id From table error: %s", err.Error())
		return dataCount, err
	}
	return dataCount, err
}

//func (t *TaskManager) GetIdList(table interface{}) (interface{}, error) {
//	data, err := t.sql.GetIdList(table)
//	if err != nil {
//		logger.Log.Error().Msgf("Get Id Error: %s", err.Error())
//		return nil, err
//	}
//	return data, nil
//}
