// +build ignore

package main

import (
	"fmt"

	"github.com/aler9/goroslib"
	"github.com/aler9/goroslib/pkg/msg"
)

// define a custom service.
// unlike the standard library, a .srv file is not needed.

type TestServiceReq struct {
	msg.Package `ros:"my_package"`
	A           float64
	B           string
}

type TestServiceRes struct {
	msg.Package `ros:"my_package"`
	C           float64
}

type TestService struct {
	TestServiceReq
	TestServiceRes
}

func onService(req *TestServiceReq) *TestServiceRes {
	fmt.Println("request:", req)
	return &TestServiceRes{
		C: 123,
	}
}

func main() {
	// create a node and connect to the master
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "goroslib_sp",
		MasterAddress: "127.0.0.1:11311",
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a service provider
	sp, err := goroslib.NewServiceProvider(goroslib.ServiceProviderConf{
		Node:     n,
		Name:     "test_srv",
		Srv:      &TestService{},
		Callback: onService,
	})
	if err != nil {
		panic(err)
	}
	defer sp.Close()

	// freeze main loop
	select {}
}
