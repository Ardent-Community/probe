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
	"github.com/thatisuday/commando"
	"sync"
)

const (
	NAME    string = "probe"
	VERSION string = "0.1.0"
)

func main() {
	fmt.Println(NAME, VERSION)

	commando.
		SetExecutableName(NAME).
		SetVersion(VERSION).
		SetDescription("probe is a CLI tool made to automate the process of solution validation for weekly challenges conducted in the Ardent-Community.\n")

	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("\nExecute `probe -h` for help.")
		})

	commando.
		Register("run").
		SetShortDescription("Validates solutions.").
		SetDescription("The `run` command does, in order, makes a request to the API, grabs the solutions, and concurrently runs all the solutions against the test cases.").
		AddArgument("challengeNumber", "The challenge number to validate the solutions for.", "").
		AddArgument("testCasesFile", "The path to the test cases file.", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			wg := sync.WaitGroup{}
			solutions := serve.GetSolutions(args["challengeNumber"].Value).Solutions

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
						TestCasesFile: args["testCasesFile"].Value,
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
		})

	commando.Parse(nil)

}
