package service

import (
	"crypto/tls"
	"hidden-record-assistant/backend/module/cmdx"
	"hidden-record-assistant/backend/module/errs"
	"hidden-record-assistant/backend/module/zlog"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type HTTPConnector struct {
	Client *http.Client
	Prefix string
	IsPing uint32
}

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

func (c *HTTPConnector) Init(pingCbfunc func()) (err error) {
	defer func() {
		if recover() != nil {
			zlog.Errorf("ping panic!")
			return
		}
		if err != nil {
			zlog.Errorf("network.Client init failed: %s", err.Error())
			return
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

	go c.Ping(pingCbfunc)
	return
}

func (c *HTTPConnector) Ping(cbs func()) {
	for {
		time.Sleep(time.Second * 1)
		zlog.Debugf("pinging...")
		_, err := cmdx.Exec("tasklist|findstr LeagueClientUx.exe")
		if err != nil {
			zlog.Error("LeagueClientUx.exe not found")
			cbs()
			return
		}
	}
}

func (c *HTTPConnector) getResponse(url string) (*http.Response, error) {
	return c.Client.Get(c.Prefix + url)
}

func (c *HTTPConnector) Get(url string) (data string, err error) {
	defer func() {
		if err != nil {
			zlog.Errorf("Get %s: %s", url, err.Error())
			return
		}
	}()
	resp, err := c.getResponse(url)
	if err != nil {
		return
	}
	binData, err := io.ReadAll(resp.Body)

	data = string(binData)

	zlog.Debugf("Get %s: %s", url, data)
	return
}
