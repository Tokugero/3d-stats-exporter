package handlers

import (
	"encoding/json"
	"net/http"
)

/*
TODO: Add Klipper core functionality to bypass API.
*/
// Object declarations for moonraker's JSON API

type KlipperHandler interface {
	jsonDecode()
	GetStatus()
	GetSystemInfo()
	GetJob()
	GetPrinter()
}

type Klipper struct {
	Endpoint string
}

func (k *Klipper) jsonDecode(uri string, typeInterface interface{}) error {
	/* Wrapper for calling upon Moonraker's API */
	resp, err := http.Get(k.Endpoint + uri)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(typeInterface)
}

func (k *Klipper) GetStatus() *MoonrakerStatus {
	status := new(MoonrakerStatus)

	k.jsonDecode("/printer/objects/query?display_status", status)

	return status
}

func (k *Klipper) GetSystemInfo() *MoonrakerSystemInfo {
	systemInfo := new(MoonrakerSystemInfo)

	k.jsonDecode("/machine/system_info", systemInfo)

	return systemInfo
}

func (k *Klipper) GetJob() {}

// /api/job - progress

func (k *Klipper) GetPrinter() {}

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
