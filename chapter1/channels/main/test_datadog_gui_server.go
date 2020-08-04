package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//http请求
func httpHandle(method, urlVal, data string) {
	client := &http.Client{}
	var req *http.Request

	if data == "" {
		urlArr := strings.Split(urlVal, "?")
		if len(urlArr) == 2 {
			urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	} else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}

	// 可以添加多个cookie
	cookie1 := &http.Cookie{Name: "authToken", Value: "6bf510bc5559d86201946513b545d9c18193ed7ee40ff7cc2251a65836d9cf26", HttpOnly: true}
	req.AddCookie(cookie1)

	//添加header, 这里的配置是为了通过gui.go中的authorizePOST验证
	req.Header.Add("Authorization", "Bearer 6bf510bc5559d86201946513b545d9c18193ed7ee40ff7cc2251a65836d9cf26")
	//req.Header.Add("type", "heartbeat")
	//req.Header.Add("url", "http://192.168.5.211:8889/data")
	//req.Header.Add("module_path", "/data/br/base")
	//req.Header.Add("module_name", "druid")

	req.Header.Add("version", "V1")
	req.Header.Add("type", "RebalanceService")
	req.Header.Add("service", "http://192.168.150.71:8889/view")
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("error: %s", err)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s \n", string(b))
	fmt.Printf("response StatusCode: %d \n", resp.StatusCode)
	fmt.Printf("response cookies: %s \n", resp.Cookies())
	fmt.Printf("response header: %s \n", resp.Header)
}

//将get请求的参数进行转义
func getParseParam(param string) string {
	return url.PathEscape(param)
}

//测试
func main() {
	args := os.Args
	log.Printf("args:%s\n", args)
	//httpHandle("POST", "http://"+args[1]+":9876/authenticate", "")
	//httpHandle("POST", "http://"+args[1]+":9876/agent/uploadcfg", "")  // 配置文件
	httpHandle("POST", "http://"+args[1]+":9876/agent/heartbeat", "") // 心跳
	//httpHandle("POST", "http://"+args[1]+":8889/data", "")
}
