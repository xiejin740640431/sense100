package controller

import (
	"github.com/gin-gonic/gin"
	"sense100/model"
	"net/http"
	"sense100/model/response"
	"sense100/constant"
	"strconv"
	"sense100/service"
)

//@Summary 创建或更新模板类型
//@Description 创建或更新模板类型
//@Accept json
//@Produce json
//@Param mouldCategory body model.SsMouldCategory true "mouldCategory"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Program
//@Router /program/createOrUpdateMouldCategory [post]
func createOrUpdateMouldCategory(c *gin.Context) {
	var mouldCategory model.SsMouldCategory
	if c.Bind(&mouldCategory) {
		c.JSON(http.StatusOK, service.CreateOrUpdateMouldCategory(&mouldCategory))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 根据parentId获取类型子集
//@Description 根据parentId获取类型子集
//@Accept json
//@Produce json
//@Param parentId path int true "parentId"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Program
//@Router /program/getMouldCategoryOnParentId/{parentId} [get]
func getMouldCategoriesOnParentId(c *gin.Context) {
	parentId, err := strconv.ParseInt(c.Params.ByName(constant.ParentId), 10, 64)
	if err == nil {
		c.JSON(http.StatusOK, service.FindMouldCategoriesOnParentId(parentId))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 根据Id删除类型
//@Description 根据Id删除类型
//@Accept json
//@Produce json
//@Param id path int true "id"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Program
//@Router /program/delMouldCategory/{id} [delete]
func delMouldCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName(constant.Id), 10, 64)
	if err == nil {
		c.JSON(http.StatusOK, service.DelMouldCategory(id))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 创建模板
//@Description 创建模板
//@Accept json
//@Produce json
//@Param programMould body model.SsProgramMould true "programMould"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Program
//@Router /program/createOrUpdateMould [post]
func createOrUpdateMould(c *gin.Context) {
	var mould model.SsProgramMould
	if c.Bind(&mould) {
		c.JSON(http.StatusOK, service.CreateOrUpdateMould(&mould))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 根据类型获取模板集合
//@Description 根据类型获取模板集合
//@Accept json
//@Produce json
//@Param categoryId path int true "categoryId"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Program
//@Router /program/getMouldListOnCategoryId/{categoryId} [get]
func getMouldListOnCategoryId(c *gin.Context) {
	categoryId, err := strconv.ParseInt(c.Params.ByName(constant.CategoryId), 10, 64)
	if err == nil {
		c.JSON(http.StatusOK, service.FindMouldListOnCategoryId(categoryId))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 获取模板信息
//@Description 获取模板信息
//@Accept json
//@Produce json
//@Param id path int true "id"
//@Success 200 {object} model.SsProgramMould
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Program
//@Router /program/getMouldInfo/{id} [get]
func getMouldInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName(constant.Id), 10, 64)
	if err == nil {
		c.JSON(http.StatusOK, service.FindMouldInfo(id))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 删除模板信息
//@Description 删除模板信息
//@Accept json
//@Produce json
//@Param id path int true "id"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Program
//@Router /program/delMould/{id} [delete]
func delMould(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName(constant.Id), 10, 64)
	if err == nil {
		c.JSON(http.StatusOK, service.DelMould(id))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

