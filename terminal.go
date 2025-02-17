/**
 * Created by Kernel.Huang
 * User: kernelman@live.com
 * Date: 2021/3/23
 * Time: 16:03
 */

package system

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jucci1887/logs"
	"os"
)

type terminal struct{}

var Terminal = new(terminal)

func (t *terminal) Call(exitCallback func(), execCallback func(t []byte) bool) {
	_, _ = fmt.Fprint(os.Stdout, ">> ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			logs.Error("Scanner error: ", scanner.Err())
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
