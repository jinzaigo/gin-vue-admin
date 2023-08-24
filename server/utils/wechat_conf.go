package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 获取公众号配置
func GetOfficialAccountConf(appId string) *config.OfficialAccount {
	for _, account := range global.GVA_CONFIG.OfficialAccountList {
		if account.AppId == appId {
			return &account
		}
	}
	return nil
}
