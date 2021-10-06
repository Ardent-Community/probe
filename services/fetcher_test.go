/*
Unit tests for the fetcher.

Author: Shravan Asati
Originally Written: 29 July 2021
Last Edited: 29 July 2021
*/

package services

import (
	"reflect"
	"testing"
)

func Test_getAPIKey(t *testing.T) {
	expected := "super-secret"
	got := getAPIKey()
	if got != expected {
		t.Errorf("getAPIKey() = %v, want %v", got, expected)
	}
}

func Test_decodeJSON(t *testing.T) {
	expected := Result{
		Ok: true,
		Solutions: map[string]map[string]string{
			"username1": {
				"language": "python",
				"code":     "def solution(n): print(n * n)",
			},
			"username2": {
				"language": "javascript",
				"code":     "const solution = (n) => {console.log(n * n)}",
			},
			"wrong_username1": {
				"language": "python",
				"code":     "def solution(n): print(n * n/n)",
			},
			"wrong_username2": {
				"language": "javascript",
				"code":     "const solution = (n) => {console.log(return n * n * n)}",
			},
		},
	}

	got := decodeJSON([]byte(readFile("../examples/response.json")))

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("decodeJSON() = %v, want %v", got, expected)
	}
}
