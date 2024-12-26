package bo

import emailpkg "github.com/ikaiguang/go-srv-kit/kit/email"

type SendEmailCodeConfig struct {
	Sender  *emailpkg.Sender
	Message emailpkg.Message
}
