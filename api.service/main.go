// Code generated by hertz generator.

package main

import (
	"api.service/biz/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	h := server.Default()
	register(h)
	h.Spin()
}
