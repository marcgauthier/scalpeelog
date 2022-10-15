package scalpeelog

import (
	"encoding/json"
	"net"
	"strconv"
	"time"
)

const constInfo = 0
const constWarning = 1
const constError = 2

var app string
var port int
var ip string

type report struct {
	Time       int64
	Importance int // 0=Info, 1=Warning, 2=Error
	App        string
	Msg        string
}

func Init(App, IP string, Port int) {
	app = App
	port = Port
	ip = IP
}

func Info(Msg string) {

	send(Msg, constInfo)
}

func Error(Msg string) {
	send(Msg, constError)

}

func Warning(Msg string) {
	send(Msg, constWarning)

}

func send(msg string, T int) error {

	m := report{time.Now().Unix(), T, app, msg}

	jsonData, _ := json.Marshal(m)

	con, err := net.Dial("udp", ip+":"+strconv.Itoa(port))
	_, err = con.Write(jsonData)
	defer con.Close()

	return err
}
