package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"flag"
	"github.com/swaggo/swag/gen"
	"sense100/util"
	"sense100/service"
	"sense100/controller"
	"log"
	"sense100/config"
)

// @title sense100
// @version 1.5
// @description sense100 接口文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 740640431@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 47.106.36.243:9494
// @BasePath /
func main() {
	isParseSwagger := flag.Bool("isParseSwagger", false, "is parse swagger")
	flag.Parse()
	if *isParseSwagger {
		gen.New().BuildSwaggerJson("./", "./main.go", "staticServer/swagger", config.Host)
	} else {
		//加载配置
		util.LoadConf()
		////连接数据库服务
		service.ConnectDB()
		//获取设置路由
		engine := controller.MapRoutes()
		err := engine.Run(config.Port)
		if err != nil {
			log.Panic(err)
		}
	}
}
