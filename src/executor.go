/*
This file contains the code for the subprocess execution of the solutions.

Author: Shravan Asati
Originially Written: 22 June 2021
Last Edited: 22 June 2021
*/

package main

import (
	// "fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func execute(command string) (string, error) {
	newCommand := strings.Fields(command)
	cmd := exec.Command(newCommand[0], newCommand[1:]...)

	stdout, e := cmd.StdoutPipe()
	if e != nil {
		log("error", "stdout failed")
		return "", e
	}

	if err := cmd.Start(); err != nil {
		log("error", "start failed")
		return "", err
	}

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		log("error", "reading failed")
		return "", err
	}

	if we := cmd.Wait(); we != nil {
		log("error", "wait failed")
	}

	return string(data), nil
}

// func main() {
// 	o, e := execute("node test.js")
// 	fmt.Println(o, e)
// }