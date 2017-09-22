package main

import (
	"com/ectongs/preplanui/gui"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"github.com/lxn/walk"
	"com/ectongs/preplanui/utils"
	"fmt"
	"os"
	"com/ectongs/preplanui/network"
)

func main() {
	//创建登录窗口
	var logined bool = false
	login := gui.NewLoginDialod(gui.EW_TITLEBAR|gui.EW_MAIN|gui.EW_ENABLE_DEBUG, 500, 300, "登录")
	login.SetFile("views\\login.html")

	login.SetWindowHandlers(func(window *window.Window) {
		window.DefineFunction("login", func(args ...*sciter.Value) *sciter.Value {
			if len(args) != 2 {
				return sciter.NewValue(false)
			}
			fmt.Printf("Username: %s --- Password: %s \n", args[0], args[1])
			if ok := network.Login(args[0].String(), args[1].String()); !ok {
				return sciter.NewValue(false)
			}

			logined = true
			return sciter.NewValue(true)
		})
	})

	login.Show()
	login.Run()

	//未登录成功，退出系统
	if !logined {
		os.Exit(0)
	}

	//创建主窗口
	window, error := walk.NewMainWindow()
	if error != nil {
		utils.MsgBoxWithError(window, INIT_MAIN_WINDOW_FAILED+error.Error())
	}
	//注册窗口
	RegistrMainWindow(window)
	defer ReleaseMainWindow()

	//启动websocket服务
	go network.WebSocketRun()
	//启动message获取服务
	go network.MessageRun()
	//启动弹窗显示服务
	go gui.PopupRun()

	//显示窗口
	ShowWindow(window)

}

func ReleaseMainWindow() {
	//listener1.Stop()
}
