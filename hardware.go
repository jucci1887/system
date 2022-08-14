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
)

type hardware struct {
	DirPath string
}

var Hardware = new(hardware)

func init() {
	Hardware.DirPath = ""
}

// 获取CPU信息
func (h *hardware) CpuInfo() []cpu.InfoStat {
	info, err := cpu.Info()
	if err != nil {
		log.Logs.Error("Get cpu info error: ", err)
	}

	return info
}

// 获取CPU平均负载
func (h *hardware) CpuLoad() *load.AvgStat {
	avg, err := load.Avg()
	if err != nil {
		log.Logs.Error("Get cpu load avg error: ", err)
	}
	return avg
}

// 获取硬盘状态
func (h *hardware) DiskState() (diskFree uint64, diskUsed float64) {
	if h.DirPath == "" {
		h.DirPath = "/"
	}

	stat, err := disk.Usage(h.DirPath)
	if err == nil {
		diskFree = stat.Free
		diskUsed = stat.UsedPercent
		return
	}

	log.Logs.Error("Get hard disk stat error: ", err)
	return
}

// 获取硬盘已使用空间
func (h *hardware) DiskFreeSpace() (diskFree uint64) {
	if h.DirPath == "" {
		h.DirPath = "/"
	}

	stat, err := disk.Usage(h.DirPath)
	if err == nil {
		diskFree = stat.Free
		return
	}

	log.Logs.Error("Get hard disk free space error: ", err)
	return
}

// 获取硬盘使用百分比
func (h *hardware) DiskUsedPercent() (diskUsed float64) {
	if h.DirPath == "" {
		h.DirPath = "/"
	}

	stat, err := disk.Usage(h.DirPath)
	if err == nil {
		diskUsed = stat.UsedPercent
		return
	}

	log.Logs.Error("Get hard disk stat error: ", err)
	return
}

// 获取内存使用率
func (h *hardware) RamUsed() (used float64) {

	stat, err := mem.VirtualMemory()
	if err == nil {
		used = stat.UsedPercent
		return
	}

	log.Logs.Error("Get memory usage error: ", err)
	return
}
