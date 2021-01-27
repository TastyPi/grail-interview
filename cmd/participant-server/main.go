package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	pb "github.com/TastyPi/grail-interview/api/participant"
	"github.com/TastyPi/grail-interview/internal/participant/server"
	"github.com/TastyPi/grail-interview/internal/participant/storage"
)

const (
	port = ":50051"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := startServer(); err != nil {
		glog.Fatal(err)
	}
}

func startServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := pb.RegisterParticipantServiceHandlerServer(
		ctx, mux, server.Create(storage.CreateInMemoryParticipantStorage()))
	if err != nil {
		return err
	}

	glog.Infoln("Listening on port", port)
	glog.Flush()
	return http.ListenAndServe(port, mux)
}
