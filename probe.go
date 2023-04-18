/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
run test cases against the solutions.

Author: Shravan Asati
Originially Written: 30 June 2021
Last Edited: 6 October 2021
*/

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	serve "github.com/Ardent-Community/probe/services"
	"github.com/olekukonko/tablewriter"
	"github.com/thatisuday/commando"
)

const (
	NAME    string = "probe"
	VERSION string = "0.1.0"
)

type processEntry struct {
	username      string
	lang          string
	code          string
	testCasesFile string
}

type winner struct {
	Username string `json:"username"`
	Language string `json:"language"`
}

type winnerDB struct {
	winners []*winner
	sync.Mutex
}

func processor(entryCh *chan *processEntry, winners *winnerDB, doneCh chan bool) {
	for entry := range *entryCh {
		serve.Log("info", fmt.Sprintf("running %v's solution written in %v", entry.username, entry.lang))
		t := serve.Tester{
			Lang:          entry.lang,
			Code:          entry.code,
			TestCasesFile: entry.testCasesFile,
		}
		passed := t.PerformTests()
		if passed {
			serve.Log("success", fmt.Sprintf("%v's code passed", entry.username))
			winners.Lock()
			winners.winners = append(winners.winners, &winner{
				Username: entry.username,
				Language: entry.lang,
			})
			winners.Unlock()
		} else {
			serve.Log("failure", fmt.Sprintf("%v's code failed", entry.username))
		}

		doneCh <- true
	}
}

func tabulate(winners *winnerDB) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Username", "Language"})
	for _, winner := range winners.winners {
		table.Append([]string{winner.Username, winner.Language})
	}
	table.Render()
}

func run(challengeNumber, testCasesFile string) {
	// test solutions
	solutions := map[string]map[string]string{
		"username1": {
			"language": "python",
			"code":     "def solution(n):print(n * n)",
		},
		"username2": {
			"language": "javascript",
			"code":     "const solution = (n) => {console.log(n * n)}",
		},
		"wrong_username1": {
			"language": "python",
			"code":     "def solution(n): print(n * 2)",
		},
		"wrong_username2": {
			"language": "javascript",
			"code":     "const solution = (n) => {console.log(return n * n * 1)}",
		},
		"wrong_username3": {
			"language": "python",
			"code": "import time\nprint('haha')",
		},
		"wrong_username4": {
			"language": "javascript",
			"code": "import {time} from datetime;",
		},
	}

	// solutions := serve.GetSolutions(challengeNumber).Solutions

	// initialize winners db
	winners := &winnerDB{
		winners: []*winner{},
	}

	// initialize channels
	entryCh := make(chan *processEntry)
	doneCh := make(chan bool, 1)

	
	// spawn processor goroutines
	maxProcs := runtime.NumCPU()
	for i := 0; i < maxProcs; i++ {
		go processor(&entryCh, winners, doneCh)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// listener goroutine
	go func() {
		for i := 0; i < len(solutions); i++ {
			// fmt.Println("waiting")
			<-doneCh
		}
		wg.Done()
	}()

	// send entries to be processed
	for username, data := range solutions {
		entryCh <- &processEntry{
			username:      username,
			lang:          data["language"],
			code:          data["code"],
			testCasesFile: testCasesFile,
		}
	}

	// closing entry channel
	close(entryCh)
	wg.Wait()

	serve.ClearClutter()

	serve.Log("info", "\nThe winners are:\n")

	tabulate(winners)
}

func main() {
	run("1", `./examples/testcases.json`)
	fmt.Println(NAME, VERSION)

	commando.
		SetExecutableName(NAME).
		SetVersion(VERSION).
		SetDescription("probe is a CLI tool made to automate the process of solution validation for weekly challenges conducted in the Ardent-Community discord server.\n")

	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			commando.Parse([]string{"help"})
		})

	commando.
		Register("run").
		SetShortDescription("Validates solutions.").
		SetDescription("The `run` command does, in order, makes a request to the API, grabs the solutions, and concurrently runs all the solutions against the test cases.").
		AddArgument("challengeNumber", "The challenge number to validate the solutions for.", "").
		AddArgument("testCasesFile", "The path to the test cases file.", "").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {

			run(
				args["challengeNumber"].Value,
				args["testCasesFile"].Value,
			)

		})

	commando.Parse(nil)

}
