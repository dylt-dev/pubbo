package main

import (
"fmt"
"io"
"net"
"os"
)

func pub (filePath string, conn net.Conn) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file")
		os.Exit(1)
	}
	defer f.Close()
	defer func (conn net.Conn) {
		fmt.Println("Closing connection ...")
		conn.Close()
	}(conn)
	io.Copy(conn, f)
}


func main () {
	filePath := os.Args[1]
	sockPath := os.Args[2]
	err := os.RemoveAll(sockPath)
	if err != nil {
		fmt.Println("error removing socket")
		os.Exit(1)
	}
	sock, err := net.Listen("unix", sockPath)
	if err != nil {
		fmt.Println("error opening socket")
		os.Exit(1)
	}
	defer func(sock net.Listener) {
		fmt.Println("closing socket ...")
		sock.Close()
	}(sock)
	err = os.Chmod(sockPath, 0664)
	if err != nil {
		fmt.Println("error chmod'ing socket")
		os.Exit(1)
	}
	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Error getting connection")
			os.Exit(1)
		}
		pub(filePath, conn)
	}
}
