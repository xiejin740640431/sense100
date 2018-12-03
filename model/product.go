package model

type SsSpProductCategory struct {
	ID          int64     `gorm:"column:id" json:"id"`
	ParentId    int64     `gorm:"column:parent_id" json:"parentId" description:"关联父级Id"`
	ProgramId   int64     `gorm:"column:program_id" json:"programId" description:"小程序Id"`
	Name        string    `gorm:"column:name" json:"name" description:"产品名称"`
	IconUrl     string    `gorm:"column:icon_url" json:"iconUrl" description:"iconUrl"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId   int64     `gorm:"column:creator_id" json:"creatorId"`
	UpdatedTime JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId   int64     `gorm:"column:updator_id" json:"updatorId"`
	Enabled     int64     `gorm:"column:enabled" json:"enabled"`
	Remarks     string    `gorm:"column:remarks" json:"remarks"`
}

func (category SsSpProductCategory) TableName() string {
	return "ss_sp_product_category"
}

type SsSpProductInfo struct {
	ID            int64     `gorm:"column:id" json:"id"`
	ProgramId     int64     `gorm:"column:program_id" json:"programId" description:"关联小程序Id"`
	CategoryId    int64     `gorm:"column:category_id" json:"categoryId" description:"关联类型Id"`
	Title         string    `gorm:"column:title" json:"title" description:"标题"`
	SalePrice     float64   `gorm:"column:sale_price" json:"salePrice" description:"售价"`
	MarketPrice   float64   `gorm:"column:market_price" json:"marketPrice" description:"市场价"`
	SkuPicUrls    string    `gorm:"column:sku_pic_urls" json:"skuPicUrls" description:"sku图"`
	VideoUrl      string    `gorm:"column:video_url" json:"videoUrl" description:"视频地址"`
	PictureUrl    string    `gorm:"column:picture_url" json:"pictureUrl" description:"默认/列表图片地址"`
	Description   string    `gorm:"column:description" json:"description" description:"描述"`
	Status        int64     `gorm:"column:status" json:"status" description:"状态(1编辑、2上架、3下架)"`
	CreatedTime   JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId     int64     `gorm:"column:creator_id" json:"creatorId"`
	UpdatedTime   JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId     int64     `gorm:"column:updator_id" json:"updatorId"`
	Enabled       int64     `gorm:"column:enabled" json:"enabled"`
	DetailPicUrls string    `gorm:"column:detail_pic_urls" json:"detailPicUrls" description:"详情图url，逗号分割"`
	Remarks       string    `gorm:"column:remarks" json:"remarks"`
}

func (productInfo SsSpProductInfo) TableName() string {
	return "ss_sp_product_info"
}
