package datastore

import (
	"errors"
	"httpproject/util/logger"
)

type DBConnector interface {
	// 기본 정보
	CreateData(input interface{}) (int64, error)
	UpsertData(input interface{}) (int64, error)
	GetCountCustomQuery(table interface{}, query map[string]interface{}) (int64, error)
	GetDataList(input interface{}) (interface{}, error)
	GetData(name string, info interface{}) (interface{}, error)
	DeleteCustomQuery(query map[string]interface{}, input interface{}) (int64, error)
	GetIndexFromTable(tableName string, idList, providerList []string) ([]int, error)
	GetCountIdProviderQuery(tableName string, id, provider string) (int64, error)
}

func (d *DataStore) CreateData(input interface{}) (int64, error) {
	var rowsAffectedCount int64
	var err error
	if d.db == nil {
		return rowsAffectedCount, errors.New("datastore connection failed")
	}
	if rowsAffectedCount, err = d.db.CreateData(input); err != nil {
		return rowsAffectedCount, err
	}
	return rowsAffectedCount, nil
}

func (d *DataStore) UpsertData(input interface{}) (int64, error) {
	var rowsAffectedCount int64
	var err error
	if d.db == nil {
		return rowsAffectedCount, errors.New("datastore connection failed")
	}

	if rowsAffectedCount, err = d.db.UpsertData(input); err != nil {
		return rowsAffectedCount, err
	}

	return rowsAffectedCount, nil
}

func (d *DataStore) GetCountCustomQuery(table interface{}, query map[string]interface{}) (count int64, err error) {
	if d.db == nil {
		return 0, errors.New("datastore connection failed")
	}

	count, err = d.db.GetCountCustomQuery(table, query)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (d *DataStore) GetDataList(input interface{}) (list interface{}, err error) {
	if d.db == nil {
		return nil, errors.New("datastore connection failed")
	}

	list, err = d.db.GetDataList(input)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DataStore) GetData(name string, info interface{}) (interface{}, error) {
	data := make([]byte, 0)
	if d.db == nil {
		return data, errors.New("datastore connection failed")
	}

	logger.Log.Debug().Msgf("name :", name)
	info, err := d.db.GetData(name, info)
	logger.Log.Debug().Msgf("GetClusterInfoByName :", info)
	if err != nil {
		return data, err
	}

	return info, nil
}

func (d *DataStore) DeleteCustomQuery(query map[string]interface{}, input interface{}) (int64, error) {
	var rowsAffectedCount int64
	if d.db == nil {
		return rowsAffectedCount, errors.New("datastore connection failed")
	}
	var err error
	if rowsAffectedCount, err = d.db.DeleteCustomQuery(query, input); err != nil {
		return rowsAffectedCount, err
	}
	return rowsAffectedCount, nil
}

func (d *DataStore) CheckDBConnection() bool {
	result := true
	if d.db == nil {
		result = false
	}
	return result
}

func (d *DataStore) GetIndexFromTable(tableName string, idList, providerList []string) ([]int, error) {
	var indexList []int
	if d.db == nil {
		return indexList, errors.New("datastore connection failed")
	}
	var err error
	if indexList, err = d.db.GetIndexFromTable(tableName, idList, providerList); err != nil {
		return indexList, err
	}
	return indexList, err
}

func (d *DataStore) GetCountIdProviderQuery(tableName string, id, provider string) (int64, error) {
	var dataCount int64
	if d.db == nil {
		return dataCount, errors.New("datastore connection failed")
	}
	var err error
	if dataCount, err = d.db.GetCountIdProviderQuery(tableName, id, provider); err != nil {
		return dataCount, err
	}
	return dataCount, err
}
