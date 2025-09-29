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

	// Remove the socket. os.RemoveAll() returns nil if no socket
	err := os.RemoveAll(socketPath)
	if err != nil {
		fmt.Println("error removing socket")
		os.Exit(1)
	}

	// Listen for connections on socket
	sock, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("error opening socket")
		os.Exit(1)
	}

	// Add a `defer` that closes the socket
	// @note maybe this should remove the socket
	defer func(sock net.Listener) {
		fmt.Println("closing socket ...")
		sock.Close()
	}(sock)
	
	// chmod the rocket to read/write for owners, readonly for world
	err = os.Chmod(socketPath, 0664)
	if err != nil {
		fmt.Println("error chmod'ing socket")
		os.Exit(1)
	}
	
	// Infinite loop: accept incoming connections + write the path's contents
	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Error getting connection")
			os.Exit(1)
		}
		fmt.Printf("Connection received. Sending content at %s ...\n", filePath)
		nWritten, err := pub(filePath, conn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		} else {
			fmt.Printf("%d byte(s) written.\n", nWritten)
		}
	}
}


func pub (filePath string, conn net.Conn) (int64, error) {
	var f *os.File
	
	// Add a `defer` to close the connection
	defer func (conn net.Conn) {
		var err error
		if f != nil {
			err = f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error closing file (%s)", filePath)
			}
		}

		fmt.Println("Closing connection ...")
		err = conn.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error closing connection")
		}
	}(conn)

	// Try and open the file
	// @note this aborts the file on temporary file unavailability
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file (%s)\n", filePath)
		return 0, err
	}

	// @note this only writes once. if there's more to write, too bad.
	var nWritten int64
	nWritten, err = io.Copy(conn, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying data: %n byte(s) written\n", nWritten)		
	}

	return nWritten, nil
}

