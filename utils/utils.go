package utils

import (
	"github.com/lxn/walk"
)

const (
	MSG_ERROR       = "错误"
	MSG_WARN        = "警告"
	MSG_INFORMATION = "信息"
)

func MsgBoxWithError(form walk.Form, msg string) {
	walk.MsgBox(form, MSG_ERROR, msg, walk.MsgBoxIconError)
	walk.App().Exit(-1)
}

func MsgBoxWithInformation(form walk.Form, msg string) {
	walk.MsgBox(form, MSG_INFORMATION, msg, walk.MsgBoxIconInformation)
}

func MsgBoxWithWarning(form walk.Form, msg string) {
	walk.MsgBox(form, MSG_WARN, msg, walk.MsgBoxIconWarning)
}
