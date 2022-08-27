package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kavanahuang/api"
	"github.com/kavanahuang/log"
	"os"
	"time"
)

type terminal struct{}

var Terminal = new(terminal)

func (t *terminal) Call(exitCallback, execCallback func()) {
	_, _ = fmt.Fprint(os.Stdout, ">> ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Logs.Error("Scanner error: ", scanner.Err())
			return
		}

		text := scanner.Bytes()
		if bytes.Equal(text, []byte("exit")) {
			exitCallback()
			break
		}

		execCallback()
		_, _ = fmt.Fprint(os.Stdout, ">> ")
	}
}
