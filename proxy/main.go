package main

import (
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/auth/grpc/client_adapter"
	clientadaptergeo "github.com/Bubotka/Microservices/proxy/internal/infrastructure/clients/geo/grpc/client_adapter"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/router"
	"github.com/Bubotka/Microservices/proxy/internal/infrastructure/server"
	clientadapteruser "github.com/Bubotka/Microservices/proxy/pkg/clients/user/grpc/client_adapter"
	_ "github.com/lib/pq"
	"log"

	"net/http"
)

func main() {
	geoConnect, err := clientadaptergeo.Connect("geo_nginx_1:8081")
	if err != nil {
		log.Fatal(err)
	}

	authConnect, err := client_adapter.Connect("auth_nginx_1:8082")
	if err != nil {
		log.Fatal(err)
	}

	userConnect, err := clientadapteruser.Connect("user_nginx_1:8083")
	if err != nil {
		log.Fatal(err)
	}

	geoClientGRpcAdapter := clientadaptergeo.NewGeoClientGRpcAdapter(geoConnect)
	userClientGRpcAdapter := clientadapteruser.NewUserClientGRpcAdapter(userConnect)
	authClientGRpcAdapter := client_adapter.NewAuthClientGRpcAdapter(authConnect)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(geoClientGRpcAdapter, userClientGRpcAdapter, authClientGRpcAdapter),
	}

	serv := server.NewServer(s)

	serv.Serve()
}
