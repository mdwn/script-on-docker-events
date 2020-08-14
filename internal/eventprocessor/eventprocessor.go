package eventprocessor

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
	"github.com/meowfaceman/script-on-docker-events/internal/config"
)

const (
	bash = "bash"

	// Special initialization type/action
	initEvent = "init"
)

// ProcessEvents will recognize events and execute scripts that match the given configuration.
func ProcessEvents(since int64) error {
	config, err := config.GetEventProcessingConfig()

	if err != nil {
		return fmt.Errorf("error getting event processing config: %v", err)
	}

	client, err := docker.NewEnvClient()

	if err != nil {
		return fmt.Errorf("error creating docker client: %v", err)
	}

	processInitEvent(config.Events)

	events, errors := client.Events(context.Background(), types.EventsOptions{
		Since: fmt.Sprintf("%d", since),
	})

	for {
		select {
		case event, open := <-events:
			if !open {
				return nil
			}

			for _, eventAndScripts := range config.Events {
				if event.Type == eventAndScripts.ObjectType && event.Action == eventAndScripts.Action {
					attributesMatch := true
					for key, value := range eventAndScripts.Attributes {
						if event.Actor.Attributes[key] != value {
							attributesMatch = false
						}
					}

					if attributesMatch {
						fmt.Printf("Event match: %s\n", eventAndScripts.ID)
						for _, command := range eventAndScripts.Commands {
							go runCommand(command)
						}
					}
				}
			}
		case error, open := <-errors:
			if !open {
				return nil
			}

			fmt.Printf("Error: %+v\n", error)
		}
	}
}

func processInitEvent(events []config.Event) {
	for _, event := range events {
		if event.ObjectType == initEvent && event.Action == initEvent {
			fmt.Println("Running initialization...")
			for _, command := range event.Commands {
				// Run these synchronously to make sure they're all set up before we start trying to process other events.
				runCommand(command)
			}
			break
		}
	}
}

func runCommand(command string) {
	execCmd := exec.Command("bash", "-c", command)

	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	fmt.Printf("Executing %s...\n", command)
	err := execCmd.Run()

	// If the script fails to run, we'll keep on rolling, but note it here.
	if err != nil {
		fmt.Printf("Error executing %s: %v\n", command, err)
	} else {
		fmt.Printf("Command %s executed successfully.\n", command)
	}
}
