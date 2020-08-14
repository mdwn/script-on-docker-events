package config

import (
	"testing"

	"github.com/go-test/deep"
)

func TestLoadData(t *testing.T) {
	tests := []struct {
		name        string
		file        string
		expected    EventProcessingConfig
		shouldError bool
	}{
		{
			name: "regular parse",
			file: "_testdata/regular.yaml",
			expected: EventProcessingConfig{
				Events: []Event{
					{
						ID:         "test event 1",
						ObjectType: "container",
						Action:     "start",
						Commands: []string{
							"echo \"1\"",
							"echo \"2\"",
						},
					},
					{
						ID:         "test event 2",
						ObjectType: "container",
						Action:     "die",
						Commands: []string{
							"echo \"3\"",
							"echo \"4\"",
						},
					},
				},
			},
		},
		{
			name:        "empty command",
			file:        "_testdata/empty-command.yaml",
			shouldError: true,
		},
	}

	for _, test := range tests {
		err := LoadConfig(test.file)

		if (err != nil) != test.shouldError {
			t.Errorf("test %s error was %t when it shouldn't have, err: %v", test.name, test.shouldError, err)
		}

		if !test.shouldError {
			eventProcessingConfig, err := GetEventProcessingConfig()

			if err != nil {
				t.Errorf("test %s error getting config: %v", test.name, err)
			}

			if diff := deep.Equal(*eventProcessingConfig, test.expected); diff != nil {
				t.Errorf("test %s expected output did not match produced output: %v", test.name, diff)
			}
		}
	}
}
