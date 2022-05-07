// Code generated by Kitex v0.3.0. DO NOT EDIT.
package api

import (
	"github.com/OnlineJudge/announcement/kitex_gen/announcement"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler announcement.Api, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
