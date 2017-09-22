package network

import (
	"net/http"
	"net/url"
	"com/ectongs/preplanui/conf"
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"com/ectongs/preplanui/mq"
	"com/ectongs/preplanui/consts"
)

type LoginInfo struct {
	ReturnCode     int    `xml:"returnCode"`
	ReturnErrorMsg string `xml:"returnErrorMsg"`

	SchemeId        string `xml:"schemeId"`
	UsedInterfaceId string `xml:"usedInterfaceId"`
	UserId          string `xml:"userId"`
	UserType        string `xml:"userType"`
	UserName        string `xml:"userName"`
	StationId       string `xml:"stationId"`
	StationName     string `xml:"stationName"`
	StationAtt      string `xml:"stationAtt"`
	OtherParams     string `xml:"otherParams"`
	SessionId       string `xml:"sessionId"`
	ServerTime      string `xml:"serverTime"`
}

func Login(userid string, password string) bool {

	val := url.Values{
		"method":   {"login"},
		"userId":   {userid},
		"password": {password},
	}
	postUrl := conf.ServerAddress()

	fmt.Printf("Login--- Addr: %s\n", postUrl)
	fmt.Printf("Login--- User: %s\n", userid)
	fmt.Printf("Login--- Passwd: %s\n", password)

	resp, err := http.PostForm(postUrl, val)
	if err != nil {
		fmt.Println("network.Login.http.PostForm:", err.Error())
		return false
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("network.Login.ioutil.ReadAll:", err.Error())
		return false
	}
	fmt.Println(string(b))

	loginInfo := new(LoginInfo)
	if err = xml.Unmarshal(b, loginInfo); err != nil {
		fmt.Println("network.Login.xml.Unmarshal:", err.Error())
		return false
	}

	if loginInfo.ReturnCode != 1 {
		return false
	}

	conf.SetLoginUser(loginInfo.UserName)
	conf.SetSessionId(loginInfo.SessionId)

	mq.NotifyTips(consts.TAG_RTIPS, loginInfo.UserName+"已登录")

	return true
}
