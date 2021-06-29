/*
This file contains the code for hunting the code using regular expressions for imports, exec
and eval functions.

Author: Shravan Asati
Originally Written: 24 June 2021
Last Edited: 24 June 2021
*/

package main

import (
	"regexp"
)

// hunt returns true if it finds exec, eval or imports in the code.
func hunt(lang, code string) bool {
	if lang == "python" {
		execEvalPattern, _ := regexp.Compile("exec[(.)]|eval[(.)]")
		execEvalMatch := execEvalPattern.MatchString(code)

		if execEvalMatch {
			return true
		}

		importPattern, _ := regexp.Compile("import .")
		importMatch := importPattern.MatchString(code)

		if importMatch {
			return true
		}

	} else if lang == "javascript" {
		evalPattern, _ := regexp.Compile("eval[(.)]")
		evalMatch := evalPattern.MatchString(code)

		if evalMatch {
			return true
		}

		importPattern, _ := regexp.Compile(".require[(.)]|import.*?from.*?")
		importMatch := importPattern.MatchString(code)

		if importMatch {
			return true
		}

	} else {
		log("error", "invalid lang to hunt for: " + lang)
	}
	return false
}