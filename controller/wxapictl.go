package controller

import (
	"github.com/gin-gonic/gin"
	"sense100/constant"
	"sense100/util"
	"net/http"
	"sense100/model/response"
	"fmt"
	"sense100/config"
	"encoding/json"
	"log"
)

//@Summary 根据code获取微信accessToken
//@Description 根据code获取微信accessToken
//@Accept json
//@Produce json
//@Param code path string true "code"
//@Success 200 {object} response.WxAccessTokenResp
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Third
//@Router /third/getWxAccessToken/{code} [get]
func getWxAccessToken(c *gin.Context) {
	code := c.Params.ByName(constant.Code)
	if !util.IsEmpty(code) {
		url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", config.WxAppId, config.WxAppSecret, code)
		log.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, response.RespWithCodeMsg(constant.ErrSystemError, err.Error()))
			return
		}
		//关闭body
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			//创建json解析实例
			enc := json.NewDecoder(resp.Body)
			var wxAccessTokenResp response.WxAccessTokenResp
			//解析json
			err = enc.Decode(&wxAccessTokenResp)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusOK, response.RespWithCodeMsg(constant.ErrSystemError, err.Error()))
			} else {
				c.JSON(http.StatusOK, response.RespWithCodeStatusData(constant.Success, &wxAccessTokenResp))
			}
		} else {
			c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrSystemError))
		}
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}
