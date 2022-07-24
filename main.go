package main

import (
	"fmt"

	"github.com/Tokugero/3d-stats-exporter/handlers"
)

func main() {
	/* 	klipper := handlers.Klipper{
		Endpoint: "http://192.168.20.230:7125",
	} */

	octoprint := handlers.Octoprint{
		Endpoint: "http://192.168.20.231",
		Token:    "3E06AEF967F24504B3450770DD9598BD",
	}

	octoprint_output := octoprint.GetJob()

	fmt.Print(octoprint_output.Progress)
	//fmt.Print(octoprint_status.Result.Status.DisplayStatus.Progress)
	//fmt.Print(octoprint_system_info.Result.SystemInfo.CpuInfo)

	return
}
