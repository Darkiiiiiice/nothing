package network

import (
	"net/url"
	"com/ectongs/preplanui/conf"
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/xml"
	"com/ectongs/preplanui/mq"
	"com/ectongs/preplanui/consts"
	"encoding/json"
)

var delay time.Duration = time.Second * 10

type Messages struct {
	ReturnCode     int       `xml:"returnCode",json:"return_code"`
	ReturnErrorMsg string    `xml:"returnErrorMsg",json:"return_error_msg"`
	Message        []Message `xml:"Message",json:"message"`
}

type Message struct {
	Title   string `xml:"title",json:"title"`
	Type    string `xml:"type",json:"type"`
	Msg     string `xml:"msg",json:"msg"`
	MsgTime string `xml:"MsgTime",json:"msg_time"`
}

func MessageRun() {

	for {
		postUrl := conf.ServerAddress()
		time.Sleep(delay)
		val := url.Values{
			"method":    {"getMessage"},
			"sessionId": {conf.SessionId()},
		}
		fmt.Println()
		fmt.Printf("Message捕获 --- Addr: %s SessionId: %s\n", postUrl, conf.SessionId())

		if conf.SessionId() == "" {
			continue
		}

		resp, err := http.PostForm(postUrl, val)
		if err != nil {
			fmt.Println("network.Login.http.PostForm:", err.Error())
			continue
		}
		defer func() {
			if resp != nil {
				resp.Body.Close()
			}
		}()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("network.Login.ioutil.ReadAll:", err.Error())
		}

		fmt.Println("消息捕获：", string(b))

		msg := new(Messages)
		if err = xml.Unmarshal(b, msg); err != nil {
			fmt.Println("network.Login.xml.Unmarshal:", err.Error())
		}

		if msg.ReturnCode == 1 {
			for _, v := range msg.Message {
				jsonStr, err := json.Marshal(v)
				if err != nil {
					fmt.Println("network.Login.json.Marshal:", err.Error())
				}
				mq.NotifyTips(consts.TAG_MTIPS, string(jsonStr))
			}

		}

	}

}
