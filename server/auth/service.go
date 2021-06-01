package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"colte.dev/env"
)

var URL = "https://github.com/login/oauth/access_token"

type Response struct {
	Error       string `json:"error"`
	AccessToken string `json:"access_token"`
}

func GetAccessToken(code string) (string, error) {
	req, _ := http.NewRequest("POST", URL, nil)
	q := req.URL.Query()
	q.Add("client_id", env.GITHUB_CLIENT_ID)
	q.Add("client_secret", env.GITHUB_CLIENT_SECRET)
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	responseBodyString, _ := ioutil.ReadAll(res.Body)

	var response Response
	err = json.Unmarshal(responseBodyString, &response)
	if err != nil {
		return "", err
	}

	return response.AccessToken, fmt.Errorf(response.Error)
}
