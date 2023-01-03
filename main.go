package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/Allengl/go-trans/server"
)


func main() {

	// 启动 gin 服务
	go server.Run()

	cmd := startBrowser()

	chSignal := listenToInterrupt()
	select {
	case <-chSignal: // 阻塞等待信号
	cmd.Process.Kill()
	}
}

func startBrowser() *exec.Cmd {
	port := "27149"
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+port+"/static/index.html")
	cmd.Start()  // 开一个新的进程	
	return cmd
}

// 监听中断信号
func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
