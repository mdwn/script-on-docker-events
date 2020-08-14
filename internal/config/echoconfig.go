package config

import "fmt"

// EchoConfig will print the config to standard out.
func EchoConfig() error {
	config, err := GetEventProcessingConfig()

	if err != nil {
		return fmt.Errorf("error getting event processing config: %v", err)
	}

	fmt.Println("Echoing the config...")

	for _, eventAndScripts := range config.Events {
		fmt.Printf("----- %s -----\n", eventAndScripts.ID)
		fmt.Printf("Event Type: %s, Action: %s\n", eventAndScripts.ObjectType, eventAndScripts.Action)
		fmt.Println("Attributes:")

		for key, value := range eventAndScripts.Attributes {
			fmt.Printf("- %s: %s\n", key, value)
		}

		fmt.Println("Commands:")

		for _, command := range eventAndScripts.Commands {
			fmt.Printf("- %s\n", command)
		}

		fmt.Println()
	}

	return nil

}
