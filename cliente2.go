package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Makkami/SisDis2/chat"
	"google.golang.org/grpc"
)

func con() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("a %v", err)
	}
}

func separarArchivo() {
	rand.Seed(time.Now().Unix())

	ports := []string{
		":9090",
		":8080",
		":8000",
	}
	n := rand.Int() % len(ports)
	port := ports[n]
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	c := chat.NewChatServiceClient(conn)
	defer conn.Close()

	var request = bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----------------------------")
		fmt.Printf("Ingrese una opcion de Orden: ")
		libro, _ := request.ReadString('\n')
		libro = strings.Trim(libro, " \r\n")
		fileToBeChunked := "./" + libro + ".pdf"

		file, err := os.Open(fileToBeChunked)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer file.Close()

		fileInfo, _ := file.Stat()

		var fileSize int64 = fileInfo.Size()

		const fileChunk = 250 * (1 << 10) // Este (1 << 10) es igual a 2^10, entonces es 250 * 1024 = 256000

		// calculate total number of parts the file will be chunked into

		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		fmt.Printf("Dividiendo el archivo en %d partes.\n", totalPartsNum)

		for i := uint64(0); i < totalPartsNum; i++ {

			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)
			file.Read(partBuffer)

			message := chat.Chunk{
				Nombre:        libro,
				Parte:         strconv.FormatUint(i, 10),
				NumPartes:     totalPartsNum,
				Buffer:        partBuffer,
				Port:          port,
				PartesTotales: strconv.FormatUint(i, 10),
			}

			var response *chat.Message

			response, _ = c.SubirChunk(context.Background(), &message)
			fmt.Printf("Hola que tal, %s \n", response.Body)

		}
	}
}

func main() {

	var request = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("----------------------------")
		fmt.Println("Ingrese una opcion de Orden:")
		fmt.Println("1.-Cargar Libro")
		fmt.Println("2.-Descargar Libro")
		fmt.Println("----------------------------")
		fmt.Print("Opcion: ")
		text, _ := request.ReadString('\n')
		text = strings.ToLower(strings.Trim(text, " \r\n"))
		if strings.Compare(text, "1") == 0 {
			separarArchivo()
		}
		if strings.Compare(text, "2") == 0 {
			separarArchivo()
		}
	}
}
