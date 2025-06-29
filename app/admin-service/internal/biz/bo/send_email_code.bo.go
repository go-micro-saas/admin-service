package bo

import (
	emailpkg "github.com/ikaiguang/go-srv-kit/kit/email"
	"time"
)

type SendEmailCodeConfig struct {
	Enable  bool
	Sender  emailpkg.Sender
	Message emailpkg.Message
}

type VerifyCodeConfig struct {
	CaptchaLen int
	CaptchaTTL time.Duration
}
