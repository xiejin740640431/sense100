package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/elazarl/go-bindata-assetfs"
	"sense100/asset"
)

func MapRoutes() *gin.Engine {
	router := gin.New()
	//要在添加路由之前配置跨域
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//swagger 文件系统
	router.ServeFiles("/swagger/*filepath", &assetfs.AssetFS{
		Asset:     asset.Asset,
		AssetDir:  asset.AssetDir,
		AssetInfo: asset.AssetInfo,
	})
	router.Static("staticServer", "./staticServer")
	commonRouter := router.Group("/common")
	{
		commonRouter.GET("/getSubsetRegion/:parentId", getSubsetRegion)
	}
	userRouter := router.Group("/user")
	{
		userRouter.POST("/registerUser", registerUser)
		userRouter.POST("/login", login)
		userRouter.POST("/addBrowseRecord", addBrowseRecord)
		userRouter.POST("/doPraise", doPraise)
	}
	companyGroup := router.Group("/company")
	{
		companyGroup.POST("/createOrUpdateFacade", createOrUpdateFacade)
		companyGroup.GET("/getFacadeOnUserId/:userId", getFacadeOnUserId)
		companyGroup.POST("/publishDynamic", publishDynamic)
		companyGroup.POST("/getDynamicListPage", getDynamicListPage)
		companyGroup.POST("/createOrUpdateTrade", createOrUpdateTrade)
		companyGroup.GET("/getTradesOnParentId/:parentId", getTradesOnParentId)
	}
	carrierGroup := router.Group("/carrier")
	{
		carrierGroup.POST("/createOrUpdateCarrier", createOrUpdateCarrier)
		carrierGroup.POST("/createOrUpdateSetMeal", createOrUpdateSetMeal)
		carrierGroup.POST("/getSetMealInfo/:carrierId/:mouldId", getSetMealInfo)
	}
	programGroup := router.Group("/program")
	{
		programGroup.POST("/createOrUpdateMouldCategory", createOrUpdateMouldCategory)
		programGroup.GET("/getMouldCategoryOnParentId/:parentId", getMouldCategoriesOnParentId)
		programGroup.DELETE("/delMouldCategory/:id", delMouldCategory)
		programGroup.POST("/createOrUpdateMould", createOrUpdateMould)
		programGroup.GET("/getMouldListOnCategoryId/:categoryId", getMouldListOnCategoryId)
		programGroup.GET("/getMouldInfo/:id", getMouldInfo)
		programGroup.DELETE("/delMould/:id", delMould)
	}
	orderGroup := router.Group("/order")
	{
		orderGroup.POST("/submitOrder", submitOrder)
	}
	thirdGroup := router.Group("/third")
	{
		thirdGroup.GET("/getWxAccessToken/:code", getWxAccessToken)
	}
	return router
}
