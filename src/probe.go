/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
runs test cases against the solutions.

Author: Shravan Asati
Originially Written: 22 June 2021
Last Edited: 22 June 2021
*/

package main



// func main() {
// 	probeDir := getProbeDir()
// 	_, e := os.Stat(probeDir)
// 	if os.IsNotExist(e) {
// 		os.Mkdir(probeDir, os.ModePerm)
// 	}
// 	os.Chdir(probeDir)
// 	clearClutter()


// 	solutions := getSolutions("1").Solutions

// 	for username, data := range solutions {
// 		lang := data["language"]
// 		code := data["code"]
// 		log("info", "running "+ username + "'s solution written in " + lang)
// 		filename := username + randomFileName(lang)
// 		writeToFile(filename, code)
// 		execute(filename)
// 	}
// }
