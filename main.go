package main

import (
	"context"
  "os"
  "io/ioutil"
  "net"
  "net/http"
  "fmt"
  "flag"

	"google.golang.org/grpc/grpclog"
  "google.golang.org/grpc"
  "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
  proto "github.com/fikrirnurhidayat/bookstoresvc/gen"
  handlers "github.com/fikrirnurhidayat/bookstoresvc/app/handlers"
)

// Grab configuration from cli
var (
	gRPCPort    = flag.Int("grpc-port", 10000, "The gRPC server port")
	gatewayPort = flag.Int("gateway-port", 11000, "The gRPC-Gateway server port")
)

var log grpclog.LoggerV2

func serveHTTP(gatewayAddr string, grpcAddr string, mux *runtime.ServeMux) error {
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
    grpc.WithInsecure(),
	)

  if err != nil {
    log.Fatalf("[gateway] Failed to dial to %s: %v", grpcAddr, err)
  }

  err = proto.RegisterBookstoreServiceHandler(context.Background(), mux, conn)

  if err != nil {
    log.Fatal("[gateway] Cannot register BookstoreServiceHandler: %v", err)
  }

  log.Infof("[gateway] Server is running at http://%s", gatewayAddr)
  http.ListenAndServe(gatewayAddr, mux)
  return nil
}

func serveRPC(s grpc.Server, grpcAddr string) error {
  tcpListener, err := net.Listen("tcp", grpcAddr)

  if err != nil {
    log.Fatalf("[tcp] Failed to listen at tcp://%s", grpcAddr)
  }

  log.Info("[grpc] Server is running at http://")
  s.Serve(tcpListener)
  return nil
}

func main() {
	log = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

  addr := fmt.Sprintf("localhost:%d", *gRPCPort)
	gatewayAddr := fmt.Sprintf("localhost:%d", *gatewayPort)

  mux := runtime.NewServeMux()
  s := grpc.NewServer();

  proto.RegisterBookstoreServiceServer(s, handlers.NewServer())

  go serveRPC(*s, addr)
  serveHTTP(gatewayAddr, addr, mux)
}
