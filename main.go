package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/Allengl/go-trans/server"
)


func main() {
	port := "27149"

	// 启动 gin 服务
	go func() {
		server.Run()
	}()  // 开一个协程

	// 启动 Chrome
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+port+"/static/index.html")
	cmd.Start()  // 开一个新的进程
	
	// 监听中断信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// 等待中断信号
	select {
	case <-chSignal: // 阻塞等待信号
	cmd.Process.Kill()
	}
}
