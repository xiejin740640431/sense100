package response

import (
	"sense100/constant"
)

type SwaggerResponse struct {
	BaseResponse BaseResponse
}

type BaseResponseSwag struct {
	Code    int    `json:"code" example:"400" description:"状态码"`
	Message string `json:"message" example:"param bad" description:"状态消息"`
	Data    string `json:"data" example:"{}" description:"数据jsonObject"`
}

type BaseResponsePageSwag struct {
	Code      int    `json:"code" example:"400" description:"状态码"`
	Message   string `json:"message" example:"param bad" description:"状态消息"`
	PageIndex int    `json:"pageIndex" example:"0" description:"当前页索引"`
	PageSize  int    `json:"pageSize" example:"0" description:"页面数据个数"`
	Data      string `json:"data" example:"{}" description:"数据jsonObject"`
}

type BaseResponse struct {
	Code    constant.ResultStatus `json:"code" example:"400"`
	Message string                `json:"message" example:"param bad"`
	Data    interface{}           `json:"data" example:"{}"`
}

type BaseResponsePage struct {
	PageIndex int         `json:"pageIndex" example:"0"`
	PageSize  int         `json:"pageSize" example:"0"`
	Data      interface{} `json:"data" example:"{}"`
}

// 创建一个response
func RespWithCodeStatusData(resultStatus constant.ResultStatus, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    resultStatus,
		Message: resultStatus.GetMsg(),
		Data:    data,
	}
}

func RespWithStatus(resultStatus constant.ResultStatus) *BaseResponse {
	return &BaseResponse{
		Code:    resultStatus,
		Message: resultStatus.GetMsg(),
		Data:    nil,
	}
}

func RespWithCodeMsg(code constant.ResultStatus, msg string) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}

func PageRespWithCodeStatusData(resultStatus constant.ResultStatus, pageIndex, pageSize int, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    resultStatus,
		Message: resultStatus.GetMsg(),
		Data: BaseResponsePage{
			PageIndex: pageIndex,
			PageSize:  pageSize,
			Data:      data,
		},
	}
}
