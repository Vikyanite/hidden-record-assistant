package backend

import (
	"hidden-record-assistant/backend/module/cmdx"
	"testing"
)

func Test(t *testing.T) {
	exec, err := cmdx.Exec("tasklist|findstr LeagueClientUx.exe")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(exec))
}
