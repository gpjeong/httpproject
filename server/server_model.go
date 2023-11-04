package server

type OjtInfo struct {
	Id      string `gorm:"primarykey;column:id"`
	Name    string `gorm:"column:name"`
	Balance string `gorm:"column:balance"`
}

func (OjtInfo) TableName() string {
	return "ojt_data_tbls"
}
