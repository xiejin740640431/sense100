package request

import "sense100/util"

/*
 提交订单请求实体
 */
type SubmitOrderReq struct {
	UserId              int64  `json:"userId" description:"用户Id" required:"true"`
	SetMealRoleId       int64  `json:"setMealRoleId" description:"关联套餐规则Id" required:"true"`
	TransactionContent  int64  `json:"transactionContent" description:"交易内容" required:"true"`
	OrderSource         int64  `json:"orderSource" description:"订单来源(1官网、2微信小程序)" required:"true"`
	Telephone           string `json:"telephone" description:"联系电话" required:"true"`
	Address             string `json:"address" description:"联系地址" required:"true"`
	Demand              string `json:"demand" description:"需求"`
	InvitationTelephone string `json:"invitationTelephone" description:"邀请人联系电话"`
}

func (req SubmitOrderReq) CheckAvailable() bool {
	if req.SetMealRoleId == 0 || req.OrderSource == 0 || util.IsEmpty(req.Telephone) {
		return false
	}
	return true
}
