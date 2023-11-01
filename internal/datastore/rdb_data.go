package datastore

import (
	"fmt"
	"gorm.io/gorm/clause"
)

func (d *PostgreDB) CreateData(inputData interface{}) (int64, error) {
	aa := d.db.Create(inputData)
	return aa.RowsAffected, aa.Error
}

func (d *PostgreDB) UpsertData(inputData interface{}) (int64, error) {
	upsertResult := d.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(inputData)
	return upsertResult.RowsAffected, upsertResult.Error
}

func (d *PostgreDB) GetCountCustomQuery(table interface{}, query map[string]interface{}) (int64, error) {
	var count int64
	aa := d.db.Model(&table)
	for q, value := range query {
		if value != nil {
			aa.Where(q, value)

		} else {
			aa.Where(q)
		}
	}
	aa.Count(&count)
	return count, aa.Error
}

func (d *PostgreDB) GetDataList(list interface{}) (interface{}, error) {
	return list, d.db.Find(&list).Error
}

func (d *PostgreDB) GetData(id string, info interface{}) (interface{}, error) {
	data := d.db.Where(id).Find(&info)
	fmt.Println(data)
	return info, data.Error
}

func (d *PostgreDB) DeleteCustomQuery(query map[string]interface{}, data interface{}) (int64, error) {
	tx := d.db.Begin()
	aa := tx.Model(data)
	for q, value := range query {
		aa.Where(q, value)
	}

	var err error
	if deleteResult := aa.Delete(&data).Error; deleteResult != nil {
		tx.Rollback()
		err = deleteResult
	}
	tx.Commit()
	return aa.RowsAffected, err
}

func (d *PostgreDB) GetIndexFromTable(tableName string, idList, providerList []string) ([]int, error) {
	var indexList []int
	for index, id := range idList {
		var dataCount int64
		result := d.db.Table(tableName).
			Where("id = ? AND provider = ?", id, providerList[index]).
			Count(&dataCount)

		if result.Error != nil {
			return indexList, result.Error
		}

		// 데이터가 없을 때만 index 추가
		if dataCount == 0 {
			indexList = append(indexList, index)
		}
	}
	return indexList, nil
}

func (d *PostgreDB) GetCountIdProviderQuery(tableName string, id, provider string) (int64, error) {
	var dataCount int64
	result := d.db.Table(tableName).Where("id = ? AND provider = ?", id, provider).Count(&dataCount)
	if result.Error != nil {
		return dataCount, result.Error
	}
	return dataCount, nil

}
