package handlers

import (
  "os"
  "io/ioutil"

	"google.golang.org/grpc/grpclog"
	proto "github.com/fikrirnurhidayat/bookstoresvc/gen"
)

type Backend struct {
  proto.BookstoreServiceServer
  Logger grpclog.LoggerV2
}

func NewServer() *Backend {
  // Setup Logger
  log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

  // Create server
	server := &Backend{
		Logger: log,
	}

  return server;
} 
