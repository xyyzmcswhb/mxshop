package forms

//用户登录表单验证

type PasswordLoginForm struct {
	Mobile    string `json:"mobile" form:"mobile" binding:"required,mobile"` //手机号码规范可循，需自定义validator
	Password  string `json:"password" form:"password" binding:"required,min=6,max=20"`
	Captcha   string `json:"captcha" form:"captcha" binding:"required,min=5,max=5"` //输入的图片验证码
	CaptchaId string `json:"captcha_id" form:"captcha_id" binding:"required"`
}

// 用户注册表单验证
type RegisterForm struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required,mobile"` //手机号码规范可循，需自定义validator
	Password string `json:"password" form:"password" binding:"required,min=6,max=20"`
	Code     string `json:"code" form:"code" binding:"required,min=6,max=6"` //输入的验证码
}
