package handlers

type OctoprintStatus struct {
	SD struct {
		Ready bool `json:"ready"`
	} `json:"sd"`
	State struct {
		Error string `json:"error"`
		Flags struct {
			Cancelling    bool `json:"cancelling"`
			ClosedOrError bool `json:"closedOrError"`
			Error         bool `json:"error"`
			Finishing     bool `json:"finishing"`
			Operational   bool `json:"operational"`
			Paused        bool `json:"paused"`
			Pausing       bool `json:"pausing"`
			Printing      bool `json:"printing"`
			Ready         bool `json:"ready"`
			Resuming      bool `json:"resuming"`
			SdReady       bool `json:"sdReady"`
		} `json:"flags"`
		Text string `json:"text"`
	} `json:"state"`
	Temperature struct {
		A     map[string]OctoprintTemperature `json:"A"`
		E     map[string]OctoprintTemperature `json:"E"`
		P     map[string]OctoprintTemperature `json:"P"`
		W     map[string]OctoprintTemperature `json:"W"`
		Bed   map[string]OctoprintTemperature `json:"bed"`
		Tool0 map[string]OctoprintTemperature `json:"tool0"`
	} `json:"temperature"`
}

type OctoprintTemperature struct {
	Actual float64 `json:"actual"`
	Offset float64 `json:"offset"`
	Target float64 `json:"target"`
}

type OctoprintJobInfo struct {
	Job struct {
		AveragePrintTime   float64 `json:"averagePrintTime"`
		EstimatedPrintTime float64 `json:"estimatedPrintTime"`
		Filament           struct {
			Tool0 struct {
				Length float64 `json:"length"`
				Volume float64 `json:"volume"`
			} `json:"tool0"`
		} `json:"filament"`
		File struct {
			Date    uint32 `json:"date"`
			Display string `json:"display"`
			Name    string `json:"name"`
			Origin  string `json:"origin"`
			Path    string `json:"path"`
			Size    uint32 `json:"size"`
		} `json:"file"`
		LastPrintTime float64 `json:"lastPrintTime"`
		User          string  `json:"user"`
	} `json:"job"`
	Progress struct {
		Completion          float64 `json:"completion"`
		FilePos             float64 `json:"filepos"`
		PrintTime           float64 `json:"printTime"`
		PrintTimeLeft       float64 `json:"printTimeLeft"`
		PrintTimeLeftOrigin string  `json:"printTimeLeftOrigin"`
	} `json:"progress"`
	State string `json:"state"`
}
