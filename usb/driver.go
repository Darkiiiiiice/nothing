package usb

import (
	"unsafe"
	"com/ectongs/preplanui/consts"
	"fmt"
)

//Usb驱动接口
type UsbDriver interface {
	BindWindow(hwnd unsafe.Pointer) bool
	UnBindWindow() bool

	Dial(number string) bool
	HangUp() bool

	StartRecord(recordFilePath string, recordType int) bool
	StopRecord() bool
}

var driver UsbDriver = NewHionU860Driver()

var conn *UsbConn

func NewUsbConn(tag string) {
	switch tag {
	case consts.HION_U860, consts.DEFAULT:
		fmt.Println("生成新的UsbConn --- ", tag)
		conn = &UsbConn{
			driver: NewHionU860Driver(),
			name:   tag,
		}
	default:

		conn = &UsbConn{
			driver: NewHionU860Driver(),
			name:   tag,
		}
	}
}

//与USB端口交互的连接
type UsbConn struct {
	driver UsbDriver
	name   string
}

func (u *UsbConn) BindWindow(hwnd unsafe.Pointer) bool {
	return u.driver.BindWindow(hwnd)
}

func (u *UsbConn) UnBindWindow() bool {
	return u.driver.UnBindWindow()
}

func (u *UsbConn) Dial(number string) bool {
	return u.driver.Dial(number)
}

func (u *UsbConn) HangUp() bool {
	return u.driver.HangUp()
}

func (u *UsbConn) StartRecord(recordFilePath string, recordType int) bool {
	return u.driver.StartRecord(recordFilePath, recordType)
}

func (u *UsbConn) StopRecord() bool {
	return u.driver.StopRecord()
}

/*
	下面属于usb 包的静态方法,在不自己创建UsbConn的情况下，推荐使用静态方法
 */

func BindWindow(hwnd unsafe.Pointer) bool {
	if conn == nil {
		NewUsbConn(consts.HION_U860)
	}
	return conn.BindWindow(hwnd)
}

func UnBindWindow() bool {
	if ok := conn.UnBindWindow(); ok {
		conn = nil
		return true
	}
	return false
}

func Dial(number string) bool {
	return conn.Dial(number)
}

func HangUp() bool {
	return conn.HangUp()
}

func StartRecord(recordFilePath string, recordType int) bool {
	return conn.StartRecord(recordFilePath, recordType)
}

func StopRecord() bool {
	return conn.StopRecord()
}
