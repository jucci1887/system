/**
 * Created by Kernel.Huang
 * User: kernelman@live.com
 * Date: 2021/3/23
 * Time: 16:03
 */

package system

import (
	"errors"
	"fmt"
	"github.com/jucci1887/common"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type DaemonServices struct {
	pid      int
	command  *exec.Cmd
	ExecDir  string
	ExecName string
	PidDir   string
	PidName  string
}

var Daemon = new(DaemonServices)

// Version Set by release.
var Version = "0.0.1"

func (d *DaemonServices) New(execDir, execName, pidDir, pidName string) *DaemonServices {
	d.ExecDir = execDir
	d.ExecName = execName
	d.PidDir = pidDir
	d.PidName = pidName
	return d
}

func (d *DaemonServices) GetPidPath() string {
	if d.PidDir != "" && d.PidName != "" {
		return GetFilepath(d.PidDir, d.PidName)
	}
	return ""
}

func (d *DaemonServices) GetProgramPath() string {
	if d.ExecDir != "" && d.ExecName != "" {
		return GetFilepath(d.ExecDir, d.ExecName)
	}
	return ""
}

func (d *DaemonServices) Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage:\n")
	_, _ = fmt.Fprintf(os.Stderr, "       %s <command>\n", os.Args[0])
	_, _ = fmt.Fprintf(os.Stderr, "\n")
	_, _ = fmt.Fprintf(os.Stderr, "The commands are:\n")

	for _, v := range d.msg() {
		_, _ = fmt.Fprintf(os.Stderr, v)
	}

	_, _ = fmt.Fprintf(os.Stderr, "\n")
	_, _ = fmt.Fprintf(os.Stderr, "Version:\n")
	_, _ = fmt.Fprintf(os.Stderr, "       %s\n", Version)
	_, _ = fmt.Fprintf(os.Stderr, "\n")
}

func (d *DaemonServices) msg() map[string]string {
	return map[string]string{
		"start":   "       start	start service\n",
		"stop":    "       stop	stop service\n",
		"restart": "       restart	restart service\n",
	}
}

func (d *DaemonServices) Restart() bool {
	stop := d.Stop()
	if stop {
		time.Sleep(common.Time.IntToSecond(3))
		return d.Start()
	}

	return false
}

func (d *DaemonServices) Start() bool {
	return d.run()
}

func (d *DaemonServices) Stop() bool {
	pid := d.getPid()
	if pid != 0 {
		return d.isStop(pid)
	}

	log.Println("Service stop success")
	return true
}

// 判断是否需要停止程序
func (d *DaemonServices) isStop(pid int) bool {
	if d.pidIsAlive(pid) {
		err := syscall.Kill(pid, 9)
		if err != nil {
			log.Println("Service stop failed")
			return false
		}
	}

	log.Println("Service stop success")
	return true
}

// 执行程序
func (d *DaemonServices) run() bool {
	pid := d.getPid()
	if pid != 0 {
		if d.pidIsAlive(pid) {
			log.Println("Service start success")
			return true
		}
	}

	program := d.GetProgramPath()
	d.command = exec.Command(program)
	d.command.Stdout = os.Stdout

	err := d.command.Start()
	if err != nil {
		log.Println("Service start failed")
		return false

	} else {
		pid := d.command.Process.Pid
		d.isSetPid(pid)
	}

	return true
}

// 判断是否需要设置pid
func (d *DaemonServices) isSetPid(pid int) {
	if pid > 0 {
		err := d.setPid(pid)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Service start success")
		}
	}
}

func (d *DaemonServices) setPid(nowPid int) (err error) {
	common.Files.Perm = 0755
	pidPath := d.GetPidPath()
	open := common.Files.Overwrite(pidPath)
	defer func() { _ = open.Close() }()

	err = syscall.Flock(int(open.Fd()), syscall.LOCK_EX)
	if err != nil {
		return
	}

	pidString := common.Format.FromInt(nowPid).ToString()
	_, err = open.WriteString(pidString)

	if err != nil {
		err = errors.New("pid file write failed")
		return
	}

	return
}

func (d *DaemonServices) getPid() int {
	pidPath := d.GetPidPath()
	exists := common.Files.IsExists(pidPath)
	if exists == nil {
		pid := common.Files.GetFile(pidPath)
		return common.Format.FromByte(pid).ToInt()
	}

	return 0
}

// 检查pid是否存活
func (d *DaemonServices) pidIsAlive(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// 给进程发送Signal
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}

	return true
}
