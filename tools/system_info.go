package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/winservices"
	"time"
)

//go get github.com/shirou/gopsutil
/*
	cpu:  		CPU相关
	disk: 		磁盘相关
	docker:   docker相关
	host:			主机相关
	mem:			内存相关
	net:			网络相关
	process:  进程相关
*/
func main() {
	getMemoryInfo()
	getDiskInfo()
	getHostInfo()
	getProcessInfo()
	getWinServiceInfo()
	getCpuInfo()
}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// memory
func getMemoryInfo() {
	v, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)

	fmt.Println(v)

	swapMemory, _ := mem.SwapMemory()
	data, _ := json.MarshalIndent(swapMemory, "", " ")
	fmt.Println(string(data))
}

func getDiskInfo() {
	mapStat, _ := disk.IOCounters()
	for name, stat := range mapStat {
		fmt.Println(name)
		data, _ := json.MarshalIndent(stat, "", "  ")
		fmt.Println(string(data))
	}

	info, _ := disk.Usage("D:/myweb")
	data, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(data))
}

func getHostInfo() {
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))

	version, _ := host.KernelVersion()
	fmt.Println(version)

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)

	users, _ := host.Users()
	for _, user := range users {
		data, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(data))
	}
}

func getProcessInfo() {
	var rootProcess *process.Process
	processes, _ := process.Processes()
	for _, p := range processes {
		if p.Pid == 0 {
			rootProcess = p
			break
		}
	}

	fmt.Println(rootProcess)

	fmt.Println("children:")
	children, _ := rootProcess.Children()
	for _, p := range children {
		fmt.Println(p)

	}
}

func getWinServiceInfo() {
	services, _ := winservices.ListServices()

	for _, service := range services {
		newservice, err := winservices.NewService(service.Name)
		if err != nil {
			fmt.Println(service.Name, err)
			continue
		}
		newservice.GetServiceDetail()
		fmt.Println("Name:", newservice.Name, "Binary Path:", newservice.Config.BinaryPathName, "State: ", newservice.Status.State)
	}
}
