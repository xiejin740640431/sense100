package model

type SsSpMenu struct {
	ID          int64     `gorm:"column:id" json:"id"`
	ProgramId   int64     `gorm:"column:program_id" json:"programId" description:"关联小程序Id"`
	Name        string    `gorm:"column:name" json:"name" description:"名称"`
	IconUrl     string    `gorm:"column:icon_url" json:"iconUrl" description:"iconUrl"`
	WebLink     string    `gorm:"column:web_link" json:"webLink" description:"网页链接"`
	Sort        int64     `gorm:"column:sort" json:"sort" description:"排序"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId   int64     `gorm:"column:creator_id" json:"creatorId"`
	UpdatedTime JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId   int64     `gorm:"column:updator_id" json:"updatorId"`
	Enabled     int64     `gorm:"column:enabled" json:"enabled"`
	Remarks     string    `gorm:"column:remarks" json:"remarks"`
	Type        int64     `gorm:"column:type" json:"type" description:"菜单类型(1导航菜单、2底部菜单)"`
}

func (menu SsSpMenu) TableName() string {
	return "ss_sp_menu"
}
