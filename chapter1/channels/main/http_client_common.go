package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func httpGet(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		// handle error
		return nil
	}

	defer response.Body.Close()

	return response
}

func httpPost(url string, requestBody string) *http.Response {
	response, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(requestBody))
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	return response
}

func httpPostForm(url string, values url.Values) *http.Response {
	response, err := http.PostForm(url, values)

	if err != nil {
		// handle error
		return nil
	}

	defer response.Body.Close()

	return response

}

func httpDo(method string, url string, cookie string, header map[string]string, requestBody string) *http.Response {
	client := &http.Client{}

	var requestBodyReader io.Reader
	if requestBody == "" {
		requestBodyReader = nil
	} else {
		requestBodyReader = strings.NewReader(requestBody)
	}

	req, err := http.NewRequest(method, url, requestBodyReader)
	if err != nil {
		// handle error
		log.Println("http do occur error when new request")
		return nil
	}

	// 添加cookie
	reqCookie := &http.Cookie{Name: "authToken", Value: cookie, HttpOnly: true}
	req.AddCookie(reqCookie)

	// 添加Header，这是必须的，目的是为了通过gui.go中的authorizePOST验证
	req.Header.Add("Authorization", "Bearer "+cookie)

	// 添加自己封装的header参数
	for key, value := range header {
		req.Header.Add(key, value)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	return response
}
