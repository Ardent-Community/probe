/*
The following code reads test cases from a json file and then executes them based on the language.

Author: Shravan Asati
Originally Written: 28 June 2021
Last Edited: 29 June 2021
*/

package services
// package main

import (
	"encoding/json"
	// "fmt"
	"path/filepath"
)

// Tester struct is the main struct which has methods related to testing the code.
type Tester struct {
	TestCases     TestCases
	TestCasesFile string
	Code          string
	Lang          string
}

// TestCases has python and javascript test cases, and is a part of the `Tester` struct.
type TestCases struct {
	PythonCases     map[string]string `json:"pythonCases"`
	JavascriptCases map[string]string `json:"javascriptCases"`
}

// getTestCases reads the test cases from the given file and assigns them to the `Tester` struct.
func (tester *Tester) getTestCases() {
	tc := TestCases{}
	fileContent := readFile(tester.TestCasesFile)

	if e := json.Unmarshal([]byte(fileContent), &tc); e != nil {
		Log("error", "unable to decode json into testcases")
		panic(e)
	}
	tester.TestCases = tc
}

// testCode writes the given code to a file, executes the file, checks the output against the test case and returns a boolean variable whether the code passed or not.
func testCode(lang, code, in, out string) bool {
	// * writing code to a file
	filename := filepath.Join(getProbeDir(), "temp", randomFileName(lang))
	writeToFile(filename, code+"\n\n"+in)

	// * getting output and error
	output, e := execute(getExecutionCommand(filename))

	// * checking if the code passed
	if e != nil {
		// Log("error", fmt.Sprintf("the code failed `%v` test", in))
		return false
	}
	if output != out {
		// Log("error", fmt.Sprintf("the code failed `%v` test", in))
		return false
	}

	// Log("info", fmt.Sprintf("the code passed `%v` test", in))
	return true
}

// PerformTests is the main tester function. It first gets the test cases, hunts the code for imports, exec and eval functions and tests the code, again returning a boolean variable.
func (tester *Tester) PerformTests() bool {
	// * getting test cases
	tester.getTestCases()

	if tester.Lang == "python" {
		// * hunting for exec, eval and imports
		if hunt("python", tester.Code) {
			// Log("error", "the code breaks the rules")
			return false
		}

		// * executing the test cases
		for in, out := range tester.TestCases.PythonCases {
			passed := testCode("python", tester.Code, in, out)
			if !passed {
				return false
			}
		}
		return true

	} else if tester.Lang == "javascript" {
		// * hunting for exec, eval and imports
		if hunt("javascript", tester.Code) {
			// Log("error", "the code breaks the rules")
			return false
		}

		// * executing the test cases
		for in, out := range tester.TestCases.JavascriptCases {
			passed := testCode("javascript", tester.Code, in, out)
			if !passed {
				return false
			}
		}
		return true

	} else {
		Log("error", "invalid language type!")
	}

	return false
}

// func main() {
// 	ts := Tester{Code: ReadFile("../examples/test.js"), TestCasesFile: "../examples/testcases.json", Lang: "javascript"}
// 	passed := ts.PerformTests()
// 	fmt.Println(passed)

// 	ts2 := Tester{Code: ReadFile("../examples/test.py"), TestCasesFile: "../examples/testcases.json", Lang: "python"}
// 	passed2 := ts2.PerformTests()
// 	fmt.Println(passed2)
// 	ClearClutter()
// }
