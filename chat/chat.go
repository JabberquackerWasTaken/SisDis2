package chat

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

//Server xd
type Server struct {
	dn1 []int
	dn2 []int
	dn3 []int
}

func conseguirChunk(Numero int, Name string) Chunk {
	var Envio Chunk
	return Envio
}

//transformarArregloS transforma una lista de ints a un string
func transformarArregloS(lista []int) string {
	Respuesta := ""
	for i := 0; i < len(lista); i++ {
		Respuesta = Respuesta + "@@" + strconv.Itoa(i)
	}
	return Respuesta
}

//SubirChunk xd
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

	if message.Parte == message.Partestotales {
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		c := NewChatServiceClient(conn)
		defer conn.Close()

		var response *PropuestaRespuesta
		response, _ = c.PedirPropuesta(context.Background(), message)
		fmt.Printf("Hola que tal \n")
		fmt.Println(response.Nombre)
		largo1, err := strconv.Atoi(response.Intn1)
		largo2, err := strconv.Atoi(response.Intn2)
		largo3, err := strconv.Atoi(response.Intn3)
		listan1 := strings.SplitN(response.Nd1, "@", largo1)
		listan2 := strings.SplitN(response.Nd1, "@", largo2)
		listan3 := strings.SplitN(response.Nd1, "@", largo3)

		for k := 0; k < largo1; k++ {
			if message.Port == ":9090" {
				break
			}
			NumeroChunk, err := strconv.Atoi(listan1[k])
			if err != nil {
				log.Fatalf("Error al extraer int: %s", err)
			}
			var Mail Chunk
			Mail = conseguirChunk(NumeroChunk, response.Nombre)
			//conectarme a port
			conn2, err := grpc.Dial(":9090", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			c2 := NewChatServiceClient(conn)
			defer conn2.Close()
			response, _ := c2.SubirChunk2(context.Background(), &Mail)
			fmt.Println(response)
		}
		for k := 0; k < largo2; k++ {
			if message.Port == ":8000" {
				break
			}
			NumeroChunk, err := strconv.Atoi(listan2[k])
			if err != nil {
				log.Fatalf("Error al extraer int: %s", err)
			}
			var Mail2 Chunk
			Mail2 = conseguirChunk(NumeroChunk, response.Nombre)
			//conectarme a port

			conn2, err := grpc.Dial(":8000", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			c2 := NewChatServiceClient(conn)
			defer conn2.Close()
			response, _ := c2.SubirChunk2(context.Background(), &Mail2)
			fmt.Println(response)

		}
		for k := 0; k < largo3; k++ {
			if message.Port == ":8080" {
				break
			}
			NumeroChunk, err := strconv.Atoi(listan3[k])
			if err != nil {
				log.Fatalf("Error al extraer int: %s", err)
			}
			var Mail3 Chunk
			Mail3 = conseguirChunk(NumeroChunk, response.Nombre)
			//conectarme a port
			conn2, err := grpc.Dial(":8080", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			c2 := NewChatServiceClient(conn)
			defer conn2.Close()
			response, _ := c2.SubirChunk2(context.Background(), &Mail3)
			fmt.Println(response)

		}
	}

	fmt.Println("Dividido en: ", fileName)
	return &Message{Body: ""}, nil
}

//SayHola envia mensajes entre servidor-cliente y siempre es un chat.Message, Lo que el string dire depende de lo que el cliente le pide
func (s *Server) SayHola(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	respuesta := Message{
		Body: "iniciando",
	}
	return &respuesta, nil
}

//PedirPropuesta manda el struct con el numero total de partes al NameNode el cual revisa los nodos abilitados y responde con la propuesta.
func (s *Server) PedirPropuesta(ctx context.Context, message *Chunk) (*PropuestaRespuesta, error) {
	//aca falta agregar el richard
	flag1 := false
	flag2 := false
	flag3 := false
	i := 0
	var dn1 []int
	var dn2 []int
	var dn3 []int
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Println("Nodo 1 caido")
	} else {
		flag1 = true
	}
	conn.Close()

	conn, err = grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Println("Nodo 2 caido")
	} else {
		flag2 = true
	}

	conn.Close()
	conn, err = grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Println("Nodo 3 caido")
	} else {
		flag3 = true
	}
	largo, err := strconv.Atoi(message.Partestotales)
	conn.Close()
	for i = 0; i < largo; i += 0 {
		if flag1 == false && flag2 == false && flag3 == false {
			log.Println("Todos los nodos caidos")
			break
		}
		if flag1 == true {
			if i == largo {
				break
			}

			dn1 = append(dn1, i)
			i++
		}
		if flag2 == true {
			if i == largo {
				break
			}
			dn2 = append(dn2, i)
			i++
		}
		if flag3 == true {
			if i == largo {
				break
			}
			dn3 = append(dn3, i)
			i++
		}
	}
	st1 := transformarArregloS(dn1)
	st2 := transformarArregloS(dn2)
	st3 := transformarArregloS(dn3)
	respuesta := PropuestaRespuesta{
		Nombre: message.Nombre,
		Total:  message.Partestotales,
		Nd1:    st1,
		Nd2:    st2,
		Nd3:    st3,
		Intn1:  strconv.Itoa(len(st1)),
		Intn2:  strconv.Itoa(len(st2)),
		Intn3:  strconv.Itoa(len(st3)),
	}
	//aca falta que la acepten y escribirla.
	return &respuesta, nil
}

//SubirChunk2 xd
func (s *Server) SubirChunk2(ctx context.Context, message *Chunk) (*Message, error) {
	// write to disk
	fileName := message.Nombre + "_" + message.Parte
	_, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// write/save buffer to disk
	ioutil.WriteFile(fileName, message.Buffer, os.ModeAppend)

	fmt.Println("Chunk: ", fileName, " recibido")
	return &Message{Body: ""}, nil
}
