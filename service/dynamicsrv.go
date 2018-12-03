package service

import (
	"sense100/model"
	"sense100/model/response"
	"sense100/constant"
	"sense100/model/request"
	"github.com/jinzhu/gorm"
)

/*
	创建动态
 */
func CreateDynamic(dynamic *model.SsSpDynamic) *response.BaseResponse {
	if dynamic == nil || !dynamic.CreateCheck() {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	if !CheckUserExist(dynamic.UserId) {
		return response.RespWithStatus(constant.ErrUserNotExist)
	}
	tx := db.Begin()
	dynamic.CreatedTime = model.DateTimeNow()
	dynamic.CreatorId = dynamic.UserId
	if err := tx.Create(dynamic).Error; err != nil {
		tx.Rollback()
		return response.RespWithStatus(constant.ErrSystemError)
	} else {
		tx.Commit()
		return response.RespWithStatus(constant.Success)
	}
}

/*
	分页查询动态集合
 */
func FindDynamicsOnPage(req *request.GetDynamicListPageReq) *response.BaseResponse {
	if req == nil || !req.Check() {
		return response.RespWithStatus(constant.ErrParamDefect)
	}
	var dynamics []model.DynamicResp
	var gormDb *gorm.DB
	if req.UserId != 0 {
		gormDb = db.Table("ss_sp_dynamic sd").
			Select(`sd.*,(SELECT count(*) FROM ss_sp_praise sp WHERE sp.praise_content_id = sd.id AND sp.enabled = ?) as praise_num,
			(SELECT EXISTS (SELECT * FROM ss_sp_praise sp2 WHERE sp2.praise_content_id = sd.id AND sp2.enabled = ? AND sp2.user_id = ?)) AS is_praise`,
			constant.Enable, constant.Enable, req.UserId)
	} else {
		gormDb = db.Table("ss_sp_dynamic sd").
			Select("sd.*,(SELECT count(*) FROM ss_sp_praise sp WHERE sp.praise_content_id = sd.id AND sp.enabled = ?) as praise_num", constant.Enable)
	}
	if req.FacadeId != 0 {
		gormDb = gormDb.Where("facade_id = ?", req.FacadeId)
	}
	if req.ProgramId != 0 {
		gormDb = gormDb.Where("program_id = ?", req.ProgramId)
	}
	if req.CompanyId != 0 {
		gormDb = gormDb.Where("company_id = ?", req.CompanyId)
	}
	if req.IsPage {
		gormDb = gormDb.Offset(req.PageIndex * req.PageSize).Limit(req.PageSize)
	}
	gormDb.Scan(&dynamics)
	return response.PageRespWithCodeStatusData(constant.Success, req.PageIndex, req.PageSize, &dynamics)
}
