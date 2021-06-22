/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
run test cases against the solutions.

Author: Shravan Asati
Originially Written: 19 June 2021
Last Edited: 19 June 2021
*/

package main

import "fmt"

func main() {
	// f := "./" + randomFileName("javascript")
	// writeToFile(f, "# hey")

	// log("success", "The app is working.")
	r := test()
	fmt.Println(r.Solutions["username1"]["language"] + "\n", r.Solutions["username1"]["code"])
}
