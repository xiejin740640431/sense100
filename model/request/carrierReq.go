package request

import "sense100/util"

type CreateOrUpdateCarrierReq struct {
	ID                       int64  `gorm:"column:id" json:"id"`
	AdminId                  int64  `json:"adminId" description:"管理员Id" required:"true"`
	Name                     string `gorm:"column:name" json:"name" description:"名称" required:"true"`
	CompanyName              string `gorm:"column:company_name" json:"companyName" description:"企业名称" required:"true"`
	ContactPerson            string `gorm:"column:contact_person" json:"contactPerson" description:"联系人姓名" required:"true"`
	ContactPhone             string `gorm:"column:contact_phone" json:"contactPhone" description:"联系电话" required:"true"`
	ContactEmail             string `gorm:"column:contact_email" json:"contactEmail" description:"联系邮箱"`
	LegalPersonName          string `gorm:"column:legal_person_name" json:"legalPersonName" description:"法人姓名" required:"true"`
	LegalPersonIdCard        string `gorm:"column:legal_person_id_card" json:"legalPersonIdCard" description:"法人身份证" required:"true"`
	LegalPersonPhone         string `gorm:"column:legal_person_phone" json:"legalPersonPhone" description:"法人联系电话" required:"true"`
	LegalPersonEmail         string `gorm:"column:legal_person_email" json:"legalPersonEmail" description:"法人联系邮箱" required:"true"`
	IdCardPositivePicUrl     string `gorm:"column:id_card_positive_pic_url" json:"idCardPositivePicUrl" description:"身份证正面照" required:"true"`
	IdCardInsteadPicUrl      string `gorm:"column:id_card_instead_pic_url" json:"idCardInsteadPicUrl" description:"身份证" required:"true"`
	IdCardHeadPositivePicUrl string `gorm:"column:id_card_head_positive_pic_url" json:"idCardHeadPositivePicUrl" required:"true"`
	BusinessLicensePicUrl    string `gorm:"column:business_license_pic_url" json:"businessLicensePicUrl" required:"true"`
	OrganizationPicUrl       string `gorm:"column:organization_pic_url" json:"organizationPicUrl" required:"true"`
}

func (req CreateOrUpdateCarrierReq) Check() bool {
	if util.IsEmpty(req.Name) || util.IsEmpty(req.CompanyName) {
		return false
	}

	return true
}
