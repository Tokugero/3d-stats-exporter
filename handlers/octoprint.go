package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://docs.octoprint.org/en/master

type OctoprintHandler interface {
	jsonDecode()
	GetStatus()
	GetSystemInfo()
	GetJob()
	GetPrinter()
}

type Octoprint struct {
	Endpoint string
	Token    string
}

func (o *Octoprint) jsonDecode(method string, uri string, typeInterface interface{}) error {
	req, err := http.NewRequest(method, o.Endpoint+uri, nil)
	if err != nil {
		return fmt.Errorf("Failed to create request to %s with %s", uri, err.Error())
	}
	req.Header.Set("X-Api-Key", o.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to send request to %s with %s", uri, err.Error())
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body from %s with %s", uri, err.Error())
	}
	fmt.Println(string(b))

	return json.NewDecoder(resp.Body).Decode(typeInterface)
}

func (o *Octoprint) GetJob() *OctoprintJobInfo {
	job := new(OctoprintJobInfo)

	o.jsonDecode("GET", "/api/job", job)

	fmt.Print(job)
	return job
}

func (o *Octoprint) GetPrinter() {}
