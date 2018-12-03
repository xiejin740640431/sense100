package model

type SsShareRecord struct {
	ID           int64     `gorm:"column:id" json:"id"`
	UserId       string    `gorm:"column:user_id" json:"userId" description:"关联用户Id"`
	ShareContent int64     `gorm:"column:share_content" json:"shareContent" description:"分享内容(1模版)"`
	CreatedTime  JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	Remarks      string    `gorm:"column:remarks" json:"remarks"`
}

func (sr SsShareRecord) TableName() string {
	return "ss_share_record"
}

type SsSpShareRecord struct {
	ID             int64     `gorm:"column:id" json:"id"`
	ProgramId      int64     `gorm:"column:program_id" json:"programId" description:"关联小程序Id"`
	WxOpenId       string    `gorm:"column:wx_open_id" json:"wx_openId" description:"微信Id"`
	ShareContentId int64     `gorm:"column:share_content_id" json:"shareContentId" description:"分享内容的Id"`
	ShareContent   int64     `gorm:"column:share_content" json:"shareContent" description:"分享的内容(1产品、2新闻)"`
	CreatedTime    JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	Remarks        string    `gorm:"column:remarks" json:"remarks"`
}

func (sr SsSpShareRecord) TableName() string {
	return "ss_sp_share_record"
}
