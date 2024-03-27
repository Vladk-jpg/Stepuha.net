package main

import (
	"fmt"
	"net"
)

func main() {
	message := "mornin'"
	fmt.Println("jpg")
	fmt.Println("jpg")
	fmt.Println("TestText")
	server, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println("Nice job! It doesn't work!")
		return
	}
	defer server.Close()

	for {
		server, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		server.Write([]byte(message))
		server.Close()
	}
}
