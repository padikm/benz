package main

import (
	"svc1/client"
	"svc1/server"
)

func init() {
	c := client.GrpcClientImpl{}
	client.IntiClient(c)
}

func main() {

	//go client.GrpcClient();
	server.StartHttpServer()
}
