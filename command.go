/**
 * Created by Kernel.Huang
 * User: kernelman@live.com
 * Date: 2021/3/23
 * Time: 16:03
 */

package system

import "os"

type CommandServices struct {
	ExecDir  string
	ExecName string
	PidDir   string
	PidName  string
}

var Command = new(CommandServices)

func (cs *CommandServices) Run(execDir, execName, pidDir, pidName string) {
	cs.ExecDir = execDir
	cs.ExecName = execName
	cs.PidDir = pidDir
	cs.PidName = pidName

	switch len(os.Args) {
	case 2:
		cs.isRun(os.Args)
	default:
		cs.usage()
	}
}

// 如果是合法指令则执行程序
func (cs *CommandServices) isRun(args []string) {
	Daemon.ExecDir = cs.ExecDir
	Daemon.ExecName = cs.ExecName
	Daemon.PidDir = cs.PidDir
	Daemon.PidName = cs.PidName

	command := args[1]
	switch command {
	case "start":
		_ = Daemon.Start()
		os.Exit(0)
	case "stop":
		_ = Daemon.Stop()
		os.Exit(0)
	case "restart":
		_ = Daemon.Restart()
		os.Exit(0)
	default:
		cs.usage()
	}
}

func (cs *CommandServices) usage() {
	Daemon.Usage()
	os.Exit(1)
}
