package main

import (
	"fmt"
	"log"

	"github.com/zhendliu/wind"
	"github.com/zhendliu/wind/server"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ser := server.NewServer()
	err := ser.RegisterName(&HelloWorld{}, "helloworld")
	if err != nil {
		log.Fatalln(err)
	}

	if err := ser.Run(server.UseTCP("0.0.0.0:8074")); err != nil {
		log.Fatalln(err)
	}
}

type HelloWorld struct{}

type HelloWorldRequest struct {
	Name string
}

type HelloWorldResponse struct {
	RPName string
}

func (s *HelloWorld) HelloWorld(ctx *light.Context, req *HelloWorldRequest, resp *HelloWorldResponse) error {
	resp.RPName = fmt.Sprintf("hello world by: %s", req.Name)
	//return errors.New(":xx")
	//fmt.Println(resp)
	return nil
}
