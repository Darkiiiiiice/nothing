package listener

import (
	"github.com/lxn/walk"
	"context"
	"fmt"
	"com/ectongs/preplanui/mq"
	"com/ectongs/preplanui/conf"
	"com/ectongs/preplanui/consts"
	"com/ectongs/preplanui/app"
	"strings"
	"com/ectongs/preplanui/usb"
	"com/ectongs/preplanui/utils"
	"time"
	"unsafe"
	"log"
	"com/ectongs/preplanui/gui"
)

const delay = 1000

type Listener struct {
	run   bool //是否启动监听
	pause bool //是否暂停监听

	ctx    context.Context
	cancle context.CancelFunc

	window *walk.MainWindow
	icon   *walk.NotifyIcon
}

func NewListener(window *walk.MainWindow, icon *walk.NotifyIcon) *Listener {
	ctx, cancle := context.WithCancel(context.Background())
	return &Listener{false, false, ctx, cancle, window, icon}
}

//配置监听器
func (l *Listener) ConfListener(ctx context.Context) {
	fmt.Println("启动 --- 配置监听器")
	confCh := mq.GetConfChan()
	for {
		//fmt.Println("配置监听中")
		conf.PrintAll()
		time.Sleep(time.Duration(time.Millisecond * delay))

		select {
		case <-ctx.Done():
			fmt.Println("停止 --- 配置监听器")
			return
		case cfg := <-confCh:
			conf.SaveConf(cfg)
			mq.NotifyState(consts.EM_CONFIG_UPDATE)
		default:
		}
	}
}

//状态监听器
func (l *Listener) StateListener(ctx context.Context) {
	fmt.Println("启动 --- 状态监听器")
	stateCh := mq.GetStateChan()
	for {
		//fmt.Println("状态监听中")
		app.PrintAll()

		time.Sleep(time.Duration(time.Millisecond * delay))
		select {
		case <-ctx.Done():
			fmt.Println("停止 --- 状态监听器")
			return
		case state := <-stateCh:
			dealstate(state)
		default:
		}

		if !app.GetPhoneState() {
			mq.NotifyCommand(consts.EC_PHONE_CONNECT)
		}

		if app.GetConfigState() {
			mq.NotifyCommand(consts.EC_UPDATE_CONFIG)
		}

	}
}

//命令监听器
func (l *Listener) CommandListener(ctx context.Context) {
	fmt.Println("启动 --- 命令监听器")
	cmdCh := mq.GetCommandChan()
	for {
		//fmt.Println("命令监听中")
		time.Sleep(time.Duration(time.Millisecond * delay))
		select {
		case <-ctx.Done():
			fmt.Println("停止 --- 命令监听器")
			return
		case cmd := <-cmdCh:
			dealcmd(cmd, l.window)
		default:
		}
	}
}

//弹窗监听器
func (l *Listener) TipsListener(ctx context.Context) {
	fmt.Println("启动 --- 弹窗监听器")
	tipCh := mq.GetTipsChan()
	for {
		//fmt.Println("弹窗监听中")
		time.Sleep(time.Duration(time.Millisecond * delay))
		select {
		case <-ctx.Done():
			fmt.Println("停止 --- 弹窗监听器")
			return
		case tip := <-tipCh:
			dealtips(tip, l.icon)
		default:
		}
	}
}

//暂停
func (l *Listener) Pause() {
	l.pause = true
}

//恢复
func (l *Listener) Recover() {
	l.pause = false
}

func (l *Listener) Run() {
	if !l.run {
		fmt.Println("监听器已启动")
		go l.ConfListener(l.ctx)
		go l.StateListener(l.ctx)
		go l.CommandListener(l.ctx)
		go l.TipsListener(l.ctx)
		l.run = true
	}
}

func (l *Listener) Stop() {
	if l.run && l.cancle != nil {
		fmt.Println("监听器已停止")
		l.cancle()
		l.run = false
	}
}

//处理状态
func dealstate(state int) {
	switch state {
	case consts.EM_PHONE_CONNECT:
		//fmt.Println("状态监听器 电话 连接")
		app.SetPhoneState(true)
	case consts.EM_PHONE_DISCONNECT:
		//fmt.Println("状态监听器 电话 断开")
		app.SetPhoneState(false)
	case consts.EM_HANGUP:
		app.SetTalkState(false)
		//fmt.Println("状态监听器 电话 挂断")
	case consts.EM_PICKUP:
		app.SetTalkState(true)
		//fmt.Println("状态监听器 电话 通话")
	case consts.EM_CONFIG_UPDATE:
		//fmt.Println("状态监听器 配置 更新")
		app.SetConfigState(true)
		mq.NotifyCommand(consts.EC_UPDATE_CONFIG)
	case consts.EM_RECORD_START:
		//fmt.Println("状态监听器 录音 开启")
		app.SetRecordState(true)
	case consts.EM_RECORD_STOP:
		//fmt.Println("状态监听器 录音 关闭")
		app.SetRecordState(false)
	}
}

//处理命令
func dealcmd(cmd string, window *walk.MainWindow) {
	cmds := strings.Split(cmd, ":")
	switch cmds[0] {
	case consts.EC_HANGUP:
		if ok := usb.HangUp(); ok {
			mq.NotifyState(consts.EM_HANGUP)
			mq.NotifyTips(consts.TAG_RTIPS, "电话已挂断")
		} else {
			utils.MsgBoxWithWarning(nil, "无法自动挂断电话，请检查电话连线")
		}
	case consts.EC_DIAL:
		if ok := usb.Dial(cmds[1]); ok {
			fmt.Println("正在拨打电话。。。")
			mq.NotifyTips(consts.TAG_RTIPS, "正在拨打电话。。。 "+cmds[1])
		} else {
			utils.MsgBoxWithWarning(nil, "无法自动拨打电话，请检查电话连线")
		}
	case consts.EC_STARTRECORD:
		if ok := usb.StartRecord(conf.GetValue("RecordFilePath"), usb.TALKING_RECORD); ok {
			mq.NotifyState(consts.EM_RECORD_START)
			mq.NotifyTips(consts.TAG_RTIPS, "通话录音已开启 ")
		} else {
			utils.MsgBoxWithWarning(nil, "无法开启通话录音，请检查电话连线")
		}
	case consts.EC_STOPRECORD:
		if ok := usb.StopRecord(); ok {
			mq.NotifyState(consts.EM_RECORD_STOP)
			mq.NotifyTips(consts.TAG_RTIPS, "通话录音已停止")
		} else {
			utils.MsgBoxWithWarning(nil, "无法停止通话录音,请检查电话连线")
		}
	case consts.EC_UPDATE_CONFIG:
		conf.ConfUpdate()
	case consts.EC_PHONE_CONNECT:
		//fmt.Println("正在连接电话")
		if ok := usb.BindWindow(unsafe.Pointer(window.Handle())); ok {
			mq.NotifyState(consts.EM_PHONE_CONNECT)
			mq.NotifyTips(consts.TAG_RTIPS, "电话已连接")
		}
	case consts.EC_PHONE_DISCONNECT:
		//fmt.Println("正在断开电话")
		if ok := usb.UnBindWindow(); ok {
			mq.NotifyState(consts.EM_PHONE_DISCONNECT)
			mq.NotifyTips(consts.TAG_RTIPS, "电话已断开")
		}
	}
}

func dealtips(tip string, icon *walk.NotifyIcon) {
	if strings.HasPrefix(tip, consts.TAG_MTIPS) {
		tip = strings.TrimLeft(tip, consts.TAG_MTIPS)
		gui.SendMsgToPopupWindow(tip)
	} else if strings.HasPrefix(tip, consts.TAG_RTIPS) {
		tip = strings.TrimLeft(tip, consts.TAG_RTIPS)
		if err := icon.ShowInfo("应急指挥通知", tip); err != nil {
			log.Fatal(err)
		}
	}
}
