package backend

import (
	"hidden-record-assistant/backend/service/support"
	"testing"
)

func Test(t *testing.T) {
	conn := support.NewConnector()
	app := NewApp(conn)
	init, err := app.InitBackend()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(init)
}
