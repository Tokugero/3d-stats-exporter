package handlers

import (
	"encoding/json"
	"net/http"
)

/*
TODO: Rebuild this to take inputs as "printer type" to determine which objects to use either in klipper or octoprint.
*/

func jsonDecode(endpoint string, typeInterface interface{}) error {
	/* Wrapper for calling upon Moonraker's API */
	resp, err := http.Get(endpoint)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(typeInterface)
}

func GetStatus(endpoint string) *MoonrakerStatus {
	status := new(MoonrakerStatus)

	jsonDecode(endpoint+"/printer/objects/query?display_status", status)

	return status
}

func GetSystemInfo(endpoint string) *MoonrakerSystemInfo {
	systemInfo := new(MoonrakerSystemInfo)

	jsonDecode(endpoint+"/machine/system_info", systemInfo)

	return systemInfo
}

func GetJob() {}

// /api/job - progress

func GetPrinter() {}

// /api/printer - heaters, print state

/*
	curl 192.168.20.230:7125/printer/objects/list | jq  >
	unmarshall these events into an anonymous object and parse everything into buckets to pass to prometheus
	let the user figure out what data is useful
		mcu
		print_stats
		bed_mesh
		heaters - configs in printer.cfg
		heater_bed - configs in printer.cfg
		fan - configs in printer.cfg
		fan_generic electronics - configs in printer.cfg
		fan_generic chamber - configs in printer.cfg
		heater_fan hotend - - configs in printer.cfg
		temperature_sensor octopus - configs in printer.cfg
		temperature_sensor chamber - configs in printer.cfg
		probe - configs in printer.cfg
		neopixel lighting - configs in printer.cfg
		motion_report
		system_stats
		toolhead - configs in printer.cfg
		extruder - configs in printer.cfg
*/
