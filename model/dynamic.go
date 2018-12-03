package model

import (
	"sense100/constant"
	"sense100/util"
)

type SsSpDynamic struct {
	ID          int64        `gorm:"column:id" json:"id"`
	UserId      int64        `gorm:"user_id" json:"userId" binding:"required" description:"用户Id"`
	ProgramId   int64        `gorm:"column:program_id" json:"programId" description:"关联小程序Id"`
	CompanyId   int64        `gorm:"column:company_id" json:"companyId" description:"关联企业Id"`
	FacadeId    int64        `gorm:"column:facade_id" json:"facadeId" description:"关联门面Id"`
	Type        int64        `gorm:"column:type" json:"type" description:"动态类型(1企业门面、2小程序、3企业)"`
	Content     string       `gorm:"column:content" json:"content"  description:"内容"`
	PictureUrls string       `gorm:"column:picture_urls" json:"pictureUrls" description:"图片url,逗号分割"`
	State       int64        `gorm:"column:state" json:"state" description:"状态(1编辑中、2已发布、3删除)"`
	Enabled     int64        `gorm:"column:enabled" json:"enabled"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId   int64        `gorm:"column:creator_id" json:"creatorId"`
	DeleteTime  JsonDateTime `gorm:"not null;default:NULL;column:delete_time" json:"deleteTime"`
	DeletorId   int64        `gorm:"column:deletor_id" json:"deletorId"`
	Remarks     string       `gorm:"column:remarks" json:"remarks"`
}

func (d SsSpDynamic) TableName() string {
	return "ss_sp_dynamic"
}

func (d SsSpDynamic) CreateCheck() bool {
	if d.UserId == 0 {
		return false
	}
	if util.IsEmpty(d.Content) && util.IsEmpty(d.PictureUrls) {
		return false
	}
	if d.Type == constant.DynamicFacade && d.FacadeId != 0 {
		return true
	}
	if d.Type == constant.DynamicSmallProgram && d.ProgramId != 0 {
		return true
	}
	if d.Type == constant.DynamicCompany && d.CompanyId != 0 {
		return true
	}
	return false
}

type DynamicResp struct {
	SsSpDynamic
	PraiseNum int  `gorm:"praise_num" json:"praiseNum"`
	IsPraise  bool `gorm:"is_praise" json:"isPraise"`
}
