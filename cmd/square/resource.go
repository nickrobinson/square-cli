package main

import (
	"fmt"
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/spf13/cobra"
)

func buildResourceCommand(sq *square.Square, resourceName string) *cobra.Command {
	return &cobra.Command{
		Use: resourceName,
		Annotations: map[string]string{
			"group": "resources",
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(sq.AccessKey)
		},
	}
}
