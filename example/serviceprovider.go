// +build ignore

package main

import (
	"fmt"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/msg"
)

type TestServiceReq struct {
	A msg.Float64
	B msg.String
}

type TestServiceRes struct {
	C msg.Float64
}

func onService(req *TestServiceReq) *TestServiceRes {
	fmt.Println("request:", req)
	return &TestServiceRes{
		C: 123,
	}
}

func main() {
	// create a node with given name and linked to given master.
	// master can be reached with an ip or hostname.
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:       "/goroslib",
		MasterHost: "127.0.0.1",
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a service provider
	sp, err := goroslib.NewServiceProvider(goroslib.ServiceProviderConf{
		Node:     n,
		Service:  "/test_srv",
		Callback: onService,
	})
	if err != nil {
		panic(err)
	}
	defer sp.Close()

	// freeze main loop
	infty := make(chan int)
	<-infty
}
