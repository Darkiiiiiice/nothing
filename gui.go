package main

import (
	"github.com/lxn/walk"
	"com/ectongs/preplanui/utils"
	"com/ectongs/preplanui/listener"
	"com/ectongs/preplanui/conf"
)

func RegistrMainWindow(window *walk.MainWindow) {
	// 从文件中读取图标
	icon, err := walk.NewIconFromFile(MAIN_ICON)
	if err != nil {
		utils.MsgBoxWithError(window, LOAD_ICON_FAILED+err.Error())
	}

	// 创建右下角托盘
	ni, err := walk.NewNotifyIcon()
	if err != nil {
		utils.MsgBoxWithError(window, LOAD_NOTIFY_ICON_FAILED+err.Error())
	}

	// 设置右下角图标和工具Tip
	if err := ni.SetIcon(icon); err != nil {
		utils.MsgBoxWithError(window, SET_NOTIFY_ICON_FAILED+err.Error())
	}
	if err := ni.SetToolTip(TOOL_TIP); err != nil {
		utils.MsgBoxWithError(window, SET_TOOL_TIP_FAILED+err.Error())
	}

	// 设置左键点击时出现的信息
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(TITLE, conf.LoginUser()+CUSTOM_INFO); err != nil {
			utils.MsgBoxWithError(window, COMMON_ERROR+err.Error())
		}
	})

	// 创建右键退出菜单
	exitAction := walk.NewAction()
	if err := exitAction.SetText(EXIT); err != nil {
		utils.MsgBoxWithError(window, COMMON_ERROR+err.Error())
	}
	exitAction.Triggered().Attach(func() {
		walk.App().Exit(0)
	})

	//创建右键设置菜单
	confAction := walk.NewAction()
	if err := confAction.SetText(CONFIG); err != nil {
		utils.MsgBoxWithError(window, COMMON_ERROR+err.Error())
	}
	confAction.Triggered().Attach(func() {
		OpenSettingDialog()
	})

	if err := ni.ContextMenu().Actions().Add(confAction); err != nil {
		utils.MsgBoxWithError(window, COMMON_ERROR+err.Error())
	}
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		utils.MsgBoxWithError(window, COMMON_ERROR+err.Error())
	}

	// T将右下角的托盘设置为可见
	if err := ni.SetVisible(true); err != nil {
		utils.MsgBoxWithError(window, COMMON_ERROR+err.Error())
	}

	//注册监听器
	ls := listener.NewListener(window, ni)
	ls.Run()
}

func ShowWindow(window *walk.MainWindow) {

	//启动主窗口
	window.Run()
}
