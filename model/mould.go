package model

type SsMouldCategory struct {
	ID           int64     `gorm:"column:id" json:"id"`
	ParentId     int64     `gorm:"column:parent_id" json:"parentId" description:"关联父级Id"`
	CategoryName string    `gorm:"column:category_name" json:"categoryName" description:"类型名称" required:"true"`
	CreatedTime  JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	UpdatedTime  JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	CreatorId    int64     `gorm:"column:creator_id" json:"creatorId" description:"管理员Id"`
	UpdatorId    int64     `gorm:"column:updator_id" json:"updatorId" description:"管理员Id"`
	Enabled      int64     `gorm:"column:enabled" json:"enabled"`
	Remarks      string    `gorm:"column:remarks" json:"remarks"`
}

func (mc SsMouldCategory) TableName() string {
	return "ss_mould_category"
}

type SsProgramMould struct {
	ID                int64     `gorm:"column:id" json:"id"`
	CategoryId        int64     `gorm:"column:category_id" json:"categoryId" required:"true"`
	Name              string    `gorm:"column:name" json:"name" description:"名称" required:"true"`
	Description       string    `gorm:"column:description" json:"description" description:"描述" required:"true"`
	SalePrice         float64   `gorm:"column:sale_price" json:"salePrice" description:"售价" required:"true"`
	MarketPrice       float64   `gorm:"column:market_price" json:"marketPrice" description:"市场价"`
	Status            int64     `gorm:"column:status" json:"status" description:"状态(1开发中、2已完成)" required:"true"`
	Enabled           int64     `gorm:"column:enabled" json:"enabled"`
	CoverUrl          string    `gorm:"column:cover_url" json:"coverUrl" description:"封面url" required:"true"`
	DetailPictureUrls string    `gorm:"column:detail_picture_urls" json:"detailPictureUrls" description:"详情图url,逗号分割" required:"true"`
	SourceUrl         string    `gorm:"column:source_url" json:"sourceUrl" description:"源码url"`
	CreatedTime       JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId         int64     `gorm:"column:creator_id" json:"creatorId" description:"管理员Id"`
	UpdatedTime       JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId         int64     `gorm:"column:updator_id" json:"updatorId" description:"管理员Id"`
	Remarks           string    `gorm:"column:remarks" json:"remarks"`
}

func (pm SsProgramMould) TableName() string {
	return "ss_program_mould"
}
