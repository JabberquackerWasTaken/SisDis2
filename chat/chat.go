package chat

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/context"
)

type Server struct {
	dn1 int
	dn2 int
	dn3 int
}

//SubirChunk sube el chunk xd
func (s *Server) SubirChunk(ctx context.Context, message *Chunk) (*Message, error) {
	// write to disk
	fileName := message.Nombre + "_" + message.Parte
	_, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// write/save buffer to disk
	ioutil.WriteFile(fileName, message.Buffer, os.ModeAppend)

	fmt.Println("Dividido en: ", fileName)
	return &Message{Body: "Done"}, nil
}

//SayHola envia mensajes entre servidor-cliente y siempre es un chat.Message, Lo que el string dire depende de lo que el cliente le pide
func (s *Server) SayHola(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	respuesta := Message{
		Body: "iniciando",
	}
	return &respuesta, nil
}
