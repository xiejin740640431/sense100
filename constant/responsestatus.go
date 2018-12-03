package constant

type ResultStatus int

const (
	Success                    ResultStatus = 200
	ErrSystemError                          = 500
	Error                                   = -1000
	ErrParamDefect                          = -1001
	ErrUserExist                            = -1002
	ErrParamError                           = -1003
	ErrNamePwdError                         = -1004
	ErrUserNotExist                         = -1005
	ErrTechnologyNotExist                   = -1006
	ErrUploadFileFailed                     = -1007
	ErrMouldCategoryIdNotExist              = -1008
	ErrSetMealIdNotExist                    = -1009
	ErrSetMealRoleIdNotExist                = -1010
	ErrPasswordError                        = -1011
	ErrRecordNotFound                       = -1012
	ErrAdminNotExist                        = -1013
	ErrTradeNoExist                         = -1014
)

var StatusMsg = map[ResultStatus]string{
	Success:                    "成功",
	ErrSystemError:             "服务器错误",
	Error:                      "失败",
	ErrParamDefect:             "参数缺失",
	ErrUserExist:               "用户已存在",
	ErrParamError:              "请求参数错误",
	ErrNamePwdError:            "用户名或密码错误",
	ErrUserNotExist:            "用户不存在",
	ErrTechnologyNotExist:      "不存在",
	ErrUploadFileFailed:        "上传文件失败",
	ErrMouldCategoryIdNotExist: "模板类型Id不存在",
	ErrSetMealIdNotExist:       "套餐Id不存在",
	ErrSetMealRoleIdNotExist:   "套餐不存在",
	ErrPasswordError:           "密码错误",
	ErrRecordNotFound:          "记录不存在",
	ErrAdminNotExist:           "管理员不存在",
	ErrTradeNoExist:            "行业类型不存在",
}

func (c ResultStatus) GetMsg() string {
	return StatusMsg[c]
}
