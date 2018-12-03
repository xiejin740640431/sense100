package model

type SsSpBrowseRecord struct {
	ID          int64     `gorm:"column:id" json:"id"`
	TouristId   int64     `gorm:"column:tourist_id" json:"touristId" description:"游客Id" required:"true"`
	ContentId   int64     `gorm:"column:content_id" json:"contentId" description:"关联浏览内容Id，如：企业门面Id、小程序Id" required:"true"`
	BrowseNum   int       `gorm:"column:browse_num" json:"browseNum" description:"浏览次数"`
	Type        int       `gorm:"type" json:"type" description:"类型(1企业门面、2小程序)" required:"true"`
	CreatedTime JsonDateTime `gorm:"column:created_time" json:"createdTime" description:"创建时间"`
	UpdatedTime JsonDateTime `gorm:"column:updated_time" json:"updatedTime" description:"更新时间"`
	Enabled     int       `gorm:"column:enabled" json:"enabled" description:"0禁用、1启用"`
}

func (br SsSpBrowseRecord) TableName() string {
	return "ss_sp_browse_record"
}

type AddBrowseRecordReq struct {
	SsTourist
	ContentId int64 `gorm:"column:content_id" json:"contentId" description:"关联浏览内容Id，如：企业门面Id、小程序Id" required:"true"`
	Type      int   `gorm:"type" json:"type" description:"类型(1企业门面、2小程序)" required:"true"`
}
