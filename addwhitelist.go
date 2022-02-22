package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type APIGet struct {
	GithubOwnedAllowed bool     `json:"github_owned_allowed"`
	PatternsAllowed    []string `json:"patterns_allowed"`
	VerifiedAllowed    bool     `json:"verified_allowed"`
}

func main() {
	fmt.Println("Action starts")

	token := os.Args[1]
	pattern := os.Args[2]
	org := os.Args[3]
	fmt.Println("Action called with token " + token + " and pattern " + pattern +
		" for the organization " + org)

	/** API Call *****/

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/orgs/%s/actions/permissions/selected-actions", org), nil)
	if err != nil {
		return
	}

	body, err := doRequest(req, token)
	if err != nil {
		return
	}

	response := APIGet{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Print("There has been an error: " + err.Error())
		return
	}

	patternFound := false
	for _, value := range response.PatternsAllowed {

		if value == pattern {
			patternFound = true
			break
		}
	}

	if !patternFound {
		fmt.Println("Pattern not found in whitelist. We proceed to add it.")
		response.PatternsAllowed = append(response.PatternsAllowed, pattern)
		rb, err := json.Marshal(response)
		if err != nil {
			fmt.Println("There has been an error during the marshalling: " + err.Error())
			return
		}

		originalString := string(rb)

		req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.github.com/orgs/%s/actions/permissions/selected-actions", org),
			strings.NewReader(originalString))
		if err != nil {
			return
		}

		_, err = doRequest(req, token)
		if err != nil {
			return
		}

	} else {
		fmt.Println("Pattern found in patterns list! We do nothing")
	}

	/** End API Call ****/

	fmt.Println("Output -list of allowed patterns-: " + strings.Join(response.PatternsAllowed, ","))
	// Print the output
	fmt.Println("::set-output name=allowedPatterns::" + strings.Join(response.PatternsAllowed, ","))

}

func doRequest(req *http.Request, token string) ([]byte, error) {

	HTTPClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	req.Header.Add("Authorization", "token "+token)

	res, err := HTTPClient.Do(req)
	if err != nil {
		log.Println(res, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(res, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted &&
		res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNotFound &&
		res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
