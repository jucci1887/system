package system

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kavanahuang/log"
	"os"
	"time"
)

type terminal struct{}

var Terminal = new(terminal)

func (t *terminal) Call(exitCallback func(), execCallback func(t []byte) bool) {
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

		if !execCallback(text) {
			break
		}
		_, _ = fmt.Fprint(os.Stdout, ">> ")
	}
}
