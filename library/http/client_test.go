package http

import (
	"fmt"
	"testing"
)

func init() {
	httpClient = createHTTPClient()
}

func Test_HttpGet(t *testing.T) {
	s1, _ := GetHttp("http://baidu.com")
	fmt.Println(string(s1))
	s2, _ := GetHttp("http://baidu.com")
	fmt.Println(string(s2))
	s3, _ := GetHttp("http://baidu.com")
	fmt.Println(string(s3))
}

func Test_PostGet(t *testing.T) {
	s1, _ := PostHttp("http://baidu.com", nil)
	fmt.Println(string(s1))
	s2, _ := PostHttp("http://baidu.com", nil)
	fmt.Println(string(s2))
}
