package models

type UnblockIp struct {
	Id       int64  `gorm:"column:id;NOT NULL;PRIMARY KEY;";json:"id";`
	Ip       string `gorm:"column:ip;NOT NULL;type:text;";json:"ip";`
	Datetime int64  `gorm:"column:datetime;NOT NULL;";json:"datetime";`
	Userid   int64  `gorm:"column:userid;NOT NULL;";json:"userid";`
}

func (UnblockIp) TableName() string {
	return "unblockip"
}
