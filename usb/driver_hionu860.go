package usb

import (
	"com/ectongs/preplanui/usblib"
	"unsafe"
	"log"
)

const (
	LOCAL_RECORD   = iota //本地录音
	TALKING_RECORD        //通话录音
	MESSAGE_RECORD        //留言录音
)

type HionU860Driver struct {
	tag string
}

func init() {
	if ok := usblib.InitDll(); !ok {
		log.Fatalln("USB初始化出错")
	}
}

func NewHionU860Driver() *HionU860Driver {
	return &HionU860Driver{"HionU860"}
}

func (u *HionU860Driver) BindWindow(hwnd unsafe.Pointer) bool {
	return usblib.BindWindow(hwnd)
}
func (u *HionU860Driver) UnBindWindow() bool {
	return usblib.UnBindWindow()
}

func (u *HionU860Driver) Dial(number string) bool {
	return usblib.StartDial(number)
}
func (u *HionU860Driver) HangUp() bool {
	return usblib.HangUpCtrl()
}

func (u *HionU860Driver) StartRecord(recordFilePath string, recordType int) bool {
	return usblib.StartRecordFile(recordFilePath, recordType)
}

func (u *HionU860Driver) StopRecord() bool {
	return usblib.StopRecordFile()
}
