/*
This file contains the code for the subprocess execution of the solutions.

Author: Shravan Asati
Originially Written: 26 June 2021
Last Edited: 29 June 2021
*/

package services

import (
	// "fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

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
		Log("error", "start failed")
		return "", err
	}

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		Log("error", "reading failed")
		return "", err
	}

	if we := cmd.Wait(); we != nil {
		Log("error", "wait failed")
	}

	return string(data), nil
}

// func main() {
// 	o, e := execute("node test.js")
// 	fmt.Println(o, e)
// }