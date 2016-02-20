package main

import (
	"encoding/json"
	"github.com/pivotal-golang/bytefmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
	"strconv"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	mem, _ := mem.VirtualMemory()
	load, _ := load.LoadAvg()
	diskUsage, _ := disk.DiskUsage(*diskDrive)

	upTime, _ := host.Uptime()
	upTimeDays := (((upTime / 60) / 60) / 24)

	totalMem := mem.Total / bytefmt.MEGABYTE
	availableMem := mem.Available / bytefmt.MEGABYTE

	diskTotal := diskUsage.Total / bytefmt.GIGABYTE
	diskUsed := diskUsage.Used / bytefmt.GIGABYTE

	stat := Stat{
		TotalMemory:         strconv.Itoa(int(totalMem)) + " MB",
		AvailableMemory:     strconv.Itoa(int(availableMem)) + " MB",
		UsedMemPercentage:   mem.UsedPercent,
		DiskTotal:           strconv.Itoa(int(diskTotal)) + " GB",
		DiskUsed:            strconv.Itoa(int(diskUsed)) + " GB",
		DiskUsagePercentage: diskUsage.UsedPercent,
		UpTimeDays:          strconv.Itoa(int(upTimeDays)) + " Days",
		CPULoad:             load.String(),
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(stat)

	if err != nil {
		log.Error(err)
	}
}
