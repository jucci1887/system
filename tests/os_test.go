/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"github.com/jucci1887/system"
	"testing"
)

func TestSystemStart(t *testing.T) {
	Test.Start("System")
}

func TestGetCurrentDir(t *testing.T) {
	msg := "/home/kernel/project/go/system/tests"
	result := system.GetCurrentDir()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetCurrentFilename(t *testing.T) {
	msg := "/home/kernel/project/go/system/tests/t.test"
	result := system.GetCurrentFilename()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestCurrentAndAbsPath(t *testing.T) {
	msg := "/home/kernel/project/go/system/tests/t.test"
	result := system.CurrentAndAbsPath()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestSetCurrentPath(t *testing.T) {
	msg := "./t.test"
	result := system.SetCurrentPath()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetAbsPath(t *testing.T) {
	msg := "/home/kernel/project/go/system/tests/t.test"
	path := system.SetCurrentPath()
	result := system.GetAbsPath(path)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetConfigDir(t *testing.T) {
	msg := "/home/kernel/project/go/system/config"
	result := system.GetConfigDir()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetLastPath(t *testing.T) {
	msg := "/home/kernel/project/go/system"
	path := system.GetCurrentDir()
	result := system.GetLastPath(path)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetRootPath(t *testing.T) {
	msg := "/home/kernel/project/go/system"
	result := system.GetRootPath()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestParseEnvVar(t *testing.T) {
	msg := "/home/kernel"
	path := "${HOME}"
	result := system.ParseEnvVar(path)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestCallBrowser(t *testing.T) {
	url := "http://www.bing.com"
	msg := "http://www.bing.com"
	err := system.CallBrowser(url)

	if err != nil {
		Test.T(t).Logs(msg).No(err)
	} else {
		Test.T(t).Logs(msg).Ok(err)
	}
}

func TestGetOS(t *testing.T) {
	msg := "linux"
	result := system.GetOS()

	if result != msg {
		Test.T(t).Logs(msg).No(result)
	} else {
		Test.T(t).Logs(msg).Ok(result)
	}
}

func TestGetFilepath(t *testing.T) {
	msg := "/home/kernel/project/go/system/config/config.toml"
	result := system.GetFilepath("config", "config.toml")

	if result != msg {
		Test.T(t).Logs(msg).No(result)
	} else {
		Test.T(t).Logs(msg).Ok(result)
	}
}

func TestSystemEnd(t *testing.T) {
	Test.End("System")
}
