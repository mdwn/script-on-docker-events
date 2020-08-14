package config

// EventProcessingConfig is the configuration for the event processor.
type EventProcessingConfig struct {
	Events []Event `yaml:"events"`
}

// Event contains a set of matching criteria and a set of scripts to run if the event occurs.
type Event struct {
	// ID is an identifier for this object. This is intended for free form text.
	ID string `yaml:"id"`

	// ObjectType is the type of docker object that is attached to this event.
	ObjectType string `yaml:"type"`

	// Action is the name of the action that needs to occur to trigger the associated scripts.
	Action string `yaml:"action"`

	// Attributes are a list of matching attributes that are present in the actor of the event.
	Attributes map[string]string `yaml:"attributes"`

	// Commands are the commands to execute.
	Commands []string `yaml:"commands"`
}
