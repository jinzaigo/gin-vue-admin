package wechat

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	OfficialAccountApi
}

var (
	officialAccountService = &service.ServiceGroupApp.WechatServiceGroup.OfficialAccountService
)
