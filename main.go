package main

import (
    "os"
    "cloudgo/service"
    flag "github.com/spf13/pflag"
)

const (
    PORT string = "8080" // 设置默认端口
)

func main() {
    // 判断是否监听到端口，若没有则使用默认端口
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }
    // 设置端口
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    // 启动预设服务
    service.NewServer(port)
}
