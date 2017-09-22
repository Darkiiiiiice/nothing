package app

import (
	"sync"
	"fmt"
)

/*
	这里保存程序的状态，比如USB连接状态等
	true 表示连接, false 表示关闭
 */
var (
	phone     bool //电话连接状态
	talking   bool //通话状态
	recording bool //录音状态

	config bool //配置更新

	lock sync.Mutex
)

func SetPhoneState(state bool) {
	lock.Lock()
	phone = state
	lock.Unlock()
}

func GetPhoneState() bool {
	return phone
}

func SetTalkState(state bool) {
	lock.Lock()
	talking = state
	lock.Unlock()
}

func GetTalkState() bool {
	return talking
}

func SetRecordState(state bool) {
	lock.Lock()
	recording = state
	lock.Unlock()
}

func GetRecordState() bool {
	return recording
}

func SetConfigState(state bool) {
	lock.Lock()
	config = state
	lock.Unlock()
}

func GetConfigState() bool {
	return config
}

func PrintAll() {
	fmt.Printf(" Phone:%v Talking:%v Record:%v Conf:%v\n", phone, talking, recording, config)
}
