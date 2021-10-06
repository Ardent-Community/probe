/*
This file contains the code for the subprocess execution of the solutions.

Author: Shravan Asati
Originially Written: 26 June 2021
Last Edited: 4 October 2021
*/

package services

// package main

import (
	// "fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

func getExecutionCommand(filename string) string {
	language := ""
	if strings.HasSuffix(filename, ".py") {
		language = "python"
	} else if strings.HasSuffix(filename, ".js") {
		language = "javascript"
	} else {
		panic("Invalid file extension")
	}

	switch language {
	case "python":
		switch runtime.GOOS {
		case "windows":
			return "python " + filename
		case "linux", "darwin":
			return "python3 " + filename
		default:
			panic("Unknown OS")
		}

	case "javascript":
		return "node " + filename

	default:
		panic("Unknown language")
	}

}

// execute executes the given command and returns the string output and error of the process.
func execute(command string) (string, error) {
	newCommand := strings.Fields(command)
	cmd := exec.Command(newCommand[0], newCommand[1:]...)

	stdout, e := cmd.StdoutPipe()
	if e != nil {
		Log("error", "stdout failed")
		return "", e
	}

	if err := cmd.Start(); err != nil {
		Log("error", "start failed "+err.Error())
		return "", err
	}

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		Log("error", "reading failed "+err.Error())
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		// Log("error", "wait failed " + err.Error())
		return "", err
	}

	return string(data), nil
}

// func main() {
// 	o, e := execute("node test.js")
// 	fmt.Println(o, e)
// }
