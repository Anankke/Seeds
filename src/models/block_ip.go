package models

type BlockIp struct {
	Id       int64  `gorm:"column:id;NOT NULL;PRIMARY KEY;" json:"id"`
	NodeId   int    `gorm:"column:nodeid;NOT NULL;" json:"nodeid"`
	Ip       string `gorm:"column:ip;NOT NULL;type:text;" json:"ip"`
	Datetime int64  `gorm:"column:datetime;NOT NULL;" json:"datetime"`
}

func (BlockIp) TableName() string {
	return "blockip"
}
