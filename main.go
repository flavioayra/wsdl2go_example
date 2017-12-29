package main

import (
	"github.com/fiorix/wsdl2go/soap"
)

func main() {

	cli := soap.Client{
		URL:       "http://server",
		Namespace: example.Namespace,
	}
	soapService := example.NewEchoService(&cli)
	echoReply, err := soapService.Echo(&example.EchoRequest{Data: "hello world"})
}
