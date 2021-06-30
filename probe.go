/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
run test cases against the solutions.

Author: Shravan Asati
Originially Written: 30 June 2021
Last Edited: 30 June 2021
*/

package main

import (
	"sync"
	serve "github.com/Ardent-Community/probe-cli/services"
)

func main() {
	wg := sync.WaitGroup{}
	solutions := serve.GetSolutions("1").Solutions
	wg.Add(len(solutions))

	for username, data := range solutions {
		go func(username string, data map[string]string) {
			lang := data["language"]
			code := data["code"]
			serve.Log("info", "running "+ username + "'s solution written in " + lang)
			
			t := serve.Tester{
				Lang: lang,
				Code: code,
				TestCasesFile: "./examples/testcases.json",
			}
			passed := t.PerformTests()
			if passed {
				serve.Log("success", username + "'s code passed")
			} else {
				serve.Log("info", username + "'s code failed")
			}
			wg.Done()
		}(username, data)
	}

	wg.Wait()
	serve.ClearClutter()
}
