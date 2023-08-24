package wechat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type OfficialAccountApi struct{}

func (e *OfficialAccountApi) Index(c *gin.Context) {
	response.Ok(c)
}

func (e *OfficialAccountApi) IndexPost(c *gin.Context) {

	fmt.Println(global.GVA_CONFIG.OfficialAccountList)
	fmt.Println(c.Get("wxOAAppId"))
	response.OkWithMessage("index success", c)

}

//
// func (e *CustomApi) IndexPost(c *gin.Context) {
// 	signature := c.DefaultQuery("signature", "")
// 	timestamp := c.DefaultQuery("timestamp", "")
// 	nonce := c.DefaultQuery("nonce", "")
// 	if !checkSignature(signature, timestamp, nonce) {
// 		c.String(500, "sign error")
// 	}
// 	//
// 	//bd, _ := ioutil.ReadAll(c.Request.Body)
// 	//log.Println(string(bd))
//
// 	services.OfficialAccountApp, _ = services.NewOfficialAccountAppService(&config.Configuration{
// 		OffiAccount: config.OffiAccount{
// 			AppID:         "",
// 			AppSecret:     "",
// 			RedisAddr:     "localhost:6379",
// 			MessageToken:  "",
// 			MessageAesKey: "",
// 		},
// 	})
//
// 	rs, err := services.OfficialAccountApp.Server.Notify(c.Request, func(event contract.EventInterface) interface{} {
// 		fmt.Println("event", event)
//
// 		// 这里需要获取到事件类型，然后把对应的结构体传递进去进一步解析
// 		// 所有包含的结构体请参考： https://github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server/handlers/models
// 		switch event.GetMsgType() {
// 		case models.CALLBACK_MSG_TYPE_TEXT:
// 			msg := models2.MessageText{}
// 			err := event.ReadMessage(&msg)
// 			if err != nil {
// 				println(err.Error())
// 				return "error"
// 			}
// 			fmt2.Dump(msg)
// 		}
//
// 		// 假设用户给应用发送消息，这里可以直接回复消息文本，
// 		// return  "I'm recv..."
//
// 		// 这里回复success告诉微信我收到了，后续需要回复用户信息可以主动调发消息接口
// 		//return kernel.SUCCESS_EMPTY_RESPONSE
//
// 		// 如果要返回xml，也可以返回message对象
// 		return messages.NewText("北京时间" + time.Now().Format("2006-01-02 15:04:05"))
//
// 	})
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// 选择1： 直接把gin context writer传入，会自动回复。
// 	err = helper.HttpResponseSend(rs, c.Writer)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return
//
// 	// 选择2： 或者是把内容读取出来
// 	//text, _ := ioutil.ReadAll(rs.Body)
// 	//c.String(http.StatusOK, string(text))
// 	//return
//
// }
