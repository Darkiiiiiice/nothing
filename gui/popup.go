package gui

import (
	"fmt"
	"github.com/sciter-sdk/go-sciter/window"
	"github.com/sciter-sdk/go-sciter"
	"com/ectongs/preplanui/network"
	"encoding/json"
)

var pw *Window

var popChan chan string = make(chan string, 10)

var row string = "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>"

func PopupRun() {
	pop := NewLoginDialod(EW_TOOL|EW_POPUP|EW_RESIZEABLE, 500, 300, "弹窗")
	pop.SetFile("views\\popup.html")

	pop.SetWindowHandlers(func(w *window.Window) {
		w.DefineFunction("Shown", func(args ...*sciter.Value) *sciter.Value {
			var msg string
			select {
			case msg = <-popChan:
			default:
			}
			return sciter.NewValue(msg)
		})
	})

	pop.Show()

	pw = pop
	pop.Run()
}

func SendMsgToPopupWindow(msg string) {
	if pw == nil || pw.window == nil {
		return
	}
	m := new(network.Message)
	if err := json.Unmarshal([]byte(msg), m); err != nil {
		fmt.Println("gui.drawTableBody.json.Unmarshal:", err.Error())
		return
	}

	html := fmt.Sprintf(row, m.Title, m.Type, m.Msg, m.MsgTime)

	fmt.Println("=======================================================================================================================")
	fmt.Println(html)
	fmt.Println("=======================================================================================================================")
	popChan <- html

}

func drawTableBody(msg string, elem *sciter.Element) bool {
	m := new(network.Message)
	if err := json.Unmarshal([]byte(msg), m); err != nil {
		fmt.Println("gui.drawTableBody.json.Unmarshal:", err.Error())
		return false
	}

	titleElem := elem.MustSelectById("title")
	titleElem.SetText(m.Title)

	tpElem := elem.MustSelectById("type")
	tpElem.SetText(m.Type)

	msgElem := elem.MustSelectById("msg")
	msgElem.SetText(m.Msg)

	timeElem := elem.MustSelectById("time")
	timeElem.SetText(m.MsgTime)

	return true
}
