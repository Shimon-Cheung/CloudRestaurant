package controller

import (
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engin *gin.Engine) {
	engin.GET("/api/sendcode", mc.sendEmailCode)
}

// http://localhost:8090/api/sendcode?email=25337942@qq.com
func (mc *MemberController) sendEmailCode(context *gin.Context) {
	// 发送验证码功能
	email, ok := context.GetQuery("email")
	if !ok {
		// 如果没有获取到邮箱
		context.JSON(200, gin.H{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	// 开始发送验证码
	memberService := service.MemberService{}
	isSend := memberService.SendCode(email)

	if isSend {
		context.JSON(200, gin.H{
			"code": 1,
			"msg":  "发送成功",
		})
		return
	}

	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "发送失败",
	})

}
