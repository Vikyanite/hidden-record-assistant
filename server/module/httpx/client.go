package httpx

import (
	"hidden-record-assistant/server/module/zlog"
	"net/http"
)

type Client struct {
	*http.Client
	Prefix string
}

var ClientInstance *Client

func InitClient(Token, Port string) {
	ClientInstance = &Client{
		Client: http.DefaultClient,
		Prefix: "https://riot:" + Token + "@127.0.0.1:" + Port,
	}
	zlog.Debugf("httpx.Client init success, your prefixURL is: %s", ClientInstance.Prefix)
}

func Get(url string) (*http.Response, error) {
	url = ClientInstance.Prefix + url
	return ClientInstance.Get(url)
}

func Post(url string) (*http.Response, error) {
	url = ClientInstance.Prefix + url
	return ClientInstance.Post(url, "", nil)
}
