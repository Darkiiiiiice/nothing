package mq

type MessageCenter struct {
	tipsChannel    chan string
	commandChannel chan string
	confChannel    chan string
	stateChannel   chan int
}

var mc *MessageCenter

func init() {
	mc = &MessageCenter{
		tipsChannel:    make(chan string, 10),
		commandChannel: make(chan string, 10),
		stateChannel:   make(chan int, 10),
		confChannel:    make(chan string),
	}
}

func NotifyTips(prefix string, msg string) {
	mc.tipsChannel <- prefix + msg
}

func NotifyCommand(cmd string) {
	mc.commandChannel <- cmd
}

func NotifyState(state int) {
	mc.stateChannel <- state
}

func NotifyConf(conf string) {
	mc.confChannel <- conf
}

func GetTipsChan() chan string {
	return mc.tipsChannel
}

func GetCommandChan() chan string {
	return mc.commandChannel
}

func GetConfChan() chan string {
	return mc.confChannel
}

func GetStateChan() chan int {
	return mc.stateChannel
}
