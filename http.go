package main

import (
	"net/http"
)

func DoRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		print("Access token expired. Refetching one...\n")
		token = GetAccessToken(etpRt) // no longer a pointer
		req.Header.Set("Authorization", "Bearer "+token)
		return DoRequest(req)
	}
	return resp, err
}
