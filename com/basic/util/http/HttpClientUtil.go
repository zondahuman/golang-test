package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
	"encoding/json"
	"bytes"
)

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

func HttpGet(httpUrl string) string {
	resp, err := http.Get(httpUrl)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body)
}

func httpPost(httpUrl string) {
	resp, err := http.Post(httpUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm(httpUrl string, request map[string]string) string {
	resp, err := http.PostForm(httpUrl,
		url.Values{})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var result = string(body)
	return result

}

func httpDo(httpUrl string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func HttpPostFormByCookie(m1 map[string]string, httpUrl string, httpLoginUrl string) string {
	jar := new(Jar)
	client := &http.Client{nil, nil, jar, 99999999999992}

	//loginResponse, err1 := client.PostForm(httpLoginUrl, nil)
	loginRequest, _ := http.NewRequest("GET", httpLoginUrl, nil)
	//loginResponse, err1 := http.Get(httpLoginUrl)
	loginResponse, err1 := client.Do(loginRequest)
	if err1 != nil {
		panic(err1.Error())
	}else{
		body1, err1 := ioutil.ReadAll(loginResponse.Body)
		if err1 != nil {
			panic("no value for $USER..............")
		}
		var result1 = string(body1)
		fmt.Println("result1=", result1)
	}

	var request = url.Values{}
	for k, v := range m1 {
		request.Add(k, v)
	}

	data := request.Encode()

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("no value for $USER")
	}
	fmt.Println(json.Marshal(body))
	var result = string(body)
	return result
}

func HttpPostJsonByCookie(m1 map[string]string, httpUrl string, httpLoginUrl string) string {
	jar := new(Jar)
	client := &http.Client{nil, nil, jar, 99999999999992}

	//loginResponse, err1 := client.PostForm(httpLoginUrl, nil)
	loginRequest, _ := http.NewRequest("GET", httpLoginUrl, nil)
	//loginResponse, err1 := http.Get(httpLoginUrl)
	loginResponse, err1 := client.Do(loginRequest)
	if err1 != nil {
		panic(err1.Error())
	}else{
		body1, err1 := ioutil.ReadAll(loginResponse.Body)
		if err1 != nil {
			panic("no value for $USER..............")
		}
		var result1 = string(body1)
		fmt.Println("result1=", result1)
	}

	var request = url.Values{}
	for k, v := range m1 {
		request.Add(k, v)
	}

	data := request.Encode()

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("no value for $USER")
	}
	fmt.Println(json.Marshal(body))
	var result = string(body)
	return result
}

func HttpPost(params map[string]string, headers map[string]string, httpUrl string) string {
	client := &http.Client{}

	var request = url.Values{}
	for k, v := range params {
		request.Add(k, v)
	}

	data := request.Encode()

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("no value for $USER")
	}
	fmt.Println(json.Marshal(body))
	var result = string(body)
	return result
}

func HttpPostJson(json string, headers map[string]string, httpUrl string) string {
	req, err := http.NewRequest("POST", httpUrl, bytes.NewBuffer([]byte(json)))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	var result = string(body)
	return result
}