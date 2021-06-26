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

type Tester struct {
	testCases TestCases
	testCasesFile string
	code string
	lang string
}

type TestCases struct {
	PythonCases     map[string]string `json:pythonCases`
	JavascriptCases map[string]string `json:javascriptCases`
}

func (tester *Tester) getTestCases() {
	tc := TestCases{}
	fileContent := readFile(tester.testCasesFile)

	if e := json.Unmarshal([]byte(fileContent), &tc); e != nil {
		log("error", "unable to decode json into testcases")
		panic(e)
	}
	tester.testCases = tc
}

func (tester *Tester) testCode() bool {
	passed := false
	// todo add regex hunting

	tester.getTestCases()

	if tester.lang == "python" {
		for in, out := range tester.testCases.PythonCases {
			filename := filepath.Join(getProbeDir(), randomFileName("python"))
			writeToFile(filename, tester.code + "\n\n" + in)
			output, e := execute("python3 " + filename)
			if e != nil {
				log("error", "The command failed to execute!")
				return false
			}
			if output != out {
				log("error", "The output isnt the same as expected!")
				return false
			}
			log("info", "passed first test")
		}
		passed = true

	} else if tester.lang == "javascript" {
		for in, out := range tester.testCases.JavascriptCases {
			filename := filepath.Join(getProbeDir(), randomFileName("python"))
			writeToFile(filename, tester.code + "\n\n" + in)
			output, e := execute("python3 " + filename)
			if e != nil {
				log("error", "The command failed to execute!")
				return false
			}
			if output != out {
				log("error", "The output isnt the same as expected!")
				return false
			}
		}
		passed = true
	
	} else {
		log("error", "invalid language type!")
	}

	clearClutter()
	return passed
}

// func main() {
// 	clearClutter()
// 	ts := Tester{code: readFile("./test.py"), testCasesFile: "./example_testcases.json", lang: "python"}
// 	passed := ts.testCode()
// 	fmt.Println(passed)
// 	// o, e := execute("python3 ./test.py")
// 	// fmt.Printf("Output: \n%v \nError: \n%v", o ,e)
// }
