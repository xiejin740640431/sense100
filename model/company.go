package model

type SsCompanyAccount struct {
	ID            int64   `gorm:"column:id" json:"id"`
	CompanyInfoId int64   `gorm:"column:company_info_id" json:"companyInfoId"`
	Balance       float64 `gorm:"column:balance" json:"balance" description:"余额"`
	Enabled       int64   `gorm:"column:enabled" json:"enabled"`
	Remarks       string  `gorm:"column:remarks" json:"remarks"`
}

func (ca SsCompanyAccount) TableName() string {
	return "ss_company_account"
}

type SsCompanyInfo struct {
	ID               int64     `gorm:"column:id" json:"id"`
	TradeId          int64     `gorm:"column:trade_id" json:"tradeId" description:"关联行业类型Id"`
	Name             string    `gorm:"column:name" json:"name" description:"名称"`
	BannerUrls       string    `gorm:"column:banner_urls" json:"bannerUrls" description:"bannerUrl,逗号分割"`
	Description      string    `gorm:"column:description" json:"description" description:"描述"`
	OfficialNetwork  string    `gorm:"column:official_network" json:"officialNetwork" description:"官网链接"`
	Longitude        float64   `gorm:"column:longitude" json:"longitude" description:"经度"`
	Latitude         float64   `gorm:"column:latitude" json:"latitude" description:"纬度"`
	Address          string    `gorm:"column:address" json:"address" description:"企业地址"`
	Telephone        string    `gorm:"column:telephone" json:"telephone" description:"联系电话"`
	RealSceneUrls    string    `gorm:"column:real_scene_urls" json:"realSceneUrls" description:"企业实景图,逗号分割"`
	ConsultBeginTime JsonDateTime `gorm:"not null;default:NULL;column:consult_begin_time" json:"consultBeginTime"`
	ConsultEndTime   JsonDateTime `gorm:"not null;default:NULL;column:consult_end_time" json:"consultEndTime"`
	Remarks          string    `gorm:"column:remarks" json:"remarks"`
	Enabled          int64     `gorm:"column:enabled" json:"enabled"`
}

func (ci SsCompanyInfo) TableName() string {
	return "ss_company_info"
}

type SsCompanyFacadeInfo struct {
	ID                int64        `gorm:"column:id" json:"id"`
	CompanyId         int64        `gorm:"column:company_id" json:"companyId" description:"关联企业Id"`
	UserId            int64        `gorm:"column:user_id" json:"userId" description:"关联用户Id"`
	Name              string       `gorm:"column:name" json:"name" description:"名称"`
	WxAccount         string       `gorm:"column:wx_account" json:"wxAccount" description:"微信账号"`
	Telephone         string       `gorm:"column:telephone" json:"telephone" description:"联系电话"`
	TradeId           int64        `gorm:"column:trade_id" json:"tradeId" description:"关联行业类型Id"`
	BannerUrls        string       `gorm:"column:banner_urls" json:"bannerUrls" description:"bannerUrl,逗号分割"`
	BusinessBeginTime JsonTimeHS   `gorm:"not null;default:NULL;column:business_begin_time" json:"businessBeginTime" time_format:"sql_time" time_location:"Local" description:"营业开始时间"`
	BusinessEndTime   JsonTimeHS   `gorm:"not null;default:NULL;column:business_end_time" json:"businessEndTime" time_format:"sql_time" time_location:"Local" description:"营业结束时间"`
	Longitude         float64      `gorm:"column:longitude" json:"longitude" description:"经度"`
	Latitude          float64      `gorm:"column:latitude" json:"latitude" description:"纬度"`
	Address           string       `gorm:"column:address" json:"address" description:"企业地址"`
	VideoUrl          string       `gorm:"column:video_url" json:"videoUrl" description:"宣传视频url"`
	PictureUrls       string       `gorm:"column:picture_urls" json:"pictureUrls" description:"图片url"`
	StartUseTime      JsonDateTime `gorm:"not null;default:NULL;column:start_use_time" json:"startUseTime" description:"购买，开始使用时间"`
	ExpiryTime        JsonDateTime `gorm:"not null;default:NULL;column:expiry_time" json:"expiryTime" description:"过期时间"`
	RentDays          int64        `gorm:"column:rent_days" json:"rentDays" description:"购买天数"`
	VisitorNumber     int64        `gorm:"column:visitor_number" json:"visitorNumber" description:"访问量"`
	State             int64        `gorm:"column:state" json:"state" description:"状态(-1欠费、0正常)"`
	Enabled           int64        `gorm:"column:enabled" json:"enabled"`
	CreatedTime       JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId         int64        `gorm:"column:creator_id" json:"creatorId"`
	UpdatedTime       JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId         int64        `gorm:"column:updator_id" json:"updatorId"`
	Remarks           string       `gorm:"column:remarks" json:"remarks"`
}

func (f SsCompanyFacadeInfo) TableName() string {
	return "ss_company_facade_info"
}

type SsFacadeOrderRelation struct {
	ID              int64 `gorm:"primary;column:id" json:"id"`
	CompanyFacadeId int64 `gorm:"company_facade_id" json:"companyFacadeId"`
	OrderInfoId     int64 `gorm:"order_info_id" json:"orderInfoId"`
}

func (foRelation SsFacadeOrderRelation) TableName() string {
	return "ss_facade_order_relation"
}

type CompanyFacadeInfoResp struct {
	SsCompanyFacadeInfo
	BrowseNum int `json:"browseNum" description:"浏览人数"`
}
