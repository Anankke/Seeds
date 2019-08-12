package models

type SpeedTest struct {
	Id               int64  `gorm:"column:id;NOT NULL;PRIMARY KEY;" json:"id";`
	NodeId           int    `gorm:"column:nodeid;NOT NULL;" json:"nodeid"`
	DateTime         int64  `gorm:"column:datetime;NOT NULL;" json:"datetime"`
	TelecomePing     string `gorm:"column:telecomping;NOT NULL;type:text;" json:"telecomping";`
	TelecomeUpload   string `gorm:"column:telecomeupload;NOT NULL;type:text;" json:"telecomeupload";`
	TelecomeDownload string `gorm:"column:telecomedownload;NOT NULL;type:text;" json:"telecomedownload";`
	UnicomPing       string `gorm:"column:unicomping;NOT NULL;type:text;" json:"unicomping";`
	UnicomUpload     string `gorm:"column:unicomupload;NOT NULL;type:text;" json:"unicomupload";`
	UnicomDownload   string `gorm:"column:unicomdownload;NOT NULL;type:text;" json:"unicomdownload";`
	CmccPing         string `gorm:"column:cmccping;NOT NULL;type:text;" json:"cmccping";`
	CmccUpload       string `gorm:"column:cmccupload;NOT NULL;type:text;" json:"cmccupload";`
	CmccDownload     string `gorm:"column:cmccdownload;NOT NULL;type:text;" json:"cmccdownload";`
}

func (SpeedTest) TableName() string {
	return "speedtest"
}
