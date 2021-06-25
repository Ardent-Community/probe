/*
The following code reads test cases from a json file and then executes them based on the language.

Author: Shravan Asati
Originally Written: 25 June 2021
Last Edited: 25 June 2021
*/

package main

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"strings"
)

type Tester struct {
	testCases TestCases
}

type TestCases struct {
	PythonCases     map[string]string `json:pythonCases`
	JavascriptCases map[string]string `json:javascriptCases`
}

func (tester *Tester) getTestCases(testfile string) {
	tc := TestCases{}
	fileContent := readFile(testfile)

	if e := json.Unmarshal([]byte(fileContent), &tc); e != nil {
		log("error", "unable to decode json into testcases")
		panic(e)
	}
	tester.testCases = tc
}

func (tester *Tester) testCode(codefile, testfile string) bool {
	passed := false
	// todo add regex hunting

	tester.getTestCases(testfile)
	lang := ""
	if strings.HasSuffix(codefile, ".py") {
		lang = "python"
	} else if strings.HasSuffix(codefile, ".js") {
		lang = "javascript"
	} else {
		panic("invalid language type")
	}

	if lang == "python" {
		for in, out := range tester.testCases.PythonCases {
			writeToFile(codefile, readFile(codefile)+"\n"+in)
			output, e := execute("python3 " + codefile)
			if e != nil {
				log("error", "The command failed to execute!")
				writeToFile(codefile, strings.ReplaceAll(codefile, in, ""))
				return false
			}
			if output != out {
				log("error", "The output isnt the same as expected!")
				writeToFile(codefile, strings.ReplaceAll(codefile, in, ""))
				return false
			}
			writeToFile(codefile, strings.ReplaceAll(codefile, in, ""))
		}
		passed = true

	} else if lang == "javascript" {
		for in, out := range tester.testCases.JavascriptCases {
			writeToFile(codefile, readFile(codefile)+"\n"+in)
			output, e := execute("node " + codefile)
			if e != nil {
				log("error", "The command failed to execute!")
				writeToFile(codefile, strings.ReplaceAll(codefile, in, ""))
				return false
			}
			if output != out {
				log("error", "The output isnt the same as expected!")
				writeToFile(codefile, strings.ReplaceAll(codefile, in, ""))
				return false
			}
			writeToFile(codefile, strings.ReplaceAll(codefile, in, ""))
		}
		passed = true
	
	} else {
		log("error", "invalid language type!")
	}

	clearClutter()
	return passed
}

func main() {
	ts := Tester{}
	passed := ts.testCode("./test.py", "./example_testcases.json")
	fmt.Println(passed)
	// o, e := execute("python3 ./test.py")
	// fmt.Printf("Output: \n%v \nError: \n%v", o ,e)
}
