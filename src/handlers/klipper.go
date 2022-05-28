package handlers

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Result struct {
		Status struct {
			DisplayStatus struct {
				Progress float32 `json:"progress"`
				Message  string  `json:"message"`
			} `json:"display_status"`
			EventTime float64 `json:"event_time"`
		} `json:"status"`
	} `json:"result"`
}

type SystemInfo struct {
	Result struct {
		SystemInfo struct {
			Python struct {
				Version       []string `json:"version"`
				VersionString string   `json:"version_string"`
			} `json:"python"`
			CpuInfo struct {
				CpuCount     int    `json:"cpu_count"`
				Bits         string `json:"bits"`
				Processor    string `json:"processor"`
				CpuDesc      string `json:"cpu_desc"`
				SerialNumber string `json:"serial_number"`
				HardwareDesc string `json:"hardware_desc"`
				Model        string `json:"model"`
				TotalMemory  uint32 `json:"total_memory"`
				MemoryUnits  string `json:"memory_units"`
			} `json:"cpu_info"`
			SdInfo struct {
				ManufacturerId   string `json:"manufacturer_id"`
				Manufacturer     string `json:"manufacturer"`
				OemId            string `json:"oem_id"`
				ProductName      string `json:"product_name"`
				ProductRevision  string `json:"product_revision"`
				SerialNumber     string `json:"serial_number"`
				ManufacturerDate string `json:"manufacturer_data"`
				Capacity         string `json:"capacity"`
				TotalBytes       uint32 `json:"total_bytes"`
			} `json:"sd_info"`
			Distribution struct {
				Name         string `json:"name"`
				Id           string `json:"id"`
				Version      string `json:"version"`
				VersionParts struct {
					Major       string `json:"major"`
					Minor       string `json:"minor"`
					BuildNumber string `json:"build_number"`
				} `json:"version_parts"`
				Like        string `json:"like"`
				Codename    string `json:"codename"`
				ReleaseInfo struct {
					Name      string `json:"name"`
					VersionId string `json:"version_id"`
					Id        string `json:"id"`
				} `json:"release_info"`
			} `json:"distribution"`
			Virtualization struct {
				VirtType       string `json:"virt_type"`
				VirtIdentifier string `json:"virt_identifier"`
			} `json:"virtualization"`
			Network           map[string]Network `json:"network"`
			AvailableServices []string           `json:"available_services"`
			ServiceState      map[string]Service `json:"service_state"`
		} `json:"system_info"`
	} `json:"result"`
}

type Service struct {
	Service struct {
		ActiveState string `json:"active_state"`
		SubState    string `json:"sub_state"`
	} `json:"service"`
}

type Network struct {
	Device struct {
		MacAddress  string      `json:"mac_address"`
		IpAddresses []IpAddress `json:"ip_address"`
	} `json:"device"`
}

type IpAddress struct {
	Family      string `json:"family"`
	Address     string `json:"address"`
	IsLinkLocal string `json:"is_link_local"`
}

func jsonDecode(endpoint string, typeInterface interface{}) error {
	resp, err := http.Get(endpoint)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(typeInterface)
}

func Get_status(endpoint string) *Status {
	status := new(Status)

	jsonDecode(endpoint+"/printer/objects/query?display_status", status)

	return status
}

func System_info(endpoint string) *SystemInfo {
	systemInfo := new(SystemInfo)

	jsonDecode(endpoint+"/machine/system_info", systemInfo)

	return systemInfo
}

func Job() {}

// /api/job - progress

func Printer() {}

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
