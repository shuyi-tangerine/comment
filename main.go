package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	server := flag.Bool("server", false, "Run server")
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	transport := flag.String("transport", "", "Use framed/buffered transport（framed/buffered）")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")

	flag.Parse()

	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(nil)
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(nil)
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(nil)
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	if *transport == "buffered" {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else if *transport == "framed" {
		transportFactory = thrift.NewTFramedTransportFactoryConf(transportFactory, cfg)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *server {
		if err := runServer(transportFactory, protocolFactory, *addr, *secure); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := runClient(transportFactory, protocolFactory, *addr, *secure, cfg); err != nil {
			fmt.Println("error running client:", err)
		}
	}
}
