package consts

//状态消息
const (
	USB_CONNECT    = iota
	USB_DISCONNECT

	PHONE_HANGUP

	EM_PHONE_CONNECT
	EM_PHONE_DISCONNECT
	EM_PICKUP
	EM_HANGUP

	EM_CONFIG_UPDATE
	EM_RECORD_START
	EM_RECORD_STOP
)

//命令消息
const (
	EC_HANGUP      = "hangup"
	EC_DIAL        = "dial"
	EC_STARTRECORD = "startrecord"
	EC_STOPRECORD  = "stoprecord"

	EC_UPDATE_CONFIG = "configupdate"

	EC_PHONE_CONNECT    = "PHONE_CONNECT"
	EC_PHONE_DISCONNECT = "PHONE_DISCONNECT"
)

//电话厂家
const (
	DEFAULT   = "DEFAULT"
	HION_U860 = "HIONU860"
)

//服务器配置
const (
	SERVER_ADDR = "http://220.194.42.91:8096"
	SERVER_URL  = "/PcClientServlet"
)

//消息前缀
const (
	TAG_RTIPS = "TIPS://" //右下角消息
	TAG_MTIPS = "MSG://"  //弹窗消息

)
