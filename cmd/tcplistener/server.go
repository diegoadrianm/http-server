package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	port := ":42069"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	fmt.Println("Listening on port 42069...")

	for {

		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())

		ch := getLinesChannel(conn)
		for val := range ch {
			fmt.Printf("read: %s\n", val)
		}
		fmt.Println("Connection to ", conn.LocalAddr(), "closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {

	curr := ""
	var parts []string
	buff := make([]byte, 8)
	channel := make(chan string)

	go func() {
		defer close(channel)
		for {
			read, err := f.Read(buff)
			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			parts = strings.Split(string(buff[:read]), "\n")
			if len(parts) == 2 {
				channel <- (curr + parts[0])
				parts = parts[1:]
				curr = parts[0]
			} else {
				curr = curr + string(buff[:read])
			}
		}
	}()
	return channel
}
