package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/JabberquackerWasTaken/SisDis/chat"
	"google.golang.org/grpc"
)

/*Funcion main
La logica.go solo es el servidor levandado. las funciones con relacion a la logica estan en chat.go :D
*/
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	sN := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &sN)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
	var request = bufio.NewReader(os.Stdin)
	fmt.Println("----------------------------")
	fmt.Println("Antes de comenzar porfavor inicie todos los nodos")
	fmt.Println("Cuando todos los nodos esten corriendo presionar enter.")
	text, _ := request.ReadString('\n')
	text = strings.ToLower(strings.Trim(text, " \r\n"))
	///
	//Conectar nodo1
	conec1, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conec1.Close()
	c1 := chat.NewChatServiceClient(conec1)

	//Conectar Nodo2
	conec2, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conec2.Close()
	c2 := chat.NewChatServiceClient(conec2)

	//Conectar Nodo3
	conec3, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conec3.Close()
	c3 := chat.NewChatServiceClient(conec3)

	///
	message := chat.Message{
		Body: "Largo",
	}
	response, err := c1.SayHola(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling server: %s", err)
	}
	fmt.Println(response.Body)
	////////////////////////////////////////////////////////////
	response, err = c2.SayHola(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling server: %s", err)
	}
	fmt.Println(response.Body)
	///////////////////////////////////////////////////
	response, err = c3.SayHola(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling server: %s", err)
	}
	fmt.Println(response.Body)
}
