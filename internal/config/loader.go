package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

var instance *EventProcessingConfig

// GetEventProcessingConfig will retrieve the global event processing config.
func GetEventProcessingConfig() (*EventProcessingConfig, error) {
	if instance == nil {
		return nil, fmt.Errorf("the event processing config has not been sucessfully loaded")
	}

	return instance, nil
}

// LoadConfig will load the event processing config.
func LoadConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("error while reading config file: %v", err)
	}

	var eventProcessingConfig EventProcessingConfig

	err = yaml.Unmarshal(data, &eventProcessingConfig)

	if err != nil {
		return fmt.Errorf("error while parsing config file: %v", err)
	}

	if err := verifyEvents(eventProcessingConfig.Events); err != nil {
		return fmt.Errorf("config failed verification: %v", err)
	}

	instance = &eventProcessingConfig

	return nil
}

func verifyEvents(events []Event) error {
	// Make sure there are no zero length commands
	for _, event := range events {
		for _, command := range event.Commands {
			if len(strings.TrimSpace(command)) == 0 {
				return fmt.Errorf("event %s has a zero length command", event.ID)
			}
		}
	}

	return nil
}
