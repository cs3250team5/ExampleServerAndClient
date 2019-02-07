package main

import (
	"fmt"
	"net"
	"os"
)

const (
	Port        = "8080"
	Host        = "localhost"
	MessageSize = 1024
)

func main() {
	listener, err := net.Listen("tcp", Host+":"+Port)
	if err != nil {
		fmt.Println("Error in Main: net.Listen: " + err.Error())
		os.Exit(1)
	}
	i := 0
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in Main: listener.connect: " + err.Error())
		}
		go handleConnection(connection, i)
		i++
	}

}

func handleConnection(connection net.Conn, num int) {
	fmt.Println("Connected to ", num)
	s := "Message Recieved"
	defer func() {
		connection.Close()
		fmt.Println("Connection ", num, " closed")
	}()
	buffer := make([]byte, MessageSize)
	numBytes, err := connection.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error in server:HandleConnection: " + err.Error())
		return
	}

	fmt.Fprintf(connection, s)

	fmt.Print(string(buffer[:numBytes]) + "\n")
}
