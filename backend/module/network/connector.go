package network

import (
	"crypto/tls"
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
	// 因为默认的http.Client会验证证书，所以我们需要创建一个不验证证书的http.Client
	// 否则会报错：x509: certificate signed by unknown authority
	// 但是这样会导致安全问题，还是需要注意这个隐患
	connector = &HTTPConnector{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
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
	defer func() {
		if err != nil {
			zlog.Errorf("network.Client init failed: %s", err.Error())
		}
	}()

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
	Token, Port := flagMap["remoting-auth-token"], flagMap["app-port"]
	c.Prefix = "https://riot:" + Token + "@127.0.0.1:" + Port
	zlog.Debugf("network.Client init success, your prefixURL is: %s", c.Prefix)
	return
}

func (c *HTTPConnector) Get(url string) (*http.Response, error) {
	return c.Client.Get(c.Prefix + url)
}
