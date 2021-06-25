/*
This file contains the code for the subprocess execution of the solutions.

Author: Shravan Asati
Originially Written: 22 June 2021
Last Edited: 22 June 2021
*/

package main

import (
	"os/exec"
	"strings"
	"bytes"
)

func execute(command string ) (string, error) {
	separated:= strings.Fields(command)

	cmd := exec.Command(separated[0], separated...)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    outStr := (stdout.String())

	return outStr, err
}
