package service

import (
	"sense100/model/response"
	"sense100/model"
	"net/http"
	"sense100/constant"
	"log"
	"github.com/jinzhu/gorm"
)

func FindRegionListOnParentId(parentId int) (int, *response.BaseResponse) {
	var regions []model.SsRegion
	err := db.Where("parent_id = ?", parentId).Find(&regions).Error
	if err != nil {
		log.Println(err)
		return http.StatusOK, response.RespWithStatus(constant.ErrSystemError)
	}
	return http.StatusOK, response.RespWithCodeStatusData(constant.Success, &regions)
}

/*
 创建或更新行业类型
 */
func CreateOrUpdateTrade(trade *model.SsTrade) *response.BaseResponse {
	if trade.Name == "" || trade.CreatorId == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	var err error
	var admin model.SsAdmin
	err = db.Find(&admin, trade.CreatorId).Error
	if err != nil || admin.ID == 0 {
		return response.RespWithStatus(constant.ErrAdminNotExist)
	}
	tx := db.Begin()
	if db.NewRecord(trade) {
		trade.Enabled = constant.Enable
		trade.CreatedTime = model.DateTimeNow()
		err = tx.Create(&trade).Error
	} else {
		trade.UpdatedTime = model.DateTimeNow()
		err = tx.Model(&trade).Updates(&trade).Error
	}

	if err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return response.RespWithStatus(constant.Success)
	}
}

/*
 根据父级Id查询行业类型列表
 */
func FindTradeOnParentId(parentId int64) *response.BaseResponse {
	var trades []model.SsTrade

	err := db.Where("parent_id = ?", parentId).Find(&trades).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.RespWithCodeStatusData(constant.Success, trades)
		} else {
			return response.RespWithStatus(constant.ErrSystemError)
		}
	} else {
		return response.RespWithCodeStatusData(constant.Success, trades)
	}
}

/*
 根据id查询行业类型
 */
func FindTradeOnId(id int64) (*model.SsTrade, error) {
	var trade model.SsTrade
	err := db.Find(&trade, id).Error
	return &trade, err
}
