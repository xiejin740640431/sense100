package service

import (
	"sense100/model/response"
	"sense100/model/request"
	"sense100/constant"
	"sense100/model"
)

/*
 创建订单信息
 */
func CreateOrderInfo(req *request.SubmitOrderReq) *response.BaseResponse {
	//请求参数可用性判断
	if req == nil || !req.CheckAvailable() {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	//根据setMealRoleId查看role是否存在
	var setMealRole model.SsSetMealRole
	db.Find(&setMealRole, req.SetMealRoleId)
	if setMealRole.ID == 0 {
		return response.RespWithStatus(constant.ErrSetMealRoleIdNotExist)
	}
	var err error
	tx := db.Begin()
	//创建订单明细实体
	orderDetail := model.SsOrderDetail{
		SetMealRoleId:       setMealRole.ID,
		Telephone:           req.Telephone,
		Address:             req.Address,
		Demand:              req.Demand,
		InvitationTelephone: req.InvitationTelephone,
		CreatorId:           req.UserId,
		CreatedTime:         model.DateTimeNow(),
		Enabled:             constant.Enable,
	}
	//创建订单明细
	err = tx.Create(&orderDetail).Error
	//错误处理，回滚事务
	if err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	}
	//创建订单信息实体
	orderInfo := model.SsOrderInfo{
		UserId:        req.UserId,
		OrderDetailId: orderDetail.ID,
		TotalAmount:   setMealRole.SalePrice,
		CreatorId:     req.UserId,
		CreatedTime:   model.DateTimeNow(),
	}
	//创建订单信息
	err = tx.Create(&orderInfo).Error
	//错误处理，回滚事务或提交事务
	if err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return response.RespWithStatus(constant.Success)
	}

}

func CreateTransactionRecord(){

}