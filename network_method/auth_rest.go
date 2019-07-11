package network_method

import (
	"../authpb"
	"../config"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func StartRESTServer(address config.ServerNodes, grpcAddress config.ServerNodes, certFile string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	//creds, err := credentials.NewClientTLSFromFile(certFile, "")
	//if err != nil {
	//	return fmt.Errorf("could not load TLS certificate: %s", err)
	//}

	// Setup the client gRPC options
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcAddress.Host+":"+grpcAddress.Port, opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}

	log.Printf("starting HTTP/1.1 REST server on %s", fmt.Sprintf("%s:%s", address.Host, address.Port))
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", address.Host, address.Port), mux)
	if err != nil {
		log.Fatalf("REST server not running : %v", err)
	}
	return nil
}