/*
Fetches solutions from the API and returns them as part of the Result struct.

Author: Shravan Asati
Originially Written: 25 June 2021
Last Edited: 29 June 2021
*/

package services 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// Result struct contains all the parameters returned by the solutions API.
type Result struct {
	Solutions map[string]map[string]string `json:solutions`
}

// GetSolutions makes a request to the API server and returns a `Result` struct filled with solutions.
func GetSolutions(challengeNumber string) Result {
	// * getting the api key
	apiKey := os.Getenv("PROBE_API_KEY")
	if apiKey == "" {
		panic("API KEY NOT FOUND")
	}

	// * defining the url and http client
	endpoint := "http://127.0.0.1:5000/api/solutions/" + challengeNumber 
	client := http.Client{}

	// * making a request object
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		panic("request not SENT")
	}

	// * defining headers
	req.Header.Set("User-Agent", "probe-cli")
	req.Header.Add("API-KEY", apiKey)

	// * performing the request
	res, err := client.Do(req)
	if err != nil {
		panic("REQUEST NOT SENT")
	}
	defer res.Body.Close()

	// * reading response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("unable to decode data in json")
	}

	// * writing response to json
	result := Result{}
	if e := json.Unmarshal(body, &result); e != nil {
		panic("unable to write json to struct")
	}

	return result
}