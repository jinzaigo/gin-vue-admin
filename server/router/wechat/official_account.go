package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OfficialAccountRouter struct{}

func (oa *OfficialAccountRouter) InitOfficialAccountRouter(Router *gin.RouterGroup) {
	r := Router.Group("officialaccount").Use(middleware.OfficialAccountAuth())
	api := v1.ApiGroupApp.WechatApiGroup.OfficialAccountApi
	{
		r.GET("index", api.Index)
		r.POST("index", api.IndexPost)
	}
}
