package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		panic(err)
	}
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		readVal, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		udpConn.Write([]byte(readVal))

	}

}
