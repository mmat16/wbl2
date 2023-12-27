package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	server, _ := net.Listen("tcp", "127.0.0.1:8080")
	defer server.Close()
	fmt.Println(server.Addr())

	for {
		conn, _ := server.Accept()
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			read := "Received msg: " + scanner.Text()
			conn.Write([]byte(read))
		}
		err := scanner.Err()
		if err != nil {
			log.Println(err)
			conn.Close()
			server.Close()
			os.Exit(1)
		}
		conn.Close()
	}
}
