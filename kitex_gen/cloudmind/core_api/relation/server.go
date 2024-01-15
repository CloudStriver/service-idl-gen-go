// Code generated by Kitex v0.8.0. DO NOT EDIT.
package relation

import (
	core_api "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/core_api"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler core_api.Relation, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
