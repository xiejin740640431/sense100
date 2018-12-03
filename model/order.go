package model

/*
 订单信息表
 */
type SsOrderInfo struct {
	ID                 int64     `gorm:"column:id" json:"id"`
	UserId             int64     `gorm:"column:user_id" json:"userId" description:"关联用户Id"`
	OrderDetailId      int64     `gorm:"column:order_detail_id" json:"orderDetailId" description:"关联订单明细Id"`
	TotalAmount        float64   `gorm:"column:total_amount" json:"totalAmount" description:"总额"`
	PaidAmount         float64   `gorm:"column:paid_amount" json:"paidAmount" description:"已付款金额"`
	TransactionContent int64     `gorm:"column:transaction_content" json:"transactionContent" description:"交易内容(1购买企业门面、2购买模版、3定制)"`
	State              int64     `gorm:"column:state" json:"state" description:"状态(1待支付、2已付部分\定金、3全额付款、4已取消)"`
	PayTypeId          int64     `gorm:"column:pay_type_id" json:"payTypeId" description:"支付类型Id"`
	PayState           int64     `gorm:"column:pay_state" json:"payState" description:"支付状态（1待支付 2已支付)"`
	OrderSource        int64     `gorm:"column:order_source" json:"orderSource" description:"订单来源（1官网、2微信小程序）"`
	CreatedTime        JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId          int64     `gorm:"column:creatorId" json:"creatorId"`
	UpdatedTime        JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId          int64     `gorm:"column:updator_id" json:"updatorId"`
	Enabled            int64     `gorm:"column:enabled" json:"enabled"`
	Remarks            string    `gorm:"column:remarks" json:"remarks"`
}

func (orderInfo SsOrderInfo) TableName() string {
	return "ss_order_info"
}

/*
 订单明细表
 */
type SsOrderDetail struct {
	ID                  int64     `gorm:"column:id" json:"id"`
	SetMealRoleId       int64     `gorm:"column:set_meal_role_id" json:"setMealRoleId" description:"关联购买套餐"`
	Telephone           string    `gorm:"column:telephone" json:"telephone" description:"联系电话"`
	Address             string    `gorm:"column:address" json:"address" description:"地址"`
	Demand              string    `gorm:"column:demand" json:"demand" description:"需求"`
	CreatedTime         JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	CreatorId           int64     `gorm:"column:creator_id" json:"creatorId"`
	UpdatedTime         JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	UpdatorId           int64     `gorm:"column:updator_id" json:"updatorId"`
	InvitationTelephone string    `gorm:"column:invitation_telephone" json:"invitationTelephone" description:"邀请人联系电话"`
	Enabled             int64     `gorm:"column:enabled" json:"enabled"`
	Remarks             string    `gorm:"column:remarks" json:"remarks"`
}

func (orderDetail SsOrderDetail) TableName() string {
	return "ss_order_detail"
}

/*
 交易记录
 */
type SsTransactionRecord struct {
	ID          int64     `gorm:"column:id" json:"id"`
	UserId      int64     `gorm:"column:user_ud" json:"userId" description:"关联用户Id"`
	OrderId     int64     `gorm:"column:order_id" json:"orderId" description:"关联订单Id"`
	PaidAmount  float64   `gorm:"column:paid_amount" json:"paidAmount" description:"支付金额"`
	CreatorId   int64     `gorm:"column:creator_id" json:"creatorId"`
	CreatedTime JsonDateTime `gorm:"not null;default:NULL;column:created_time" json:"createdTime"`
	UpdatorId   int64     `gorm:"column:updator_id" json:"updatorId"`
	UpdatedTime JsonDateTime `gorm:"not null;default:NULL;column:updated_time" json:"updatedTime"`
	Remarks     string    `gorm:"remarks" json:"remarks"`
}

func (tr SsTransactionRecord) TableName() string {
	return "ss_transaction_record"
}
