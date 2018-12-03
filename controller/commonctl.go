package controller

import (
	"github.com/gin-gonic/gin"
	"sense100/constant"
	"strconv"
	"net/http"
	"sense100/model/response"
	"sense100/service"
)

//@Summary 获取子集城市集合
//@Description
//@Accept json
//@Produce json
//@Param parentId path int	 true "parentId"
//@Success 200 {array} model.SsRegion
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Common
//@Router	/common/getSubsetRegion/{parentId} [get]
func getSubsetRegion(c *gin.Context) {
	parentId, err := strconv.Atoi(c.Params.ByName(constant.ParentId))
	if err != nil {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrSystemError))
	} else {
		c.JSON(service.FindRegionListOnParentId(parentId))
	}
}
