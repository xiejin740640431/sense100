package model

type SsAdmin struct {
	ID          int64     `gorm:"column:id" json:"id" description:"id"`
	UserName    string     `gorm:"column:user_name" json:"userName" description:"用户名" required:"true"`
	Password    string     `gorm:"column:password" json:"password" description:"密码"`
	RealName    string     `gorm:"column:real_name" json:"realName" description:"真实姓名"`
	Mobile      string     `gorm:"column:mobile" json:"mobile" description:"手机号"`
	Description string     `gorm:"column:description" json:"description" description:"描述"`
	IsSys       bool      `gorm:"column:is_sys" json:"isSys" description:"是否是系统管理员"`
	LoginTime   JsonDateTime`gorm:"column:login_time" json:"loginTime" description:"最近登录时间"`
	CreatedTime JsonDateTime `gorm:"column:created_time" json:"createdTime" description:"创建时间"`
	UpdatedTime JsonDateTime `gorm:"column:updated_time" json:"updatedTime" description:"更新时间"`
	Enabled     bool      `gorm:"column:enabled" json:"enabled" description:"是否启用(0禁用、1启用)"`
}

func (admin SsAdmin) TableName() string {
	return "ss_admin"
}
