// Test file for server.go

package main

import (
	"testing"
)

func TestProcessCommand(t *testing.T) {
	tests := []struct {
		command  string
		expected string
	}{
		{"Test hardcoded response", "+PONG\r\n"},
		//{"PING", "+PONG\r\n"},
		//{"ECHO Hello", "+PONG\r\n"}, // Assuming the function always returns "+PONG\r\n"
		//{"SET key value", "+PONG\r\n"},
	}

	for _, test := range tests {
		t.Run(test.command, func(t *testing.T) {
			result := processCommand(test.command)
			if result != test.expected {
				t.Errorf("processCommand(%q) = %q; want %q", test.command, result, test.expected)
			}
		})
	}
}
