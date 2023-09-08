package support

import (
	"crypto/tls"
	"hidden-record-assistant/backend/model"
	"hidden-record-assistant/backend/module/cmdx"
	"hidden-record-assistant/backend/module/errs"
	"hidden-record-assistant/backend/module/zlog"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
)

type Connector struct {
	Client *http.Client
	Prefix string
	IsPing uint32
}

func NewConnector() (connector *Connector) {
	// 因为默认的http.Client会验证证书，所以我们需要创建一个不验证证书的http.Client
	// 否则会报错：x509: certificate signed by unknown authority
	// 但是这样会导致安全问题，还是需要注意这个隐患
	connector = &Connector{
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

func (c *Connector) Init(pingCbfunc func()) (data model.Auth, err error) {
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

	data = model.Auth{
		Token: Token,
		Port:  Port,
	}

	zlog.Debugf("network.Client init success, your prefixURL is: %s", c.Prefix)

	go c.Ping(pingCbfunc)

	return
}

func (c *Connector) Ping(cbs func()) {
	if atomic.LoadUint32(&c.IsPing) == 1 {
		return
	}
	atomic.StoreUint32(&c.IsPing, 1)
	defer func() {
		atomic.StoreUint32(&c.IsPing, 0)
	}()
	for {
		time.Sleep(time.Second * 1)
		_, err := cmdx.Exec("tasklist|findstr LeagueClientUx.exe")
		if err != nil {
			zlog.Error("LeagueClientUx.exe not found")
			cbs()
			return
		}
	}
}

func (c *Connector) getResponse(url string) (*http.Response, error) {
	return c.Client.Get(c.Prefix + url)
}

func (c *Connector) Get(url string) (data []byte, err error) {
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
	data, err = io.ReadAll(resp.Body)
	//if err == nil {
	//	zlog.Debugf("Get [%s] succeed!", url)
	//}
	return
}

func ExecuteTaskConcurrently(tasks ...func() error) (errChan chan error, num int) {
	num = len(tasks)
	errChan = make(chan error, num)
	for i := range tasks {
		go func(index int) {
			errChan <- tasks[index]()
		}(i)
	}
	return
}
