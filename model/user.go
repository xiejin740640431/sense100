package model

type SsUserAccount struct {
	ID        int64       `gorm:"primary;column:id" json:"id" description:"用户Id"`
	UserId    int64       `gorm:"column:user_id" json:"userId" description:"关联用户信息Id"`
	CompanyId int64       `gorm:"column:company_id" json:"companyId" description:"关联企业Id"`
	UserName  string      `gorm:"column:user_name" json:"userName"  description:"用户名"`
	WxId      string      `gorm:"column:wx_id" json:"wxId" description:"微信Id"`
	Password  string      `gorm:"column:password" json:"password" description:"密码"`
	IpAddress string      `gorm:"column:ip_address" json:"ipAddress" description:"ip地址"`
	Enabled   int64       `gorm:"column:enabled" json:"enabled" description:"启用状态"`
	Remarks   string      `gorm:"column:remarks" json:"remarks" description:"备注"`
	UserInfo  *SsUserInfo `gorm:"ForeignKey:ID;AssociationForeignKey:UserId" json:"userInfo"`
}

func (userAccount SsUserAccount) TableName() string {
	return "ss_user_account"
}

type SsUserInfo struct {
	ID          int64        `gorm:"primary;column:id" json:"id" description:"Id"`
	RealName    string       `gorm:"column:real_name" json:"realName" description:"真实姓名"`
	NickName    string       `gorm:"column:nickName" json:"nickName" description:"昵称"`
	HeadUrl     string       `gorm:"column:head_url" json:"headUrl" description:"头像url"`
	Mobile      string       `gorm:"column:mobile" json:"mobile" description:"手机号"`
	Address     string       `gorm:"column:address" json:"address" description:"地址"`
	Email       string       `gorm:"column:email" json:"email" description:"邮箱"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	UpdatedTime JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	Enabled     int64        `gorm:"column:enabled" json:"enabled"`
	Remarks     string       `gorm:"column:remarks" json:"remarks" description:"备注"`
}

func (userInfo SsUserInfo) TableName() string {
	return "ss_user_info"
}
