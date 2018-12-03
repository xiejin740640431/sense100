package request

type RegisterUserReq struct {
	Mobile    string `json:"mobile" description:"手机号"`
	WxId      string `json:"wxId" description:"微信id"`
	Password  string `json:"password" description:"密码"`
	IpAddress string `json:"ipAddress"  description:"ip地址"`
	NickName  string `json:"nickName" description:"昵称"`
	HeadUrl   string `json:"headUrl" description:"头像url"`
	Type      int    `json:"type" description:"1:手机号，2:微信"`
}

type LoginReq struct {
	UserName  string `json:"userName" example:"18664394383" description:"用户名"`
	Password  string `json:"password" example:"123456" description:"密码(md5加密)"`
	WxId      string `json:"wxId" description:"微信id"`
	LoginType int    `json:"loginType" description:"登录类型(1:手机号,2:微信)" required:"true"`
}
