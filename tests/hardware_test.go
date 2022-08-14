/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"github.com/kavanahuang/system"
	"testing"
)

func TestStringStart(t *testing.T) {
	Test.Start("Hardware")
}

func TestHardware(t *testing.T) {
	cpu := system.Hardware.CpuInfo()
	avg := system.Hardware.CpuLoad()
	ram := system.Hardware.RamUsed()
	free, used := system.Hardware.DiskState()

	Test.T(t).Logs("Cpu info: ").Ok(cpu)
	Test.T(t).Logs("Cpu load avg: ").Ok(avg.String())
	Test.T(t).Logs("Hard disk free space").Ok(free)
	Test.T(t).Logs("Hard disk used: ").Ok(used)
	Test.T(t).Logs("Ram used: ").Ok(ram)
}

func TestStringEnd(t *testing.T) {
	Test.End("Hardware")
}
