package network

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"encoding/json"
	"crypto/md5"
	"errors"
	"com/ectongs/preplanui/mq"
	"com/ectongs/preplanui/app"
	"com/ectongs/preplanui/consts"
	"com/ectongs/preplanui/conf"
)

const (
	CODE_SUCCESS = iota
	CODE_FAILED
)

const (
	MSG_SUCCESS = "SUCCESS"
	MSG_FAILED  = "FAILED"

	GETTOKEN = "getToken"
	DIAL     = "dial"
)

//网页发送给服务器的消息
//请求命令时的数据结构
type Request struct {
	Command     string                 //命令类型， dial：拨打电话   hangup：挂断电话   startrecord: 开始录音  stoprecord: 停止录音
	PhoneNumber string                 //电话号码， 只有Command是dial 的时候才需要，其他情况可以为空
	Token       string                 //校验码，在第一次连接的时候会返回一个token，这里每次请求加上那个token就可以了
	Method      string `json:"method"` //请求方法
	//下面两个只在第一次建立WebSocket时，才会有
	Tag string //值恒等于ectongs
	Id  string //随机生成的32位字符串
}

//服务器返回的消息
type Response struct {
	ReturnCode   int    //服务器返回的状态码 0: 请求成功   1： 请求失败
	ReturnMsg    string //服务器返回的状态描述   SUCCESS： 表示成功   FAILED： 表示失败
	RetureDetial string //服务器返回的消息描述
	Method       string `json:"method"`
}

var requestMap map[string]string = make(map[string]string)

func WebSocketRun() {
	http.Handle("/ws", websocket.Handler(OnWebSocket))

	// 开始服务
	if err := http.ListenAndServe(":12138", nil); err != nil {
		fmt.Println("服务失败 /// ", err)
	}
}

func OnWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	var err error
	var str string

	for {
		if err = websocket.Message.Receive(ws, &str); err != nil {
			fmt.Println("OnWebSocket.websocket.Message.Receive:", err.Error())
			break
		}
		fmt.Println("从客户端收到: ", str)

		var req Request = Request{}

		if err := json.Unmarshal([]byte(str), &req); err != nil {
			ResponseFailed(ws, err, "")
			break
		}

		//第一次建立连接
		if req.Method == "getToken" && req.Tag == "ectongs" && req.Id != "" {

			fmt.Println("第一次连接")

			mq.NotifyTips(consts.TAG_RTIPS, "客户端已连接")

			bId := md5.Sum([]byte(req.Id))
			token := fmt.Sprintf("%x", bId)

			requestMap[token] = req.Id
			ResponseSuccess(ws, token, req.Method)

		} else if req.Token != "" { //接收到普通消息
			fmt.Println("接收到普通消息")

			if _, ok := requestMap[req.Token]; ok {

				if !app.GetPhoneState() {
					fmt.Println("电话未连接")
					ResponseFailed(ws, errors.New("电话未连接"), req.Method)
				} else if app.GetTalkState() {
					fmt.Println("电话未连接")
					ResponseFailed(ws, errors.New("正在通话中"), req.Method)
				} else if app.GetRecordState() {
					fmt.Println("电话未连接")
					ResponseFailed(ws, errors.New("正在录音中"), req.Method)
				} else {
					command := fmt.Sprintf("%s:%s%s", req.Command, conf.TeleNo(), req.PhoneNumber)
					fmt.Println("接收到命令", command)
					mq.NotifyCommand(command)
					ResponseSuccess(ws, "", req.Method)
				}
			} else {
				fmt.Println("非法的连接")
				ResponseFailed(ws, errors.New("非法的连接"), req.Method)
			}
		} else {
			fmt.Println("什么都没有")
		}
	}
}

func ResponseFailed(ws *websocket.Conn, err error, method string) {
	res := Response{CODE_FAILED, MSG_FAILED, err.Error(), method}

	str, err := json.Marshal(res)
	if err != nil {
		fmt.Println("ResponseFailed.json.Marshal:", err.Error())
		return
	}

	if err := websocket.Message.Send(ws, string(str)); err != nil {
		fmt.Println("ResponseFailed.websocket.Message.Send:", err.Error())
		return
	}
	fmt.Println("出现错误:", string(str))
}

func ResponseSuccess(ws *websocket.Conn, msg string, method string) {
	res := Response{CODE_SUCCESS, MSG_SUCCESS, msg, method}

	str, err := json.Marshal(res)
	if err != nil {
		fmt.Println("ResponseSuccess.json.Marshal:", err.Error())
		return
	}

	if err := websocket.Message.Send(ws, string(str)); err != nil {
		fmt.Println("ResponseSuccess.websocket.Message.Send:", err.Error())
		return
	}
	fmt.Println("向客户端发送:", string(str))
}
