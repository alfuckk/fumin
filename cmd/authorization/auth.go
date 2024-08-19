package main

import (
	"net"
	"os"

	"github.com/alfuckk/fumin/pkg/authorization/endpoints"
	"github.com/alfuckk/pkg/user"
	"github.com/go-kit/log"
)

func main() {
	var (
		logger   log.Logger
		httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
		grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestamp)

	var (
		service   = user.NewService()
		endpoints = endpoints.NewEndpoint(service)
	)
}
