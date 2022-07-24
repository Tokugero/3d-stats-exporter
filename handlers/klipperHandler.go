package handlers

type MoonrakerStatus struct {
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

type MoonrakerSystemInfo struct {
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
			Network           map[string]MoonrakerNetwork `json:"network"`
			AvailableServices []string                    `json:"available_services"`
			ServiceState      map[string]MoonrakerService `json:"service_state"`
		} `json:"system_info"`
	} `json:"result"`
}

type MoonrakerService struct {
	Service struct {
		ActiveState string `json:"active_state"`
		SubState    string `json:"sub_state"`
	} `json:"service"`
}

type MoonrakerNetwork struct {
	Device struct {
		MacAddress  string               `json:"mac_address"`
		IpAddresses []MoonrakerIpAddress `json:"ip_address"`
	} `json:"device"`
}

type MoonrakerIpAddress struct {
	Family      string `json:"family"`
	Address     string `json:"address"`
	IsLinkLocal string `json:"is_link_local"`
}
