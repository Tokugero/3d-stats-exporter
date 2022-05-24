package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Status struct {
	Result struct {
		Status struct {
			DisplayStatus struct {
				Progress float32
				Message  string
			}
			EventTime float64
		}
	}
}

func get_status(endpoint string, status interface{}) error {
	// apikey := "0572f6d72a954316b397c3d54820c30a"

	resp, err := http.Get(endpoint + "/printer/objects/query?display_status")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(status)
}

func main() {
	endpoint := "http://192.168.20.230:7125"

	status := new(Status)
	get_status(endpoint, status)

	fmt.Print(status.Result.Status.DisplayStatus.Progress)

	return
}
