/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
runs test cases against the solutions.

Author: Shravan Asati
Originially Written: 22 June 2021
Last Edited: 22 June 2021
*/

package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	solutions := getSolutions("1").Solutions
	wg.Add(len(solutions))

	for username, data := range solutions {
		go func(username string, data map[string]string) {
			lang := data["language"]
			code := data["code"]
			log("info", "running "+ username + "'s solution written in " + lang)
			
			t := Tester{
				lang: lang,
				code: code,
				testCasesFile: "./example_testcases.json",
			}
			passed := t.performTests()
			if passed {
				log("success", username + "'s code passed")
			} else {
				log("info", username + "'s code failed")
			}
			wg.Done()
		}(username, data)
	}

	wg.Wait()
	clearClutter()
}
