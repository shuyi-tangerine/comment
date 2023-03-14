package main

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/comment/gen-go/base"
	"github.com/shuyi-tangerine/comment/gen-go/tangerine/comment"

	"github.com/apache/thrift/lib/go/thrift"
)

var defaultCtx = context.Background()

func handleClient(client comment.CommentHandler) (err error) {
	res, err := client.GenCommentID(defaultCtx, &comment.GenCommentIDRequest{Base: &base.RPCRequest{}})
	if err != nil {
		return err
	}
	fmt.Println("res ==>", res.Base.Code, res.Base.Message)
	fmt.Println("CommentIDs ==>", res.CommentIds)
	return
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return handleClient(comment.NewCommentHandlerClient(thrift.NewTStandardClient(iprot, oprot)))
}
