package service

import (
	"CloudRestaurant/tool"
	"fmt"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) SendCode(email string) bool {
	// 生成一个随机验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	// 调用api进行发送验证码
	err := tool.SendMail(email, "【云餐厅】验证码服务", fmt.Sprintf("<h1>验证码:%s</h1>", code))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	//接受返回结果，并判断发送状态

	return true

}
