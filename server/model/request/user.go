package request

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type Register struct {
	Username  string `json:"username" validate:"required,min=6,max=64"`
	Password  string `json:"password" validate:"required,min=8,max=64"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
