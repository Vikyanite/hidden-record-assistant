package backend

import (
	"testing"
)

func Test(t *testing.T) {
	app := NewApp()
	err := app.InitBackend()
	binData, err := app.GetSummonerByName("remakeoui9")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(binData)

}
