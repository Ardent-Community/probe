/*
This file contains the code for the subprocess execution of the solutions.

Author: Shravan Asati
Originially Written: 22 June 2021
Last Edited: 22 June 2021
*/

package main

import (
	"fmt"
	"os/exec"
	// "strings"
	"bytes"
)

func execute(command string ) (string, error) {
	// separated:= strings.Fields(command)
	cmd := exec.Command("python3", "./test.py")
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    outStr := (stdout.String())

	return outStr, err
}

func main() {
	out, e := execute("python3 ./test.py")
	if e != nil {
		panic(e)
	}
	if out == "25" {
		fmt.Println("hoorray")
	} else {
		fmt.Println("nope")
	}
}