package main

import (
	"fmt"
	"os"

	"pubbo/cmd"
)

func main () {
/*
	filePath := s.Args[1]
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
*/
	os.Exit(Run())
}

func Run () int {
	rootCmd := cmd.CreateRootCommand()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

