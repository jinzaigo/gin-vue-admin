package wechat

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"os"
	"sync"
)

type OfficialAccountService struct {
	lock sync.Mutex
	apps map[string]*officialAccount.OfficialAccount // 存储公众号实例
}

func (oa *OfficialAccountService) GetApp(c *gin.Context) *officialAccount.OfficialAccount {
	tmp, ok := c.Get(constant.WechatOfficialAccountAppId)
	if !ok || tmp == "" {
		panic("OfficialAccountService GetApp failed: appId is empty")
	}

	appId := tmp.(string)
	if app, ok := oa.apps[appId]; ok && app != nil {
		return app
	}

	app, err := newApp(appId)
	if err != nil {
		panic("OfficialAccountService GetApp error: " + err.Error())
	}

	oa.lock.Lock()
	defer oa.lock.Unlock()
	if len(oa.apps) == 0 { // 初始化
		oa.apps = make(map[string]*officialAccount.OfficialAccount)
	}
	oa.apps[appId] = app

	return app
}

func newApp(appId string) (*officialAccount.OfficialAccount, error) {
	conf := utils.GetOfficialAccountConf(appId)

	var cache kernel.CacheInterface
	// if conf.RedisAddr != "" {
	// 	cache = kernel.NewRedisClient(&kernel.RedisOptions{
	// 		Addr: conf.RedisAddr,
	// 	})
	// }

	app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:        conf.AppId,     // 小程序、公众号或者企业微信的appid
		Secret:       conf.AppSecret, // 商户号 appID
		Token:        conf.MessageToken,
		AESKey:       conf.MessageAesKey,
		ResponseType: os.Getenv("response_type"),
		Log: officialAccount.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
	})

	return app, err
}
