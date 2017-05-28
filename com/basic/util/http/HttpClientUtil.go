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

func httpGet(httpUrl string) {
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

func HttpDoi(m1 map[string]string, httpUrl string) string {
	client := &http.Client{}

	var request = url.Values{}
	for k, v := range m1 {
		request.Add(k, v)
	}

	data := request.Encode()

	req, err := http.NewRequest("POST", httpUrl, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Cookie", "name=anny")

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