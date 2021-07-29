/*
Unit tests for the hunter.

Author: Shravan Asati
Originally Written: 26th July 2021
last Edited: 26th July 2021
*/

package services

import (
	"testing"
)

type executorTest struct {
	command string
	output  string
	err     bool
}

func TestExecute(t *testing.T) {
	tests := []executorTest{
		{"sh -c \"echo 'this is a test'\"", "this is a test", false},
		{"sh -c \"echo 'this is fun'\"", "this is fun", false},
		{"cat ./something", "", true},
	}

	for _, test := range tests {
		output, err := execute(test.command)
		if err != nil && !test.err {
			t.Errorf("Error: %s", err)
		}
		if output != test.output {
			t.Errorf("Expected: %s, Actual: %s", test.output, output)
		}
	}
}