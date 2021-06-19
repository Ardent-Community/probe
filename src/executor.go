/*
This file contains the code for the subprocess execution of the solutions.

Author: Shravan Asati
Originially Written: 19 June 2021
Last Edited: 19 June 2021
*/

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeFile(username, filename string, testCases []string) {
	executable := ""

	if strings.HasSuffix(filename, ".py") {
		executable = "python3"
	} else if strings.HasSuffix(filename, ".js") {
		executable = "node"
	} else {
		log("error", "invalid filename!")
		return
	}

	cmd := exec.Command(executable, filename)
	out, err := cmd.Output()

	if err != nil {
		log("error", "the command failed to execute!")
		fmt.Println(err)
		return
	}

	if string(out) == "test" {
		log("success", "the test passed")
	}
}
