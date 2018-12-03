package service

import (
	"sense100/model/response"
	"sense100/model"
	"sense100/util"
	"sense100/constant"
	"time"
)

/*
 创建或更新模版类型信息
 */
func CreateOrUpdateMouldCategory(mouldCategory *model.SsMouldCategory) *response.BaseResponse {
	//获取是否是新记录
	isNewRecord := db.NewRecord(&mouldCategory)
	//获取adminId
	var adminId = util.If(isNewRecord, mouldCategory.CreatorId, mouldCategory.UpdatorId).(int64)
	if util.IsEmpty(mouldCategory.CategoryName) || adminId == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	//判断管理员是否存在
	if !CheckAdminExist(adminId) {
		return response.RespWithStatus(constant.ErrUserNotExist)
	}
	//开启事务
	var err error
	tx := db.Begin()
	if isNewRecord {
		mouldCategory.CreatedTime = model.DateTimeNow()
		mouldCategory.Enabled = constant.Enable
		err = tx.Create(&mouldCategory).Error
	} else {
		mouldCategory.UpdatedTime = model.DateTimeNow()
		err = tx.Model(&mouldCategory).Updates(&mouldCategory).Error
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
 根据父级id查询模版类型结婚
 */
func FindMouldCategoriesOnParentId(parentId int64) *response.BaseResponse {
	if parentId == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	var mouldCategories []model.SsMouldCategory
	db.Where(map[string]interface{}{"parent_id": parentId, "enabled": constant.Enable}).Find(&mouldCategories)
	return response.RespWithCodeStatusData(constant.Success, &mouldCategories)
}

/*
 删除模版类型
 */
func DelMouldCategory(id int64) *response.BaseResponse {
	if id == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	tx := db.Begin()
	if err := tx.Table("ss_mould_category").Where("id = ?", id).Updates(map[string]interface{}{"enabled": constant.Disable, "updated_time": time.Now()}).Error; err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return response.RespWithStatus(constant.Success)
	}
}

/*
 创建或更新模版
 */
func CreateOrUpdateMould(mould *model.SsProgramMould) *response.BaseResponse {
	//获取是否是新记录
	isNewRecord := db.NewRecord(&mould)
	//获取adminId
	var adminId = util.If(isNewRecord, mould.CreatorId, mould.UpdatorId).(int64)
	if adminId == 0 || util.IsEmpty(mould.Name) || util.IsEmpty(mould.Description) ||
		mould.SalePrice == 0 || util.IsEmpty(mould.CoverUrl) || util.IsEmpty(mould.DetailPictureUrls) || util.IsEmpty(mould.SourceUrl) {
		return response.RespWithStatus(constant.ErrParamDefect)
	}

	//判断管理员是否存在
	if !CheckAdminExist(adminId) {
		return response.RespWithStatus(constant.ErrUserNotExist)
	}

	//判断模版类型是否存在
	var category model.SsMouldCategory
	db.Where("id = ?", mould.CategoryId).Find(&category)
	if category.ID == 0 {
		return response.RespWithStatus(constant.ErrMouldCategoryIdNotExist)
	}

	//开启事务
	var err error
	tx := db.Begin()
	if isNewRecord {
		mould.CreatedTime = model.DateTimeNow()
		mould.Enabled = constant.Enable
		err = tx.Create(&mould).Error
	} else {
		mould.UpdatedTime = model.DateTimeNow()
		err = tx.Model(&mould).Updates(&mould).Error
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
 根据类型查找模版集合
 */
func FindMouldListOnCategoryId(categoryId int64) *response.BaseResponse {
	if categoryId == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	var moulds []model.SsProgramMould
	db.Where(map[string]interface{}{"category_id": categoryId, "enabled": constant.Enable}).Find(&moulds)
	return response.RespWithCodeStatusData(constant.Success, &moulds)
}

/*
 根据id获取模版信息
 */
func FindMouldInfo(id int64) *response.BaseResponse {
	if id == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	var mouldInfo model.SsProgramMould
	db.Where("id = ?", id).Find(&mouldInfo)
	return response.RespWithCodeStatusData(constant.Success, &mouldInfo)
}

/*
 根据id删除模版
 */
func DelMould(id int64) *response.BaseResponse {
	if id == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	tx := db.Begin()
	if err := tx.Table("ss_program_mould").Where("id = ?", id).Updates(map[string]interface{}{"enabled": constant.Disable, "updated_time": time.Now()}).Error; err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return response.RespWithStatus(constant.Success)
	}
}
