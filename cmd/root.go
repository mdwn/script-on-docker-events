package cmd

import (
	"fmt"
	"os"

	"github.com/meowfaceman/script-on-docker-events/internal/config"
	"github.com/spf13/cobra"
)

var (
	// Event processor config file
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "script-on-docker-events",
		Short: "Runs user configurable scripts when encountering user configurable Docker events.",
		Long: `script-on-docker-events will run a given set of scripts when encountering
Docker events.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "the event processing config")

	rootCmd.AddCommand(echoConfig)
	rootCmd.AddCommand(processEvents)
}

func initConfig() {
	if cfgFile != "" {
		err := config.LoadConfig(cfgFile)

		if err != nil {
			fmt.Printf("error loading event processing config: %v", err)
			os.Exit(1)
		}
	}
}

// Execute will execute the event processor and its subcommands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
