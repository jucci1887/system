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
	msg := "Get Hardware Info: "
	free, used := system.Hardware.DiskState()
	cpu := system.Hardware.CpuInfo()
	avg := system.Hardware.CpuLoad()

	Test.T(t).Logs(msg).Ok(cpu)
	Test.T(t).Logs(msg).Ok(avg.String())
	Test.T(t).Logs(msg).Ok(free)
	Test.T(t).Logs(msg).Ok(used)
}

func TestStringEnd(t *testing.T) {
	Test.End("Hardware")
}
