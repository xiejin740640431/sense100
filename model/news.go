package model

type SsSpNews struct {
	ID          int64  `gorm:"column:id" json:"id"`
	ProgramId   int64  `gorm:"column:program_id" json:"programId" description:"关联小程序Id"`
	Title       string `gorm:"column:title" json:"title" description:"标题"`
	Description string `gorm:"column:description" json:"description" description:"描述"`
	PictureUrl  string `gorm:"column:picture_url" json:"pictureUrl" description:"图片url"`
	RichText    string `gorm:"column:rich_text" json:"richText" description:"富文本内容"`
	SourceUrl   string `gorm:"column:source_url" json:"sourceUrl" description:"原网页url"`
	Enabled     int64  `gorm:"column:enabled" json:"enabled"`
	Remarks     string `gorm:"column:remarks" json:"remarks"`
}

func (news SsSpNews) TableName() string {
	return "ss_sp_news"
}
