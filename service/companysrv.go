package service

import (
	"sense100/model"
	"sense100/model/response"
	"net/http"
	"sense100/constant"
	"time"
	"sense100/util"
	"github.com/jinzhu/gorm"
)

func CreateOrUpdateFacade(facade *model.SsCompanyFacadeInfo) (int, *response.BaseResponse) {
	if facade == nil || facade.UserId == 0 || util.IsEmpty(facade.Name) || facade.TradeId == 0 {
		return http.StatusOK, response.RespWithStatus(constant.ErrParamDefect)
	}
	if !CheckUserExist(facade.UserId) {
		return http.StatusOK, response.RespWithStatus(constant.ErrUserNotExist)
	}
	trade, err := FindTradeOnId(facade.TradeId)
	if err != nil || trade == nil || trade.ID == 0 {
		return http.StatusOK, response.RespWithStatus(constant.ErrTradeNoExist)
	}
	tx := db.Begin()
	if db.NewRecord(facade) {
		facade.CreatedTime = model.JsonDateTime{T: time.Now()}
		facade.CreatorId = facade.UserId
		//默认欠费状态，购买套餐后，才更改为正常状态
		facade.State = constant.FacadeStateArrears
		err = tx.Create(&facade).Error
	} else {
		facade.UpdatedTime = model.JsonDateTime{T: time.Now()}
		facade.UpdatorId = facade.UserId
		err = tx.Model(&facade).Updates(facade).Error
	}
	if err != nil {
		tx.Rollback()
		return http.StatusOK, response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return http.StatusOK, response.RespWithStatus(constant.Success)
	}
}

/*
 根据userId查询企业门面信息
 */
func FindFacadeOnUserId(userId int) (int, *response.BaseResponse) {
	var facade model.CompanyFacadeInfoResp
	err := db.Table("ss_company_facade_info cfi").
		Select("cfi.*,sbr.browse_num").
		Joins("LEFT JOIN ss_sp_browse_record sbr ON sbr.content_id = cfi.id").
		Where("cfi.user_id = ?", userId).
		Find(&facade).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusOK, response.RespWithStatus(constant.ErrRecordNotFound)
		} else {
			return http.StatusOK, response.RespWithStatus(constant.ErrSystemError)
		}
	}
	return http.StatusOK, response.RespWithCodeStatusData(constant.Success, &facade)
}
