package middleware

import (
	"crypto/sha1"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"sort"
	"strings"
)

// 微信公众号验签
func OfficialAccountAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.DefaultQuery("signature", "")
		timestamp := c.DefaultQuery("timestamp", "")
		nonce := c.DefaultQuery("nonce", "")
		echostr := c.DefaultQuery("echostr", "")
		appId := c.DefaultQuery("appId", "")

		err := checkSignature(signature, timestamp, nonce, appId)
		if err != nil {
			response.FailWithDetailed(gin.H{
				"signature": signature,
				"timestamp": timestamp,
				"nonce":     nonce,
				"echostr":   echostr,
				"appId":     appId,
			}, "OfficialAccountAuth failed:"+err.Error(), c)
			c.Abort()
			return
		}

		// 验签通过：暂且认为GET请求就是接入，输出echoStr
		if c.Request.Method == "GET" {
			c.String(http.StatusOK, echostr)
			c.Abort()
			return
		}

		c.Set(constant.WechatOfficialAccountAppId, appId)
		c.Next()
	}
}

// 校验签名
// 接入指南 https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html
func checkSignature(signature, timestamp, nonce, appId string) error {
	conf := utils.GetOfficialAccountConf(appId)
	if conf == nil {
		return errors.New("get appId config error")
	}

	arr := []string{timestamp, nonce, conf.Token}
	sort.Strings(arr)
	str := strings.Join(arr, "")
	res := toSha1(str)
	if res != signature {
		return errors.New("signature error:" + res)
	}
	return nil
}

// 哈希算法
func toSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
