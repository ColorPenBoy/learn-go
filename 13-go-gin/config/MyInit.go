package config

import (
	"log"
	"os"
	"os/signal"
)

var ServerSigChan chan os.Signal

func init() {
	ServerSigChan = make(chan os.Signal)
}
func ShutdownServer(err error) {
	log.Println(err)
	ServerSigChan <- os.Interrupt
}
func ServerNotify() {
	signal.Notify(ServerSigChan, os.Interrupt)
	<-ServerSigChan
}
