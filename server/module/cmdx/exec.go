package cmdx

import (
	"hidden-record-assistant/server/module/errs"
	"os/exec"
	"strings"
)

func Exec(name string, arg ...string) ([]byte, error) {
	arg = append([]string{"/c", name}, arg...)
	cmd := exec.Command("cmd", arg...)
	return cmd.CombinedOutput()
}

func ExecWmicToMap() (retMap map[string]string, err error) {
	res, err := Exec("wmic PROCESS WHERE name='LeagueClientUx.exe' GET commandline")
	if err != nil {
		err = errs.InternalError(err)
		return
	}

	var ClientIsStarted = strings.HasPrefix(string(res), "CommandLine")

	retMap = FlagsToMap(res)
	if len(retMap) == 0 {
		if ClientIsStarted {
			err = errs.ErrNeedAdmin
		} else {
			err = errs.ErrCantFindClient
		}
		return
	}
	return
}
