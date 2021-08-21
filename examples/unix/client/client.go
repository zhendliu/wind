package main

import (
	"fmt"
	"log"

	"github.com/zhendliu/wind"
	"github.com/zhendliu/wind/client"
	"github.com/zhendliu/wind/discovery"
	"github.com/zhendliu/wind/transport"
)

type MethodTestReq struct {
	Name string
}

type MethodTestResp struct {
	RPName string
}

func main() {
	client := client.NewClient(discovery.NewSimplePeerToPeer("./unix.sock", transport.UNIX))
	connect, err := client.NewConnect("helloworld")
	if err != nil {
		log.Fatalln(err)
		return
	}

	req := MethodTestReq{
		Name: "hello",
	}
	resp := MethodTestResp{}
	err = connect.Call(light.DefaultCtx(), "HelloWorld", &req, &resp)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}
