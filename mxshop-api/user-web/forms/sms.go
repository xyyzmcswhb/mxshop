package forms

type SendSmsForm struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required,mobile"` //手机号码规范可循，需自定义validator
	Type   uint   `json:"type" form:"type" binding:"required,oneof=1 2"`  //可能不同地方都有短信发送的需求：1.注册发送 2.动态验证码登陆发送
}
