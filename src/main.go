/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
run test cases against the solutions.

Author: Shravan Asati
Originially Written: 19 June 2021
Last Edited: 19 June 2021
*/

package main

func main() {
	f := "./" + randomFileName("javascript")
	writeToFile(f, "# hey")

	log("success", "The app is working.")
}
