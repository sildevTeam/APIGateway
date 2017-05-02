package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"./testservice"
	"flag"
	"fmt"
)

var host string
var port int

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "APIGateway host listen address,127.0.0.1 for default")
	flag.IntVar(&port, "port", 9090, "APIGateway's serve port,9090 for default")
	flag.Parse()
}
func main() {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("error on creating server socket : %s", err.Error())
		return
	}
	handler := testservice.TestService{}
	processor := testservice.NewSilGatewayTestServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	if err = server.Serve(); err != nil {
		panic(err)
	}

}
