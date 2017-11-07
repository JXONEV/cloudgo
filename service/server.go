package service

import (
    // 使用martini框架
    "github.com/go-martini/martini"
)

func NewServer(port string) {
    m := martini.Classic()
    // 添加参数
    m.Get("/hello/:id", func(params martini.Params) string {
        return "Hello " + params["id"]
    })
    // 对应端口
    m.RunOnAddr(":"+port)
}
