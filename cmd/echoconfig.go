package cmd

import (
	"github.com/meowfaceman/script-on-docker-events/internal/config"
	"github.com/spf13/cobra"
)

var echoConfig = &cobra.Command{
	Use:   "echo-config",
	Short: "Echoes the config for Docker events and their corresponding actions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.EchoConfig()
	},
}
