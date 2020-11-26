package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Makkami/SisDis2/chat"
	"google.golang.org/grpc"
)

func separarArchivo() {
	var request = bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----------------------------")
		fmt.Printf("Ingrese una opcion de Orden:")
		libro, _ := request.ReadString('\n')
		libro = strings.ToLower(strings.Trim(libro, " \r\n"))
		fileToBeChunked := "./" + libro + ".pdf"

		file, err := os.Open(fileToBeChunked)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer file.Close()

		fileInfo, _ := file.Stat()

		var fileSize int64 = fileInfo.Size()

		const fileChunk = 250 * (1 << 10)
		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

		for i := uint64(0); i < totalPartsNum; i++ {

			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			message := chat.Chunk{
				Nombre:    libro,
				Parte:     strconv.FormatUint(i, 10),
				NumPartes: totalPartsNum,
				Buffer:    partBuffer,
			}

			var response *chat.Message

			response, _ = c1.SubirChunk(context.Background(), &message)
			log.Printf("Holi 1 %s", response.Body)
		}
	}
}

func main() {

	var con1 *grpc.ClientConn
	con1, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer con1.Close()
	c1 := chat.NewChatServiceClient(con1)
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
	/*		// just for fun, let's recombine back the chunked files in a new file

			newFileName := "NEWbigfile.zip"
			_, err = os.Create(newFileName)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			//set the newFileName file to APPEND MODE!!
			// open files r and w

			file, err = os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			// IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
			// defer file.Close()

			// just information on which part of the new file we are appending
			var writePosition int64 = 0

			for j := uint64(0); j < totalPartsNum; j++ {

					//read a chunk
					currentChunkFileName := "bigfile_" + strconv.FormatUint(j, 10)

					newFileChunk, err := os.Open(currentChunkFileName)

					if err != nil {
							fmt.Println(err)
							os.Exit(1)
					}

					defer newFileChunk.Close()

					chunkInfo, err := newFileChunk.Stat()

					if err != nil {
							fmt.Println(err)
							os.Exit(1)
					}

					// calculate the bytes size of each chunk
					// we are not going to rely on previous data and constant

					var chunkSize int64 = chunkInfo.Size()
					chunkBufferBytes := make([]byte, chunkSize)

					fmt.Println("Appending at position : [", writePosition, "] bytes")
					writePosition = writePosition + chunkSize

					// read into chunkBufferBytes
					reader := bufio.NewReader(newFileChunk)
					_, err = reader.Read(chunkBufferBytes)

					if err != nil {
							fmt.Println(err)
							os.Exit(1)
					}

					// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
					// write/save buffer to disk
					//ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

					n, err := file.Write(chunkBufferBytes)

					if err != nil {
							fmt.Println(err)
							os.Exit(1)
					}

					file.Sync() //flush to disk

					// free up the buffer for next cycle
					// should not be a problem if the chunk size is small, but
					// can be resource hogging if the chunk size is huge.
					// also a good practice to clean up your own plate after eating

					chunkBufferBytes = nil // reset or empty our buffer

					fmt.Println("Written ", n, " bytes")

					fmt.Println("Recombining part [", j, "] into : ", newFileName)
			}

			// now, we close the newFileName
			file.Close()*/
}
