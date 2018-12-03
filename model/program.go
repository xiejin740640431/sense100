package model

type SsSmallProgram struct {
	ID                        int64     `gorm:"column:id" json:"id"`
	CompanyId                 int64     `gorm:"column:company_id" json:"company_id" description:"企业Id"`
	UserId                    int64     `gorm:"column:user_id" json:"user_id" description:"用户Id"`
	MouldId                   int64     `gorm:"column:mould_id" json:"mould_id" description:"模版Id"`
	ProgramName               string    `gorm:"column:program_name" json:"program_name" description:"程序名称"`
	SourceUrl                 string    `gorm:"column:source_url" json:"sourceUrl"  description:"源码地址"`
	Url                       string    `gorm:"column:url" json:"url" description:"项目地址"`
	QrCode                    string    `gorm:"column:qr_code" json:"qr_code" description:"二维码字符串"`
	ProjectStartupTime        JsonDateTime `gorm:"not null;default:NULL;column:project_startup_time" json:"project_startup_time" description:"项目启动时间"`
	ExpectProjectFinishedTime JsonDateTime `gorm:"not null;default:NULL;column:expect_project_finished_time" json:"expect_project_finished_time" description:"预期项目结束时间"`
	ProjectFinishedTime       JsonDateTime `gorm:"not null;default:NULL;column:project_finished_time" json:"project_finished_time" description:"项目结束时间"`
	LogoUrl                   string    `gorm:"column:logo_url" json:"logo_url" description:"loglUrl"`
	DetailPictureUrls         string    `gorm:"column:detail_picture_urls" json:"detail_picture_urls" description:"详情图url，逗号分割"`
	ProgramData               string    `gorm:"column:program_data" json:"program_data" description:"程序自定义数据"`
	SpAppId                   string    `gorm:"column:sp_app_id" json:"sp_app_id" description:"appId"`
	SpAppSecret               string    `gorm:"column:sp_app_secret" json:"sp_app_secret" description:"appSecret"`
	StartUseTime              JsonDateTime `gorm:"not null;default:NULL;column:start_use_time" json:"start_use_time" description:"开通时间"`
	ExpiryTime                JsonDateTime `gorm:"not null;default:NULL;column:expiry_time" json:"expiry_time" description:"过期时间"`
	RentDays                  int64     `gorm:"column:rent_days" json:"rent_days" description:"购买天数"`
	Remarks                   string    `gorm:"column:remarks" json:"remarks" description:"备注"`
}

func (program SsSmallProgram) TableName() string {
	return "ss_small_program"
}

type SsSmallProgramOrderRelation struct {
	ID             int64 `gorm:"primary;column:id" json:"id"`
	SmallProgramId int64 `gorm:"small_program_id" json:"smallProgramId"`
	OrderInfoId    int64 `gorm:"order_info_id" json:"orderInfoId"`
}

func (spoRelation SsSmallProgramOrderRelation) TableName() string {
	return "ss_small_program_order_relation"
}
