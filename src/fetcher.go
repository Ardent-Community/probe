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

	endpoint := "https://domain.name/api/solutions/" + challengeNumber + "?apiKey=" + apiKey

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

func tempSolutions() Result {
	body := []byte(`{
	"ok": true,
	"solutions":  {
		"username1": {
			"language": "python",
			"code": "from time import sleep\nsleep(1)"
		},
		"username2": {
			"language": "javascript",
			"code": "console.log('hey')"
		}
	}
}
	`)
	result := Result{}
	if e := json.Unmarshal(body, &result); e != nil {
		panic(e)
	}
	return result
}