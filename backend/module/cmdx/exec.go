package cmdx

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func Exec(name string, arg ...string) (ret []byte, err error) {
	arg = append([]string{"/c", name}, arg...)
	cmd := exec.Command("cmd", arg...)
	//zlog.Debugf("cmd: %s", cmd.String())
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	ret = out.Bytes()
	return
}
