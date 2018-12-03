package model

type SsSpLeavingMessage struct {
	ID          int64  `gorm:"column:id" json:"id"`
	ProgramId   int64  `gorm:"column:program_id" json:"programId"`
	Name        string `gorm:"column:name" json:"name"`
	Telephone   string `gorm:"column:telephone" json:"telephone"`
	Content     string `gorm:"column:content" json:"content"`
	PictureUrls string `gorm:"column:picture_urls" json:"pictureUrls"`
	Remarks     string `gorm:"column:remarks" json:"remarks"`
}

func (msg SsSpLeavingMessage) TableName() string {
	return "ss_sp_leaving_message"
}
