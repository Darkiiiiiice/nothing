package usblib

//#cgo CFLAGS : -I ./include
//#cgo LDFLAGS: -L ./lib -lUsbDll
// #include <stdio.h>
// #include <stdlib.h>
// #include <UsbDll.h>
/*
//// 若无说明，则返回 0 代表成功
////

////初始化，仅调用一次,最先调用
int CInitDll(){
	return InitDll();
}
////绑定窗口，所有事件发送到此窗口，初始化时调用；系统接收消息 WM_DEVICECHANGE 后调用此函数
int CBindWindow(void *hwnd){
	return BindWindow(hwnd);
}
////解除窗口绑定, 则所有事件将不会发送到此窗口，窗口销毁时调用此函数
int CUnBindWindow(){
	return UnBindWindow();
}

////发送摘机命令
int COffHookCtrl(){
	return OffHookCtrl(0);
}

////发送挂机命令
int CHangUpCtrl() {
	return HangUpCtrl(0);
}

////发送拨号命令
int CStartDial(const char *szDest){
	return StartDial(0,szDest);
}

////设置振铃开或关,mode: 0--关闭; 1--打开
int CBell(char mode){
	return Bell(0,mode);
}

////与StartDial功能相似,供二次拨号用
int CSendDTMF(const char *szDest){
	return SendDTMF(0,szDest);
}

////查询话机摘挂机状态 0->挂机,1->摘机
int CQueryPhoneStatus() {
	return QueryPhoneStatus(0);
}

////本地录放音，开 pc 时， rec: false--关闭;true--打开
int CSetLocalRecord(int rec) {
	return setLocalRecord(0,rec);
}

////通话录音，rec: false--关闭;true--打开
int CSetTalkRecord(int rec){
	return setTalkRecord(0,rec);
}

////留言录音，rec: false--关闭;true--打开
int CSetLeaveRecord(int rec) {
	return setLeaveRecord(0,rec);
}

////获取序列号
int CGetSerialNo(char *seq) {
	return GetSerialNo(0,seq);
}

////设置序列号
int CSetSerialNo(const char *seq) {
	return SetSerialNo(0,seq);
}

////闪断一下，ivalue--Flash操作的时间长度,取值为0--100ms,1--180ms,2--300ms,3--600ms,4--1000ms之间。
int CFlash(unsigned int ivalue){
	return Flash(0,ivalue);
}

////设置拨号音开或关,mode: 0--关闭; 1--打开
int CSetDialTone(char mode) {
	return SetDialTone(0,mode);
}

////设置自动接听开或关,mode: 0--关闭; 1--打开
int CSetAutoAnswer(char mode) {
	return SetAutoAnswer(0,mode);
}

////设置Flash值, ivalue 取值为0--100ms,1--180ms,2--300ms,3--600ms,4--1000ms之间
int CSetFlashTime(unsigned int ivalue){
	return SetFlashTime(0,ivalue);
}

////设置出局码，最多3位
int CSetOutcode(const char * code) {
	return SetOutcode(0,code);
}

////开始录音操作, strFileName: 录音文件名，完整的路径
////如:"C:\\record\\sound.wav"。iType: 录音类型：0:本地录音；1:通话录音；2:留言录音
int CStartRecordFile(const char* strFileName, int iType) {
	return StartRecordFile(0,strFileName, iType);
}

////停止录音
int CStopRecordFile() {
	return StopRecordFile(0);
}

////转拨闪断一下，ivalue--转拨操作的时间长度,取值为0--100ms,1--180ms,2--300ms,3--600ms,4--1000ms之间。
int CZhuanBo(unsigned int ivalue) {
	return ZhuanBo(0,ivalue);
}

////bOn--1:开启保留；0:关闭保留
int CHold(int bOn) {
	return Hold(0,bOn);
}

//// bOn->1:开启闭音；0:关闭闭音；
int CMute(int bOn) {
	return Mute(0,bOn);
}
*/
import "C"
import "unsafe"

const (
	CLOSE = 0
	OPEN  = 1

	SUCCESS = 0

	HANGUP = 0
	PICKUP = 1
)

func InitDll() bool {
	r := int(C.CInitDll())

	if r == SUCCESS {
		return true
	}

	return false
}

func BindWindow(hwnd unsafe.Pointer) bool {
	r := int(C.CBindWindow(hwnd))
	if r == SUCCESS {
		return true
	}

	return false
}

func UnBindWindow() bool {
	r := int(C.CUnBindWindow())
	if r == SUCCESS {
		return true
	}

	return false
}

func OffHookCtrl() bool {
	r := int(C.COffHookCtrl())
	if r == SUCCESS {
		return true
	}

	return false
}

func HangUpCtrl() bool {
	r := int(C.CHangUpCtrl())
	if r == SUCCESS {
		return true
	}

	return false
}

func StartDial(szDest string) bool {
	cs := C.CString(szDest)
	r := int(C.CStartDial(cs))
	C.free(unsafe.Pointer(cs))
	if r == SUCCESS {
		return true
	}

	return false
}

func Bell(mode byte) bool {
	c := C.char(mode)
	r := int(C.CBell(c))
	if r == SUCCESS {
		return true
	}

	return false
}

func SendDTMF(szDest string) bool {
	cs := C.CString(szDest)
	r := int(C.CStartDial(cs))
	C.free(unsafe.Pointer(cs))
	if r == SUCCESS {
		return true
	}

	return false
}

func QueryPhoneStatus() int {
	r := int(C.CQueryPhoneStatus())
	if r == HANGUP {
		return HANGUP
	} else if r == PICKUP {
		return PICKUP
	}

	return -1
}

func SetLocalRecord(rec byte) bool {
	i := C.int(int(rec))
	r := int(C.CSetLocalRecord(i))
	if r == SUCCESS {
		return true
	}

	return false
}

func SetTalkRecord(rec byte) bool {
	i := C.int(int(rec))
	r := int(C.CSetTalkRecord(i))
	if r == SUCCESS {
		return true
	}

	return false
}

func SetLeaveRecord(rec byte) bool {
	i := C.int(int(rec))
	r := int(C.CSetLeaveRecord(i))
	if r == SUCCESS {
		return true
	}

	return false
}

// Deprecated: Have some problem..
func GetSerialNo() string {
	cs := C.CString("")
	r := int(C.CGetSerialNo(cs))

	gs := C.GoString(cs)

	C.free(unsafe.Pointer(cs))
	if r == SUCCESS {
		return gs
	}

	return ""
}

func SetSerialNo(seq string) bool {
	cs := C.CString(seq)
	r := int(C.CSetSerialNo(cs))
	C.free(unsafe.Pointer(cs))
	if r == SUCCESS {
		return true
	}

	return false
}

func Flash(ivalue uint32) bool {
	i := C.uint(ivalue)
	r := int(C.CFlash(i))
	if r == SUCCESS {
		return true
	}

	return false
}

func SetDialTone(mode byte) bool {
	c := C.char(mode)
	r := int(C.CSetDialTone(c))
	if r == SUCCESS {
		return true
	}

	return false
}

func SetAutoAnswer(mode byte) bool {
	c := C.char(mode)
	r := int(C.CSetAutoAnswer(c))
	if r == SUCCESS {
		return true
	}

	return false
}

func SetFlashTime(ivalue uint32) bool {
	i := C.uint(ivalue)
	r := int(C.CSetFlashTime(i))
	if r == SUCCESS {
		return true
	}

	return false
}

func SetOutcode(seq string) bool {
	cs := C.CString(seq)
	r := int(C.CSetOutcode(cs))
	C.free(unsafe.Pointer(cs))
	if r == SUCCESS {
		return true
	}

	return false
}

func StartRecordFile(strFileName string, iType int) bool {
	cs := C.CString(strFileName)
	i := C.int(iType)
	r := int(C.CStartRecordFile(cs, i))
	C.free(unsafe.Pointer(cs))
	if r == SUCCESS {
		return true
	}

	return false
}

func StopRecordFile() bool {
	r := int(C.CStopRecordFile())
	if r == SUCCESS {
		return true
	}

	return false
}

func ZhuanBo(ivalue uint32) bool {
	i := C.uint(ivalue)
	r := int(C.CZhuanBo(i))
	if r == SUCCESS {
		return true
	}

	return false
}

func Hold(bOn int) bool {
	i := C.int(bOn)
	r := int(C.CHold(i))
	if r == SUCCESS {
		return true
	}
	return false
}

// bOn->1:开启闭音；0:关闭闭音；
func Mute(bOn int) bool {
	i := C.int(bOn)
	r := int(C.CMute(i))
	if r == SUCCESS {
		return true
	}
	return false
}
