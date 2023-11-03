package server

type OjtInfo struct {
	Id      string `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	Balance string `gorm:"column:balance"`
}

func (OjtInfo) TableName() string {
	return "ojt_data_tbls"
}
