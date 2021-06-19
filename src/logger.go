/*
This file contains code to log the stella messages and the actual program output, each in
different color codes.

Author: Shravan Asati
Originially Written: 19 June 2021
Last Edited: 19 June 2021
*/

package main

import (
	"fmt"
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorBlue  = "\033[36m"
	colorReset = "\033[0m"
)


func logger(severity, message string) {
	if severity == "info" {
		fmt.Println(colorGreen, message, colorReset)

	} else if severity == "error" {
		fmt.Println(colorGreen, message, colorReset)

	} else if severity == "success" {
		fmt.Println(colorGreen, message, colorReset)

	} else {
		fmt.Println("Fatal error. Invalid logging severity:", severity)
	}
}