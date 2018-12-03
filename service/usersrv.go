package service

import (
	"sense100/model/response"
	"sense100/model/request"
	"sense100/constant"
	"net/http"
	"sense100/model"
	"errors"
	"sense100/util"
	"log"
)

func RegisterUser(req *request.RegisterUserReq) (int, *response.BaseResponse) {
	if !checkRegisterUserReq(req) {
		return http.StatusOK, response.RespWithStatus(constant.ErrParamDefect)
	}
	//验证用户是否存在
	value := getAvailableValue(req)
	userAccount := GetUserAccount(value)
	if userAccount != nil && userAccount.ID != 0 {
		return http.StatusOK, response.RespWithStatus(constant.ErrUserExist)
	} else {
		var err error
		tx := db.Begin()
		userInfo := model.SsUserInfo{Mobile: req.Mobile, NickName: req.NickName, HeadUrl: req.HeadUrl, CreatedTime: model.DateTimeNow(), Enabled: constant.Enable}
		err = tx.Create(&userInfo).Error
		if err != nil {
			tx.Rollback()
			return http.StatusOK, response.RespWithStatus(constant.ErrSystemError)
		} else {
			userAccount := model.SsUserAccount{UserId: userInfo.ID, WxId: req.WxId, Password: req.Password, IpAddress: req.IpAddress, Enabled: constant.Enable}
			err = tx.Create(&userAccount).Error
			if err != nil {
				tx.Rollback()
				return http.StatusOK, response.RespWithStatus(constant.ErrSystemError)
			} else {
				tx.Commit()
				return http.StatusOK, response.RespWithStatus(constant.Success)
			}
		}
	}
}

/*
 根据参数查询用户信息
 */
func FindUser(req *request.LoginReq) *response.BaseResponse {
	if req.LoginType != constant.LoginTypeWx && req.LoginType != constant.LoginTypeMobile {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	if req.LoginType == constant.LoginTypeMobile && (util.IsEmpty(req.UserName) || util.IsEmpty(req.Password)) {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	if req.LoginType == constant.LoginTypeWx && util.IsEmpty(req.WxId) {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	var userAccount model.SsUserAccount
	newDB := db.Table("ss_user_account").Select("ss_user_account.*").
		Joins("left join ss_user_info on ss_user_account.user_id = ss_user_info.id")
	if !util.IsEmpty(req.UserName) {
		log.Println("---->", req.UserName)
		newDB = newDB.Where("ss_user_info.mobile = ?", req.UserName).Or("ss_user_account.user_name = ?", req.UserName)
	}
	if !util.IsEmpty(req.WxId) {
		log.Println("---->", req.WxId)
		newDB = newDB.Where("ss_user_account.wx_id = ?", req.WxId)
	}
	newDB = newDB.Preload("UserInfo")
	newDB.Find(&userAccount)
	if userAccount.ID == 0 || userAccount.UserInfo == nil || userAccount.UserInfo.ID == 0 {
		return response.RespWithStatus(constant.ErrUserNotExist)
	} else if req.LoginType == constant.LoginTypeMobile && userAccount.Password != req.Password {
		return response.RespWithStatus(constant.ErrPasswordError)
	}
	return response.RespWithCodeStatusData(constant.Success, &userAccount)
}

/*
 根据id检测用户是否存在
 */
func CheckUserExist(id int64) bool {
	var userAccount model.SsUserAccount
	db.Find(&userAccount, id)
	if userAccount.ID != 0 {
		return true
	}
	return false
}

/*
 创建用户账号信息
 */
func GetUserAccount(value string) *model.SsUserAccount {
	var user model.SsUserAccount
	db.Table("ss_user_account").Select("ss_user_account.*").
		Joins("left join ss_user_info on ss_user_account.user_id = ss_user_info.id").
		Where("ss_user_account.user_name = ?", value).
		Or("ss_user_account.wx_id = ?", value).
		Or("ss_user_info.mobile = ?", value).
		Find(&user)
	return &user
}

/*
	判断req是否可用
 */
func checkRegisterUserReq(req *request.RegisterUserReq) bool {
	if req == nil {
		return false
	}
	if req.Type == constant.RegisterTypeWx && req.WxId != "" {
		return true
	} else if req.Type == constant.RegisterTypeMobile && req.Mobile != "" && req.Password != "" {
		return true
	}
	return false
}

/*
 获取账号中可用的值，如微信id、手机号
 */
func getAvailableValue(req *request.RegisterUserReq) string {
	if req == nil {
		return ""
	}
	if req.Type == constant.RegisterTypeWx && req.WxId != "" {
		return req.WxId
	} else if req.Type == constant.RegisterTypeMobile && req.Mobile != "" && req.Password != "" {
		return req.Mobile
	}
	return ""
}

/*
 创建或更新浏览记录信息
 */
func CreateOrUpdateBrowseRecord(req *model.AddBrowseRecordReq) *response.BaseResponse {
	if req == nil || req.WxId == "" || req.ContentId == 0 || req.Type == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	//创建或更新游客信息，并返回游客信息
	tourist, err := CreatedOrUpdateTourist(&req.SsTourist)
	if err != nil {
		return response.RespWithStatus(constant.ErrSystemError)
	}
	//根据条件 查询记录是否已存在
	browseRecord := model.SsSpBrowseRecord{
		TouristId: tourist.ID,
		ContentId: req.ContentId,
		Type:      req.Type,
	}
	db.Find(&browseRecord)

	//执行创建或更新操作
	tx := db.Begin()
	if db.NewRecord(browseRecord) {
		browseRecord.Enabled = constant.Enable
		browseRecord.CreatedTime = model.DateTimeNow()
		browseRecord.BrowseNum = 1
		err = tx.Create(&browseRecord).Error
	} else {
		browseRecord.UpdatedTime = model.DateTimeNow()
		browseRecord.BrowseNum++
		err = tx.Model(&browseRecord).Updates(&browseRecord).Error
	}
	//检测是否有错误
	if err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return response.RespWithStatus(constant.Success)
	}
}

/**
 创建或更新游客信息
 */
func CreatedOrUpdateTourist(newTourist *model.SsTourist) (*model.SsTourist, error) {
	if newTourist == nil || newTourist.WxId == "" {
		return nil, errors.New("param defect")
	}
	var tourist model.SsTourist
	var err error
	tx := db.Begin()

	err = db.Where("wx_id = ?", newTourist.WxId).Find(&tourist).Error

	if err != nil || db.NewRecord(&tourist) {
		err = tx.Create(&newTourist).Error
	} else {
		err = tx.Model(&tourist).Updates(&tourist).Error
	}

	if err != nil {
		tx.Rollback()
		return nil, err
	} else {
		tx.Commit()
		return &tourist, nil
	}
}

/*
 创建点赞记录
 */
func CreatePraise(praise *model.SsSpPraise) *response.BaseResponse {
	if praise == nil || praise.UserId == 0 || praise.PraiseContentId == 0 || praise.Type == 0 {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	db.Where(map[string]interface{}{"user_id": praise.UserId, "praise_content_id": praise.PraiseContentId, "type": praise.Type}).Find(&praise)

	var err error
	tx := db.Begin()
	if db.NewRecord(praise) {
		praise.CreatedTime = model.DateTimeNow()
		praise.Enabled = constant.Enable
		err = tx.Create(&praise).Error
	} else {
		if praise.Enabled == constant.Enable {
			praise.Enabled = constant.Disable
		} else {
			praise.Enabled = constant.Enable
		}
		err = tx.Table("ss_sp_praise").Where("id = ?", praise.ID).Updates(map[string]interface{}{"enabled": praise.Enabled}).Error
	}

	if err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return response.RespWithStatus(constant.Success)
	}
}
