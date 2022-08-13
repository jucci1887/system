/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

type BaseTests struct {
	t       *testing.T
	file    string
	line    string
	content string
}

var test = "Tested "
var tested = " ---> "
var ok = "==> ok: "
var no = "<== no: "
var skip = "=== skip: "
var debug = "Debug: "
var Test = BaseTests{}

func (b *BaseTests) T(t *testing.T) *BaseTests {
	b.t = t
	return b
}

func (b *BaseTests) Logs(content string) *BaseTests {
	if content == "" {
		b.t.Error("Content required")
	}

	b.content = content
	return b
}

func (b *BaseTests) Ok(out interface{}) {
	log.Printf("%s\033[0;40;32m%s\033[0m\n", test+tested+b.content+" "+ok, out)
}

func (b *BaseTests) No(out interface{}) {
	_, file, line, _ := runtime.Caller(1)
	b.t.Errorf("%s\033[0;40;31m%s\033[0m\n", test+filepath.Base(file)+":"+strconv.Itoa(line)+tested+b.content+" "+no, out)
}

func (b *BaseTests) Debug(out interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s\033[0;40;34m%s\033[0m\n", debug+filepath.Base(file)+":"+strconv.Itoa(line)+" ", out)
}

func (b *BaseTests) Skip(out interface{}) {
	b.t.Fatal(test+b.content+" "+skip, out)
}

func (b *BaseTests) Start(content string) {
	log.Println("Testing: " + content + " start: **********************************************************************")
	log.Println("")
	log.Println("")
}

func (b *BaseTests) End(content string) {
	log.Println("")
	log.Println("")
	log.Println("Testing: " + content + " End: **********************************************************************")
	log.Println("")
	log.Println("")
}
