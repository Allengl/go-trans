package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/zserge/lorca"
)


func main()  {
	var ui lorca.UI 
	ui, _ = lorca.New("https://baidu.com", "", 800, 600, "--disable-sync","--disable-translate") // 创建一个窗口
	chSignal := make(chan os.Signal, 1)   // 1.创建一个信号通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM) //监听系统信号
	select {     // 阻塞主线程 
	case <-ui.Done():	 // 2.监听窗口关闭事件
	case <-chSignal:   // 3.监听系统信号
}
ui.Close() // 4.关闭窗口
}
