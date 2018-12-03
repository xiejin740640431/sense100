package service

import (
	"sense100/model/request"
	"sense100/model/response"
	"sense100/constant"
	"sense100/model"
	"sense100/util"
	"time"
)

/*
	创建或更新承运商
 */
func CreateOrUpdateCarrier(req *request.CreateOrUpdateCarrierReq) *response.BaseResponse {
	//请求参数可用性检测
	if req != nil && req.Check() {
		if !CheckAdminExist(req.AdminId) {
			return response.RespWithStatus(constant.ErrUserNotExist)
		}
		//创建承运商结构实体
		carrierInfo := model.SsCarrierInfo{
			ID:                       req.ID,
			Name:                     req.Name,
			CompanyName:              req.CompanyName,
			ContactPerson:            req.ContactPerson,
			ContactPhone:             req.ContactPhone,
			ContactEmail:             req.ContactEmail,
			LegalPersonEmail:         req.LegalPersonEmail,
			LegalPersonIdCard:        req.LegalPersonIdCard,
			LegalPersonName:          req.LegalPersonName,
			LegalPersonPhone:         req.LegalPersonPhone,
			IdCardHeadPositivePicUrl: req.IdCardHeadPositivePicUrl,
			IdCardInsteadPicUrl:      req.IdCardInsteadPicUrl,
			IdCardPositivePicUrl:     req.IdCardHeadPositivePicUrl,
			BusinessLicensePicUrl:    req.BusinessLicensePicUrl,
			OrganizationPicUrl:       req.OrganizationPicUrl,
			CreatorId:                req.AdminId,
			Enabled:                  constant.Enable,
		}
		var err error
		tx := db.Begin()
		if db.NewRecord(&carrierInfo) {
			//执行新增操作
			carrierInfo.CreatedTime = time.Now()
			err = tx.Create(&carrierInfo).Error
			if err == nil {
				carrierAccount := model.SsCarrierAccount{
					CarrierInfoId: carrierInfo.ID,
					CreatedTime:   model.DateTimeNow(),
					Enabled:       constant.Enable,
				}
				err = tx.Create(&carrierAccount).Error
			}
		} else {
			//执行更新操作
			carrierInfo.UpdatedTime = time.Now()
			err = tx.Model(&carrierInfo).Updates(&carrierInfo).Error
		}
		//如果有错误，则回滚事务，并返回错误，如果没有错误，则提交事务，返回操作成功
		if err != nil {
			tx.Rollback()
			return response.RespWithStatus(constant.ErrSystemError)
		} else {
			tx.Commit()
			return response.RespWithStatus(constant.Success)
		}
	} else {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
}

/*
	创建或更新套餐
 */
func CreateOrUpdateSetMeal(setMeal *model.SsSetMeal) *response.BaseResponse {
	if setMeal.CarrierId != 0 && setMeal.MouldId != 0 && !util.IsEmpty(setMeal.Name) {
		if !CheckUserExist(setMeal.CreatorId) {
			return response.RespWithStatus(constant.ErrUserNotExist)
		}
		isNewRecord := db.NewRecord(&setMeal)
		//如果有role，则检测role的可用性
		if setMeal.MealRoles != nil && len(setMeal.MealRoles) != 0 {
			for _, role := range setMeal.MealRoles {
				//检测role可用性
				if checkMealRole(role) {
					//设置创建或更新时间
					if isNewRecord {
						role.CreatedTime = model.DateTimeNow()
					} else {
						role.UpdatedTime = model.DateTimeNow()
					}
				} else {
					return response.RespWithStatus(constant.ErrParamDefect)
				}
			}
		}
		var err error
		tx := db.Begin()
		//判断是否是新数据，然后执行相应的创建和更新操作
		if isNewRecord {
			setMeal.CreatedTime = model.DateTimeNow()
			err = tx.Create(&setMeal).Error
		} else {
			setMeal.UpdatedTime = model.DateTimeNow()
			err = tx.Save(&setMeal).Error
		}
		if err != nil {
			tx.Rollback()
			return response.RespWithStatus(constant.ErrSystemError)
		} else {
			tx.Commit()
			return response.RespWithStatus(constant.Success)
		}
	} else {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
}

/*
	查找套餐信息
 */
func FindSetMealInfo(carrierId, mouldId int64) *response.BaseResponse {
	if carrierId == 0 || mouldId == 0 {
		return response.RespWithStatus(constant.ErrSystemError)
	}
	var setMeal model.SsSetMeal
	db.Where(model.SsSetMeal{CarrierId: carrierId, MouldId: mouldId}).
		Preload("MealRoles", "enabled = ?", constant.Enable).
		Find(&setMeal)
	return response.RespWithCodeStatusData(constant.Success, &setMeal)
}

/*
	检测role可用性
 */
func checkMealRole(role model.SsSetMealRole) bool {
	if util.IsEmpty(role.Name) || role.SalePrice == 0 || role.SetMealId == 0 {
		return false
	}
	return true
}
