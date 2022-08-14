/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: hardware
 Date: 8/13/22 1:46 PM
*/
package system

import (
	"github.com/kavanahuang/log"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"runtime"
)

type HardwareInfo struct {
	CPU             float64 `json:"cpu"`
	DiskFree        uint64  `json:"diskFree"`
	DiskUsedPercent float64 `json:"diskUsedPercent"`
	MemUsedPercent  float64 `json:"memUsedPercent"`
	GoMemory        int64   `json:"goMemory"`
	GoRoutines      int     `json:"goRoutines"`
	DirPath         string
}

var Hardware = new(HardwareInfo)

func init() {
	Hardware.DirPath = ""
}

// 检查磁盘空间是否已满
func (info *HardwareInfo) IsDiskFull(dir string) bool {
	Hardware.DirPath = dir
	if Hardware.DiskFreeSpace() < 10*1024*1024 {
		log.Logs.Error("Disk is full.")
		return true
	}

	return false
}

func (info *HardwareInfo) GetCpuPercent() (cpuInfo float64) {
	percent, err := cpu.Percent(0, false)
	if err == nil {
		cpuInfo = percent[0]
	}

	return
}

// 获取系统运行时的硬件状态
func (info *HardwareInfo) GetHardwareState() {
	info.CPU = info.GetCpuPercent()
	info.DiskUsedPercent = info.GetDiskUsedPercent()
	info.DiskFree = info.DiskFreeSpace()
	info.MemUsedPercent = info.RamUsed()
	info.GoMemory = info.CountGoUseRam()
	info.GoRoutines = info.CountGoroutine()
}

// 计算Golang运行时分配的内存总字节数
func (info *HardwareInfo) CountGoUseRam() int64 {
	ram := runtime.MemStats{}
	runtime.ReadMemStats(&ram)
	return int64(ram.Sys)
}

// 计算当前Goroutine的数量
func (info *HardwareInfo) CountGoroutine() int {
	return runtime.NumGoroutine()
}

// 获取CPU信息
func (info *HardwareInfo) CpuInfo() []cpu.InfoStat {
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Logs.Error("Get cpu info error: ", err)
	}

	return cpuInfo
}

// 获取CPU平均负载
func (info *HardwareInfo) CpuLoad() *load.AvgStat {
	avg, err := load.Avg()
	if err != nil {
		log.Logs.Error("Get cpu load avg error: ", err)
	}
	return avg
}

// 获取硬盘状态
func (info *HardwareInfo) DiskState() (diskFree uint64, diskUsed float64) {
	if info.DirPath == "" {
		info.DirPath = "/"
	}

	stat, err := disk.Usage(info.DirPath)
	if err == nil {
		diskFree = stat.Free
		diskUsed = stat.UsedPercent
		return
	}

	log.Logs.Error("Get hard disk stat error: ", err)
	return
}

// 获取硬盘已使用空间
func (info *HardwareInfo) DiskFreeSpace() (diskFree uint64) {
	if info.DirPath == "" {
		info.DirPath = "/"
	}

	stat, err := disk.Usage(info.DirPath)
	if err == nil {
		diskFree = stat.Free
		return
	}

	log.Logs.Error("Get hard disk free space error: ", err)
	return
}

// 获取硬盘使用百分比
func (info *HardwareInfo) GetDiskUsedPercent() (diskUsed float64) {
	if info.DirPath == "" {
		info.DirPath = "/"
	}

	stat, err := disk.Usage(info.DirPath)
	if err == nil {
		diskUsed = stat.UsedPercent
		return
	}

	log.Logs.Error("Get hard disk stat error: ", err)
	return
}

// 获取内存使用率
func (info *HardwareInfo) RamUsed() (used float64) {

	stat, err := mem.VirtualMemory()
	if err == nil {
		used = stat.UsedPercent
		return
	}

	log.Logs.Error("Get memory usage error: ", err)
	return
}
