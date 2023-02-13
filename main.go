package Chapo

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

type AntiVM struct {
	ApiKey    string
	Blacklist struct {
		ComputerNames []string
		Usernames     []string
		IPs           []string
		Macs          []string
		GUIDS         []string
		UUIDS         []string
	}
}

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    struct {
		ComputerNames []string `json:"computer_names"`
		Usernames     []string `json:"usernames"`
		IPAddresses   []string `json:"ip_addresses"`
		MacAddresses  []string `json:"mac_addresses"`
		MachineGuids  []string `json:"machine_guids"`
		MachineUuids  []string `json:"machine_uuids"`
	} `json:"data"`
}

func New(apiKey string) *AntiVM {
	return &AntiVM{
		ApiKey: apiKey,
	}
}

func (c *AntiVM) GetBlacklistedData() (bool, error) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI("http://45.79.53.89/blacklist")
	req.Header.SetMethod("GET")
	req.Header.Set("X-ApiKey", c.ApiKey)

	err := fasthttp.Do(req, res)
	if err != nil {
		return false, err
	}

	var response Response
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		return false, err
	}

	if response.Error {
		return false, errors.New(response.Message)
	}

	c.Blacklist.ComputerNames = response.Data.ComputerNames
	c.Blacklist.Usernames = response.Data.Usernames
	c.Blacklist.IPs = response.Data.IPAddresses
	c.Blacklist.Macs = response.Data.MacAddresses
	c.Blacklist.GUIDS = response.Data.MachineGuids
	c.Blacklist.UUIDS = response.Data.MachineUuids

	return true, nil
}
