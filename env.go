/**
 * Created by Kernel.Huang
 * User: kernelman@live.com
 * Date: 2021/3/23
 * Time: 16:03
 */

package system

import (
	"log"
	"os"
)

type env struct {
	name string
}

var Env = new(env)

func (e *env) New(envName string) *env {
	e.name = os.Getenv(envName)
	if e.name == "" {
		log.Println("Env name not found.")
	}

	return e
}

func (e *env) Toml() string {
	switch e.name {
	case "dev":
		return "dev.toml"
	case "prd":
		return "prd.toml"
	case "test":
		return "test.toml"
	default:
		return "dev.toml"
	}
}

func (e *env) Ini() string {
	switch e.name {
	case "dev":
		return "dev.ini"
	case "prd":
		return "prd.ini"
	case "test":
		return "test.ini"
	default:
		return "dev.ini"
	}
}

func (e *env) Yaml() string {
	switch e.name {
	case "dev":
		return "dev.yaml"
	case "prd":
		return "prd.yaml"
	case "test":
		return "test.yaml"
	default:
		return "dev.yaml"
	}
}

func (e *env) Json() string {
	switch e.name {
	case "dev":
		return "dev.json"
	case "prd":
		return "prd.json"
	case "test":
		return "test.json"
	default:
		return "dev.json"
	}
}
