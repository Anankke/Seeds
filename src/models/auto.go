package models

type Auto struct {
	Id       int64  `gorm:"column:id;NOT NULL;PRIMARY KEY;" json:"id";`
	Type     int    `gorm:"column:type;NOT NULL;" json:"type";`
	Value    string `gorm:"column:value;NOT NULL;type:longtext;" json:"value";`
	Sign     string `gorm:"column:sign;NOT NULL;type:longtext;" json:"sign";`
	Datetime int64  `gorm:"column:datetime;NOT NULL;" json:"datetime";`
}

func (Auto) TableName() string {
	return "auto"
}
