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
	"os"
	"testing"
)

func TestEnvStart(t *testing.T) {
	Test.Start("Env")
}

func TestTomlEnv(t *testing.T) {
	_ = os.Setenv("project", "dev")
	msg := "dev.toml"
	result := system.Env.New("project").Toml()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestYamlEnv(t *testing.T) {
	_ = os.Setenv("project", "test")
	msg := "test.yaml"
	result := system.Env.New("project").Yaml()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestIniEnv(t *testing.T) {
	_ = os.Setenv("project", "prd")
	msg := "prd.ini"
	result := system.Env.New("project").Ini()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestJsonEnv(t *testing.T) {
	_ = os.Setenv("project", "prd")
	msg := "prd.json"
	result := system.Env.New("project").Json()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestEnvEnd(t *testing.T) {
	Test.End("Env")
}
