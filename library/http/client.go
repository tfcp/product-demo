package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"tfpro/library/log"
	"time"
)

var (
	httpClient    *http.Client
	defaultExpire = 5 * time.Minute
)

type ResHttp struct {
	ErrCode int `json:"errcode"`
	ErrMsg  int `json:"errmsg"`
}

func init() {
	httpClient = createHTTPClient()
}

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     time.Duration(90) * time.Second,
		},
	}
	return client
}

func GetHttp(queryurl string, timeout ...time.Duration) ([]byte, error) {
	resCache, err := gcache.Get(queryurl)
	if err != nil {
		log.Logger.Errorf("GetHttp.gcache.GetError:%v", err)
	}
	if resCache != nil {
		return resCache.([]byte), nil
	}
	u, _ := url.Parse(queryurl)
	retstr, err := httpClient.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("http.Get %v", err)
	}
	result, err := ioutil.ReadAll(retstr.Body)
	retstr.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %v", err)
	}
	expire := defaultExpire
	if len(timeout) > 0 {
		expire = timeout[0]
	}
	if err = gcache.Set(queryurl, result, expire); err != nil {
		log.Logger.Errorf("GetHttp.gcache.SetError:%v", err)
	}
	return result, nil
}

func PostHttp(queryurl string, postdata map[string]string, timeout ...time.Duration) ([]byte, error) {
	postKey := fmt.Sprintf("post#%s#%s", queryurl, gconv.String(postdata))
	resCache, err := gcache.Get(postKey)
	if err != nil {
		log.Logger.Errorf("PostHttp.gcache.GetError:%v", err)
	}
	if resCache != nil {
		return resCache.([]byte), nil
	}
	data, err := json.Marshal(postdata)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal %v", err)
	}

	body := bytes.NewBuffer([]byte(data))
	resp, err := httpClient.Post(queryurl, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Post %v", err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %v", err)
	}
	expire := defaultExpire
	if len(timeout) > 0 {
		expire = timeout[0]
	}
	if err = gcache.Set(postKey, result, expire); err != nil {
		log.Logger.Errorf("PostHttp.gcache.SetError:%v", err)
	}
	return result, nil
}
