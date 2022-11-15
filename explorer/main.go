package main

// import (
// 	"fmt"
// 	"net"
// 	"os"

// 	explorer "github.com/TudorEsan/FinanceAppGo/explorer/proto"
// 	"github.com/TudorEsan/FinanceAppGo/explorer/service"
// 	"github.com/hashicorp/go-hclog"
// 	"github.com/joho/godotenv"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )


// func main() {
// 	godotenv.Load()
// 	l := hclog.Default()
// 	port, ok := os.LookupEnv("PORT")
// 	if !ok {
// 		fmt.Println("PORT: ", port)
// 	}


// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	}

// 	s := grpc.NewServer()
	
// 	// explorer service
// 	explorerS := service.NewAddressExplorerServer(l)
// 	explorer.RegisterAddressExplorerServer(s, explorerS)
// 	// use reflection to register all services
// 	reflection.Register(s)
	
// 	s.Serve(lis)
// }