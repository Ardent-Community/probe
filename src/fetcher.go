/*
Fetches solutions from the API and returns them as part of the Result struct.

Author: Shravan Asati
Originially Written: 22 June 2021
Last Edited: 23 June 2021
*/

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Result struct {
	Solutions map[string]map[string]string `json:solutions`
}

func getSolutions(challengeNumber string) Result {
	apiKey := os.Getenv("PROBE_API_KEY")
	if apiKey == "" {
		panic("API KEY NOT FOUND")
	}

	endpoint := "http://127.0.0.1:5000/api/solutions/" + challengeNumber + "?apiKey=" + apiKey

	res, err := http.Get(endpoint)
	if err != nil {
		panic("REQUEST NOT SENT")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("unable to decode data in json")
	}

	result := Result{}

	if e := json.Unmarshal(body, &result); e != nil {
		panic("unable to write json to struct")
	}

	return result
}