/*
Unit tests for all the functions defined in the utils.go file.

Author: Shravan Asati
Originally Written: 26 July 2021
Last Edited: 26 July 2021
*/

package services

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRandomFileName(t *testing.T) {
	fileName := randomFileName("python")
	if !strings.HasSuffix(fileName, ".py") {
		t.Error("Wrong random filename generated for python.")
	}

	filename := randomFileName("javascript")
	if !strings.HasSuffix(filename, ".js") {
		t.Error("Wrong random filename generated for javascript.")
	}
}

type readWriteTest struct {
	fileName, content string
}

func TestReadWriteFile(t *testing.T) {
	readWriteTests := []readWriteTest{
		{"test.txt", "This is a test file."},
		{"test.json", `{"test": "This is a test file."}`},
	}

	for _, test := range readWriteTests {
		writeToFile(test.fileName, test.content)
		if readFile(test.fileName) != test.content {
			t.Error("Wrong content written to file.")
		}
	}

	t.Cleanup(func() {
		for _, test := range readWriteTests {
			cwd, e := os.Getwd()
			if e != nil {
				t.Error("Could not get current working directory.")
			}
			os.Remove(filepath.Join(cwd, test.fileName))
		}
	})
}

func Test_getProbeDir_ClearClutter(t *testing.T) {
	probeDir := getProbeDir()

	_, e := os.Stat(probeDir)
	if os.IsNotExist(e) {
		t.Error("Probe directory does not exist.")
	}

	writeToFile(filepath.Join(probeDir, "temp", "test.txt"), "Test")
	ClearClutter()

	contents, er := ioutil.ReadDir(filepath.Join(probeDir, "temp"))
	if er != nil {
		t.Error("unable to get contents of ~/.probe/temp: ", er.Error())
	}
	if len(contents) != 0 {
		t.Error("Clutter not cleared.")
	}
}
