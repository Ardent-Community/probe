/*
The following code reads test cases from a json file and then executes them based on the language.

Author: Shravan Asati
Originally Written: 25 June 2021
Last Edited: 25 June 2021
*/

package main

import (
	"encoding/json"
	// "fmt"
	"path/filepath"
)

// Tester struct is the main struct which has methods related to testing the code.
type Tester struct {
	testCases     TestCases
	testCasesFile string
	code          string
	lang          string
}

// TestCases has python and javascript test cases, and is a part of the `Tester` struct.
type TestCases struct {
	PythonCases     map[string]string `json:pythonCases`
	JavascriptCases map[string]string `json:javascriptCases`
}

// getTestCases reads the test cases from the given file and assigns them to the `Tester` struct.
func (tester *Tester) getTestCases() {
	tc := TestCases{}
	fileContent := readFile(tester.testCasesFile)

	if e := json.Unmarshal([]byte(fileContent), &tc); e != nil {
		log("error", "unable to decode json into testcases")
		panic(e)
	}
	tester.testCases = tc
}

// testCode writes the given code to a file, executes the file, checks the output against the test case and returns a boolean variable whether the code passed or not.
func testCode(lang, code, in, out string) bool {
	// * writing code to a file
	filename := filepath.Join(getProbeDir(), randomFileName(lang))
	writeToFile(filename, code+"\n\n"+in)

	// * intialising output and error variables
	var output string
	var e error

	// * getting output and error
	if lang == "python" {
		output, e = execute("python3 " + filename)
	} else if lang == "javascript" {
		output, e = execute("node " + filename)
	}

	// * checking if the code passed
	if e != nil {
		// log("error", fmt.Sprintf("the code failed `%v` test", in))
		return false
	}
	if output != out {
		// log("error", fmt.Sprintf("the code failed `%v` test", in))
		return false
	}

	// log("info", fmt.Sprintf("the code passed `%v` test", in))
	return true
}

// performTests is the main tester function. It first gets the test cases, hunts the code for imports, exec and eval functions and tests the code, again returning a boolean variable.
func (tester *Tester) performTests() bool {
	// * getting test cases
	tester.getTestCases()

	if tester.lang == "python" {
		// * hunting for exec, eval and imports
		if hunt("python", tester.code) {
			// log("error", "the code breaks the rules")
			return false
		}

		// * executing the test cases
		for in, out := range tester.testCases.PythonCases {
			passed := testCode("python", tester.code, in, out)
			if !passed {
				return false
			}
		}
		return true

	} else if tester.lang == "javascript" {
		// * hunting for exec, eval and imports
		if hunt("javascript", tester.code) {
			// log("error", "the code breaks the rules")
			return false
		}

		// * executing the test cases
		for in, out := range tester.testCases.JavascriptCases {
			passed := testCode("javascript", tester.code, in, out)
			if !passed {
				return false
			}
		}
		return true

	} else {
		log("error", "invalid language type!")
	}

	return false
}

// func main() {
// 	ts := Tester{code: readFile("./test.js"), testCasesFile: "./example_testcases.json", lang: "javascript"}
// 	passed := ts.performTests()
// 	fmt.Println(passed)

// 	ts2 := Tester{code: readFile("./test.py"), testCasesFile: "./example_testcases.json", lang: "python"}
// 	passed2 := ts2.performTests()
// 	fmt.Println(passed2)
// 	clearClutter()
// }
