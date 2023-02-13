package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"
)

type Data struct {
	Usernames []string `json:"Username"`
}

func check() {
	// Make API request
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://45.79.53.89/data", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("api-key", "KEYHERE")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse JSON response
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Get current username
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}
	username := strings.Split(currentUser.Username, `\`)[1]

	// Check if username is in the list
	for _, u := range data.Usernames {
		if username == u {
			fmt.Println("FOUND")
			return
		}
	}
	fmt.Println("NOT FOUND")
}
