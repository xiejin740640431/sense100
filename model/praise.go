package model

type SsSpPraise struct {
	ID              int64     `gorm:"column:id" json:"id"`
	UserId          int64     `gorm:"column:user_id" json:"userId" description:"关联用户Id"`
	PraiseContentId int64     `gorm:"column:praise_content_id" json:"praiseContentId" description:"关联点赞内容Id,如:关联点赞动态Id"`
	CreatedTime     JsonDateTime `gorm:"column:created_time" json:"createdTime" description:"创建时间"`
	Type            int       `gorm:"column:type" json:"type" description:"类型(1点赞动态、2点赞资讯)"`
	Enabled         int       `gorm:"column:enabled" json:"enabled" description:"0禁用、1启用"`
}

func (p SsSpPraise) TableName() string {
	return "ss_sp_praise"
}
