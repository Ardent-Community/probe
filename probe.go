/*
The main functioning of the probe cli. Spawns multiple goroutines to asynchronously
run test cases against the solutions.

Author: Shravan Asati
Originially Written: 30 June 2021
Last Edited: 30 June 2021
*/

package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"

	serve "github.com/Ardent-Community/probe/services"
	"github.com/thatisuday/commando"
	// "github.com/thatisuday/commando"
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
	Username string
	Language string
}

type winnerDB struct {
	winners []*winner
	sync.Mutex
}

func processor(entryCh *chan *processEntry, mutateCh *chan *winner, doneCh chan bool) {
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
			*mutateCh <- &winner{
				Username: entry.username,
				Language: entry.lang,
			}
		} else {
			serve.Log("failure", fmt.Sprintf("%v's code failed", entry.username))
		}

		doneCh <- true
	}
	// close(*mutateCh)
}

func mutater(db *winnerDB, mutateCh *chan *winner) {
	for winner := range *mutateCh {
		db.Lock()
		db.winners = append(db.winners, winner)
		db.Unlock()
	}
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
			"code":     "const solution = (n) => {console.log(return n * n * n)}",
		},
	}

	// solutions := serve.GetSolutions(challengeNumber).Solutions

	// initialize winners as atomic
	winners := &winnerDB{
		winners: []*winner{},
	}

	// initialize channels
	entryCh := make(chan *processEntry)
	mutateCh := make(chan *winner)
	doneCh := make(chan bool)

	// spawn processor and mutater goroutines
	go mutater(winners, &mutateCh)
	for i := 0; i < runtime.NumCPU(); i++ {
		go processor(&entryCh, &mutateCh, doneCh)
	}

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

	for i := 0; i < len(solutions); i++ {
		<-doneCh
	}

	// serve.ClearClutter()

	serve.Log("info", "\nThe winners are:\n")

	result, e := json.MarshalIndent(winners.winners, "", "  ")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(result))
}

func main() {
	// run("1", `./examples/testcases.json`)
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
