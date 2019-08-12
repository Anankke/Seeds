package models

type DetectList struct {
	Id    int64  `gorm:"column:id;NOT NULL;PRIMARY KEY;" json:"id";`
	Name  string `gorm:"column:name;NOT NULL;type:longtext;" json:"name";`
	Text  string `gorm:"column:text;NOT NULL;type:longtext;" json:"text";`
	Regex string `gorm:"column:regex;NOT NULL;type:longtext;" json:"regex";`
	Type  int    `gorm:"column:type;NOT NULL;" json:"type";`
}

func (DetectList) TableName() string {
	return "detect_list"
}
