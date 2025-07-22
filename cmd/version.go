package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AppVersion string

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Git AI CLI",
		Long:  `All software has versions. This is Git AI CLI's.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("GitAI CLI Version: %s\n", AppVersion)
		},
	}
} 