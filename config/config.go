package config

var ConfigJson = `
{
  "Server": "http://47.106.36.243:9494",
  "ImageServer": "http://47.106.36.243:9494/static/images",
  "StaticServer": "http://47.106.36.243:9494/static",
  "RuntimeModel": "dev",
  "LogLevel": "debug",
  "mysql": "root:xhshop123!@#@(47.106.36.243:3306)/small_program_test?charset=utf8mb4&parseTime=true&loc=Local",
  "Port": ":9494",
  "SwaggerJson": "swagger.json"
}
`
var (
	WxAppId     = "wxb951051ae74e9f0e"
	WxAppSecret = "d271c13642bb3f0c41ca568e10551e80"
)
