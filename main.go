package main

import (
	"fmt"

	"github.com/Tokugero/3d-stats-exporter/handlers"
)

func main() {
	endpoint := "http://192.168.20.230:7125"

	status := handlers.Get_status(endpoint)
	system_info := handlers.System_info(endpoint)
	// system_info(endpoint)

	fmt.Print(status.Result.Status.DisplayStatus.Progress)
	fmt.Print(system_info.Result.SystemInfo.CpuInfo)

	return
}
