package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/gin-gonic/gin"
)


func main() {
	go func() {
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.GET("/", func(c *gin.Context) {
			c.Writer.Write([]byte("Hello World")) //
		})
		router.Run(":8080")
	}()  // 开一个协程

	chromePath := "D:\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/")
	cmd.Start()  // 开一个新的进程
	
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	select {
	case <-chSignal: // 阻塞等待信号
	cmd.Process.Kill()
	}
}
