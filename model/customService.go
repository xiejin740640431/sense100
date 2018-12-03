package model

type SsSpCustomService struct {
	ID          int64     `gorm:"column:id" json:"id"`
	ProgramId   int64     `gorm:"column:program_id" json:"programId" description:"关联小程序Id" required:"true"`
	CompanyId   int64     `gorm:"column:company_id" json:"companyId" description:"关联企业Id"`
	Name        string    `gorm:"column:name" json:"name"  description:"客服名称" required:"true"`
	Type        int64     `gorm:"column:type" json:"type" description:"客服类型(1电话、2QQ、3微信)" required:"true"`
	Number      string    `gorm:"column:number" json:"number"  description:"客服号(电话/QQ/微信)" required:"true"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId   int64     `gorm:"column:creator_id" json:"creatorId" required:"true"`
	UpdatedTime JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId   int64     `gorm:"column:updator_id" json:"updatorId"`
	Enabled     int64     `gorm:"column:enabled" json:"enabled"`
	Remarks     string    `gorm:"column:remarks" json:"remarks"`
}

func (customService SsSpCustomService) TableName() string {
	return "ss_sp_custom_service"
}
