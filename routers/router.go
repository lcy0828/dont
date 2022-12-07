package routers

import (
	"dont/middleware/jwt"
	"dont/pkg/setting"
	"dont/routers/api"
	"dont/routers/api/v1"
	"dont/routers/mod"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	apimod := r.Group("/mod")
	//apimod.Use(jwt.JWT())
	{
		//获取标签列表
		apimod.GET("/tags", mod.SearchMod)
		apimod.GET("/add", mod.AddMod)
		apimod.POST("/down", mod.DownloadMod)
		apimod.GET("/down", mod.DownloadMod)
		//新建标签
		//apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		//apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		//apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
