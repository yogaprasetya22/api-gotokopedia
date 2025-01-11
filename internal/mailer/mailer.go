package mailer

import (
	"embed"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data interface{}, isSendBox bool) error
}
