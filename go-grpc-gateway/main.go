package main

import (
	"context"
	"flag"
	"fmt"
	"go-grpc-gateway/src/api/entity"
	"go-grpc-gateway/src/api/transport"
	"go-grpc-gateway/src/config"
	"go-grpc-gateway/src/proto/gRPC/accounts"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func main() {
	if db, err := config.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		entities := []interface{}{&entity.Account{}}
		db.AutoMigrate(entities...)

		address := "0.0.0.0:50051"
		if lis, err := net.Listen("tcp", address); err != nil {
			panic("Error while start gRPC server (account service) at: " + address)
		} else {
			flag.Parse()
			mux := runtime.NewServeMux()
			opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
			if err := accounts.RegisterAccountServiceHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts); err != nil {
				panic("Error while start gRPC Gateway: " + err.Error())
			} else {
				s := grpc.NewServer()
				server := transport.NewGrpcServer(db)
				accounts.RegisterAccountServiceServer(s, server)
				go s.Serve(lis)
				fmt.Println("Server (account service) is running at: " + address)

				http.ListenAndServe(":3001", mux)
			}
		}
	}

}
