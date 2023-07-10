package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	con, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	var b []byte
	_, err = con.Read(b)
	if err != nil {
		fmt.Println("Error reading sent data")
		os.Exit(1)
	}
	_, err = con.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("Sending PONG failed")
		os.Exit(1)
	}
}
