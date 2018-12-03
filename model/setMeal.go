package model

type SsSetMeal struct {
	ID          int64           `gorm:"column:id" json:"id"`
	CarrierId   int64           `gorm:"column:carrier_id" json:"carrierId" description:"承运商Id" required:"true"`
	MouldId     int64           `gorm:"column:mould_id" json:"mouldId" description:"模版Id" required:"true"`
	Name        string          `gorm:"column:name" json:"name" description:"套餐名称" required:"true"`
	Description string          `gorm:"column:description" json:"description" description:"套餐描述"`
	CoverUrl    string          `gorm:"column:cover_url" json:"coverUrl" description:"封面url" required:"true"`
	CreatedTime JsonDateTime       `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId   int64           `gorm:"column:creator_id" json:"creatorId" required:"true"`
	UpdatedTime JsonDateTime       `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId   int64           `gorm:"column:updator_id" json:"updatorId"`
	Enabled     int64           `gorm:"column:enabled" json:"enabled"`
	Remarks     string          `gorm:"column:remarks" json:"remarks"`
	MealRoles   []SsSetMealRole `gorm:"ForeignKey:SetMealId;AssociationForeignKey:ID" json:"mealRoles"`
}

func (setMeal SsSetMeal) TableName() string {
	return "ss_set_meal"
}

type SsSetMealRole struct {
	ID            int64     `gorm:"column:id" json:"id"`
	SetMealId     int64     `gorm:"column:set_meal_id" json:"setMealId" description:"关联套餐Id" required:"true"`
	Name          string    `gorm:"column:name" json:"name" description:"名称" required:"true"`
	Description   string    `gorm:"column:description" json:"description"  description:"描述"`
	OriginalPrice float64   `gorm:"column:original_price" json:"originalPrice" description:"原价"`
	SalePrice     float64   `gorm:"column:sale_price" json:"salePrice" description:"销售价" required:"true"`
	PictureUrl    string    `gorm:"column:picture_url" json:"pictureUrl" description:"图片url" required:"true"`
	CreatedTime   JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId     int64     `gorm:"column:creator_id" json:"creatorId" required:"true"`
	UpdatedTime   JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId     int64     `gorm:"column:updator_id" json:"updatorId"`
	Enabled       int64     `gorm:"column:enabled" json:"enabled"`
	Remarks       string    `gorm:"column:remarks" json:"remarks"`
}

func (role SsSetMealRole) TableName() string {
	return "ss_set_meal_role"
}
