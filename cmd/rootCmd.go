package cmd

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/spf13/cobra"
)

func CreateRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pubbo",
		Short: "pubbo",
		Long: "Publish a file & make it availble on a Unix socket",
		RunE: runCmd,
	}
	cmd.Flags().String("file-path", "", "")
	cmd.Flags().String("socket-path", "", "")
	cmd.MarkFlagRequired("file-path")
	cmd.MarkFlagRequired("socket-path")
	return cmd
}

func runCmd (cmd *cobra.Command, args []string) error {
	filePath, _ := cmd.Flags().GetString("file-path")
	socketPath, _ := cmd.Flags().GetString("socket-path")
	fmt.Printf("Pubbo'ing %s at %s ...\n", filePath, socketPath)

	err := os.RemoveAll(socketPath)
	if err != nil {
		fmt.Println("error removing socket")
		os.Exit(1)
	}
	sock, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("error opening socket")
		os.Exit(1)
	}
	defer func(sock net.Listener) {
		fmt.Println("closing socket ...")
		sock.Close()
	}(sock)
	err = os.Chmod(socketPath, 0664)
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
		fmt.Printf("Connection received. Sending content at %s ...\n", filePath)
		pub(filePath, conn)
	}
}


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

