package model

/*
 行业类型
 */
type SsTrade struct {
	ID          int64     `gorm:"column:id" json:"id"`
	ParentId    int64     `gorm:"column:parent_id" json:"parentId" description:"关联父级id" required:"true"`
	Name        string    `gorm:"column:name" json:"name" description:"名称" required:"true"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId   int64     `gorm:"column:creator_id" json:"creatorId" description:"管理员Id" required:"true"`
	UpdatedTime JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId   int64     `gorm:"column:updator_id" json:"updatorId" description:"管理员Id"`
	Enabled     int64     `gorm:"column:enabled" json:"enabled"`
	Remarks     string    `gorm:"column:remarks" json:"remarks"`
}

func (trade SsTrade) TableName() string {
	return "ss_trade"
}
