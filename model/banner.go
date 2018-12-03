package model

type SsSpBanner struct {
	ID          int64     `gorm:"column:id" json:"id"`
	ProgramId   int64     `gorm:"column:program_id" json:"programId" description:"关联小程序Id" required:"true"`
	PictureUrl  string    `gorm:"column:picture_url" json:"pictureUrl" description:"图片url" required:"true"`
	WebLink     string    `gorm:"column:web_link" json:"webLink" description:"web链接" required:"true"`
	Description string    `gorm:"column:description" json:"description" description:"描述" required:"true"`
	Enabled     int64     `gorm:"column:enabled" json:"enabled" description:"是否启用(0禁用、1启用)"`
	CreatorId   int64     `gorm:"column:creator_id" json:"creatorId" required:"true"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	UpdatorId   int64     `gorm:"column:updator_id" json:"updatorId"`
	UpdatedTime JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	Remarks     string    `gorm:"column:remarks" json:"remarks"`
}

func (sb SsSpBanner) TableName() string {
	return "ss_sp_banner"
}
