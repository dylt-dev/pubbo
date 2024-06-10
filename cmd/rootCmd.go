package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CreateRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pubbo",
		Short: "pubbo",
		Long: "Publish a file & make it availble on a Unix socket",
	}
	return cmd
}
