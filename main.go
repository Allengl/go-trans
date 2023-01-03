//go:generate go-winres make --product-version=git-tag
package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/Allengl/go-trans/config"
	"github.com/Allengl/go-trans/server"
)

func main() {
	chChromeDie := make(chan struct{})
	chBackendDie := make(chan struct{})
	chSignal := listenToInterrupt()
	go server.Run()
	go startBrowser(chChromeDie, chBackendDie)
	for {
		select {
		case <-chSignal: // 阻塞等待信号

		case <-chChromeDie:
			os.Exit(0)
		}
	}
}

func startBrowser(chChromeDie chan struct{}, chBackendDie chan struct{}) {
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+config.GetPort()+"/static/index.html")
	cmd.Start() // 开一个新的进程
	go func() {
		<-chBackendDie
		cmd.Process.Kill()
	}()
	go func() {
		cmd.Wait()
		chChromeDie <- struct{}{}
	}()
}

func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
