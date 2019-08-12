package models

type Relay struct {
	Id           int64  `gorm:"column:id;NOT NULL;PRIMARY KEY;" json:"id";`
	UserId       int64  `gorm:"column:user_id;NOT NULL;" json:"user_id";`
	SourceNodeId int64  `gorm:"column:source_node_id;NOT NULL;" json:"source_node_id";`
	DistNodeId   int64  `gorm:"column:dist_node_id;NOT NULL;" json:"dist_node_id";`
	DistIp       string `gorm:"column:dist_ip;NOT NULL;type:text;" json:"dist_ip";`
	Port         int    `gorm:"column:port;NOT NULL;" json:"port";`
	Priority     int    `gorm:"column:priority;NOT NULL;" json:"priority";`
}

func (Relay) TableName() string {
	return "relay"
}
