package model

import "time"

type SsCarrierAccount struct {
	ID            int64     `gorm:"column:id" json:"id"`
	CarrierInfoId int64     `gorm:"column:carrier_info_id" json:"carrierInfoId" description:"关联承运商信息Id"`
	Enabled       int64     `gorm:"column:enabled" json:"enabled"`
	CreatedTime   JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId     int64     `gorm:"column:creator_id" json:"creatorId"`
	UpdatedTime   JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId     int64     `gorm:"column:updator_id" json:"updatorId"`
	Remarks       string    `gorm:"column:remarks" json:"remarks"`
}

func (ca SsCarrierAccount) TableName() string {
	return "ss_carrier_account"
}

type SsCarrierInfo struct {
	ID                       int64     `gorm:"column:id" json:"id"`
	Name                     string    `gorm:"column:name" json:"name" description:"账号名称"`
	CompanyName              string    `gorm:"column:company_name" json:"companyName" description:"企业名称"`
	ContactPerson            string    `gorm:"column:contact_person" json:"contactPerson" description:"联系人姓名"`
	ContactPhone             string    `gorm:"column:contact_phone" json:"contactPhone" description:"联系人电话"`
	ContactEmail             string    `gorm:"column:contact_email" json:"contactEmail" description:"联系人邮箱"`
	LegalPersonName          string    `gorm:"column:legal_person_name" json:"legalPersonName" description:"法人名称"`
	LegalPersonIdCard        string    `gorm:"column:legal_person_id_card" json:"legalPersonIdCard" description:"法人身份证号"`
	LegalPersonPhone         string    `gorm:"column:legal_person_phone" json:"legalPersonPhone" description:"法人联系电话"`
	LegalPersonEmail         string    `gorm:"column:legal_person_email" json:"legalPersonEmail" description:"法人联系邮箱"`
	IdCardPositivePicUrl     string    `gorm:"column:id_card_positive_pic_url" json:"idCardPositivePicUrl" description:"身份证正面"`
	IdCardInsteadPicUrl      string    `gorm:"column:id_card_instead_pic_url" json:"idCardInsteadPicUrl"`
	IdCardHeadPositivePicUrl string    `gorm:"column:id_card_head_positive_pic_url" json:"idCardHeadPositivePicUrl" description:"手持身份证"`
	BusinessLicensePicUrl    string    `gorm:"column:business_license_pic_url" json:"businessLicensePicUrl" description:"营业执照url"`
	OrganizationPicUrl       string    `gorm:"column:organization_pic_url" json:"organizationPicUrl" description:"组织机构照url"`
	CreatedTime              time.Time `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId                int64     `gorm:"column:creator_id" json:"creatorId"`
	UpdatedTime              time.Time `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId                int64     `gorm:"column:updator_id" json:"updatorId"`
	State                    int64     `gorm:"column:state" json:"state" description:"状态(-1欠费 0正常)"`
	Enabled                  int64     `gorm:"column:enabled" json:"enabled"`
	Remarks                  string    `gorm:"column:remarks" json:"remarks"`
}

func (ci SsCarrierInfo) TableName() string {
	return "ss_carrier_info"
}
