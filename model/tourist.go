package model

type SsTourist struct {
	ID          int64     `gorm:"primary;column:id" json:"id" description:"Id"`
	WxId        string    `gorm:"column:wx_id" json:"wxId" description:"微信Id"`
	NickName    string    `gorm:"column:nickName" json:"nickName" description:"昵称"`
	HeadUrl     string    `gorm:"column:head_url" json:"headUrl" description:"头像url"`
	Mobile      string    `gorm:"column:mobile" json:"mobile" description:"手机号"`
	Address     string    `gorm:"column:address" json:"address" description:"地址"`
	Email       string    `gorm:"column:email" json:"email" description:"邮箱"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	Enabled     int64     `gorm:"column:enabled" json:"enabled"`
	Remarks     string    `gorm:"column:remarks" json:"remarks" description:"备注"`
}

func (t SsTourist) TableName() string {
	return "ss_tourist"
}
