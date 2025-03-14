package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

// trunked: [chunk size][chunk data][chunk boundary]
func Client() {
	resp, err := http.Get("http://127.0.0.1:8081/name")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.TransferEncoding)

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			fmt.Println(line)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	Client()
}
