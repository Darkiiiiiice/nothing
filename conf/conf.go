package conf

import (
	"com/ectongs/preplanui/consts"
	"reflect"
	"strings"
	"com/ectongs/preplanui/usb"
	"com/ectongs/preplanui/app"
	"com/ectongs/preplanui/mq"
	"fmt"
)

type Config struct {
	PhoneType      string
	ServerAddress  string
	TeleNo         string
	SessionId      string
	LoginUser      string
	RecordFilePath string
}

var conf *Config

func init() {
	conf = new(Config)
	conf.PhoneType = consts.HION_U860
	conf.TeleNo = "100"
	conf.RecordFilePath = "./"
	conf.ServerAddress = consts.SERVER_ADDR + consts.SERVER_URL
}

func SaveConf(cfg string) {
	cfgs := strings.Split(cfg, "(|=|)")
	for i, v := range cfgs {
		mutable := reflect.ValueOf(conf).Elem()
		mutable.Field(i).SetString(v)
	}
}

func GetValue(name string) string {
	immutable := reflect.ValueOf(conf).Elem()
	return immutable.FieldByName(name).String()
}

func PhoneType() string {
	return conf.PhoneType
}

func ServerAddress() string {
	return conf.ServerAddress
}

func SetLoginUser(user string) {
	conf.LoginUser = user
}
func LoginUser() string {
	return conf.LoginUser
}

func SetSessionId(id string) {
	conf.SessionId = id
}
func SessionId() string {
	return conf.SessionId
}

func SetTeleNo(id string) {
	conf.SessionId = id
}
func TeleNo() string {
	return conf.SessionId
}

//让程序需的新配置生效
func ConfUpdate() {
	//重新生成UsbConn
	usb.NewUsbConn(conf.PhoneType)
	conf.RecordFilePath = "./"
	app.SetConfigState(false)

	mq.NotifyTips(consts.TAG_RTIPS, "配置文件已生效")
}

func PrintAll() {
	fmt.Printf(" PhoneType:%s ServerAddress:%s TeleNo:%s RecordFilePath:%s LoginUser:%s SessionId:%s\n", conf.PhoneType, conf.ServerAddress, conf.TeleNo, conf.RecordFilePath, conf.LoginUser, conf.SessionId)
}
