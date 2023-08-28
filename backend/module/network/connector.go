package network

import (
	"hidden-record-assistant/backend/module/cmdx"
	"hidden-record-assistant/backend/module/errs"
	"hidden-record-assistant/backend/module/zlog"
	"net/http"
	"regexp"
	"strings"
)

type HTTPConnector struct {
	Client *http.Client
	Prefix string
}

var DefaultConnector = NewHTTPConnector()

func NewHTTPConnector() (connector *HTTPConnector) {
	connector = &HTTPConnector{
		Client: http.DefaultClient,
	}
	return
}

func InitDefaultConnector() error {
	return DefaultConnector.Init()
}

func flagsToMap(res []byte) map[string]string {
	// 使用正则表达式来提取键值对
	re := regexp.MustCompile(`"--([^=]+)=([^ ]+)"`)
	matches := re.FindAllStringSubmatch(string(res), -1)

	// 创建一个map来存储键值对
	configMap := make(map[string]string)

	// 将提取的键值对存储到map中
	for _, match := range matches {
		key := match[1]
		value := match[2]
		configMap[key] = value
	}

	//// 打印map中的键值对
	//for key, value := range configMap {
	//	fmt.Printf("%s: %s\n", key, value)
	//}
	return configMap
}

func (c *HTTPConnector) Init() (err error) {
	res, err := cmdx.Exec("wmic PROCESS WHERE name='LeagueClientUx.exe' GET commandline")
	if err != nil {
		err = errs.InternalError(err)
		return
	}
	// 通过判断是否有CommandLine来判断是否启动客户端
	var FoundClientExe = strings.HasPrefix(string(res), "CommandLine")

	flagMap := flagsToMap(res)
	if len(flagMap) == 0 {
		err = errs.ErrNeedAdmin
		if !FoundClientExe {
			err = errs.ErrCantFindClient
		}
		return
	}
	Token, Port := flagMap["remoting-auth-token"], flagMap["riotclient-app-port"]
	c.Prefix = "https://riot:" + Token + "@127.0.0.1:" + Port

	zlog.Debugf("network.Client init success, your prefixURL is: %s", c.Prefix)
	return
}

func (c *HTTPConnector) Get(url string) (*http.Response, error) {
	return c.Client.Get(c.Prefix + url)
}
