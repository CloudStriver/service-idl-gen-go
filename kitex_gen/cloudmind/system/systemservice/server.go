// Code generated by Kitex v0.8.0. DO NOT EDIT.
package systemservice

import (
	system "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/system"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler system.SystemService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}