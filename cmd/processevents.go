package cmd

import (
	"time"

	"github.com/meowfaceman/script-on-docker-events/internal/eventprocessor"
	"github.com/spf13/cobra"
)

var startMinutesAgo int

var processEvents = &cobra.Command{
	Use:   "process-events",
	Short: "Processes Docker events and runs configured actions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		timestamp := time.Now().Add(time.Duration(-startMinutesAgo) * time.Minute).Unix()
		return eventprocessor.ProcessEvents(timestamp)
	},
}

func init() {
	processEvents.Flags().IntVar(&startMinutesAgo, "start-minutes-ago", 5, "look at events for the previous start-minutes-ago minutes on startup (default: 5)")
}
