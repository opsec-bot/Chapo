package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"
)

const (
	APIKEY = ""
	API    = "http://45.79.53.89/data"
)

type Data struct {
	Usernames []string `json:"Username"`
	PublicIP  []string `json:"Public IP"`
}

func check() {
	// Make API request
	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("api-key", APIKEY)
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

func Washer() {
	// Get current IP address
	resp, err := http.Get("http://checkip.amazonaws.com")
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	currentIP := strings.TrimSpace(string(body))

	// Make API request
	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("api-key", APIKEY)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
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

	// Check if IP address is in the list
	for _, ip := range data.PublicIP {
		if currentIP == ip {
			fmt.Println("FOUND")
			return
		}
	}
	fmt.Println("NOT FOUND")
}
