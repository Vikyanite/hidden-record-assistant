package cmdx

import (
	"os/exec"
)

func Exec(name string, arg ...string) ([]byte, error) {
	arg = append([]string{"/c", name}, arg...)
	cmd := exec.Command("cmd", arg...)
	return cmd.CombinedOutput()
}
