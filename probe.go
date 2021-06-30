/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
run test cases against the solutions.

Author: Shravan Asati
Originially Written: 30 June 2021
Last Edited: 30 June 2021
*/

package main

import (
	"fmt"
	serve "github.com/Ardent-Community/probe/services"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	solutions := serve.GetSolutions("1").Solutions

	winners := []string{}
	wg.Add(len(solutions))

	for username, data := range solutions {
		go func(username string, data map[string]string) {
			lang := data["language"]
			code := data["code"]
			serve.Log("info", "running "+username+"'s solution written in "+lang)

			t := serve.Tester{
				Lang:          lang,
				Code:          code,
				TestCasesFile: "./examples/testcases.json",
			}
			passed := t.PerformTests()
			if passed {
				serve.Log("success", username+"'s code passed")
				winners = append(winners, username)
			} else {
				serve.Log("failure", username+"'s code failed")
			}
			wg.Done()
		}(username, data)
	}

	wg.Wait()
	serve.ClearClutter()

	serve.Log("info", "\nThe winners are:\n")
	for i, v := range winners {
		serve.Log("success", fmt.Sprintf("%v. %v", i+1, v))
	}
}
