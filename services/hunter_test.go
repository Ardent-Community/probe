/*
Unit tests for the hunter.

Author: Shravan Asati
Originally Written: 26th July 2021
last Edited: 26th July 2021
*/

package services

import "testing"

type hunterTest struct {
	lang string
	code string
	pass bool
}

func TestHunter(t *testing.T) {
	hunterTests := []hunterTest{
		// python tests
		{"python", "# this is just an example\n\n \t\t import zen of python", true},
		{"python", "# again this is an example\n\n\nprint('hello') from time import sleep", true},
		{"python", "this test will pass \n\n\n\n\n \t\t def add(a,b): return a + b", false},

		// javascript tests
		{"javascript", "import {x, y} from z; // haha this will fail", true},
		{"javascript", "\n\n const something = none; undefined \n\n", false},
		{"javascript", "const http = require('http')\n\nhttp.get() \n\n", true},
		{"javascript", "\n\nconsole.log('this test will pass')\n\n", false},
	}

	for _, test := range hunterTests {
		result := hunt(test.lang, test.code)
		if test.pass != result {
			t.Errorf("Hunter test failed.\n Language: %s \nCode: %s \nExpected: %v \nGot: %v", test.lang, test.code, test.pass, result)
		}
	}
}
