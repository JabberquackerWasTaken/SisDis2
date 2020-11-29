package main

import (
	"log"
	"net"

	"github.com/JabberquackerWasTaken/SisDis/chat"
	"google.golang.org/grpc"
)

/*Funcion main
La logica.go solo es el servidor levandado. las funciones con relacion a la logica estan en chat.go :D
*/
func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
	}

	s3 := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s3)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8000: %v", err)
	}

}
