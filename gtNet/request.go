package gtNet

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/*
* POST请求
* @headers http请求头;可以使用NewJsonHeader()生成json请求头;
* @client 允许传入自定义的Client对象;传入nil，每次请求将会临时生成默认Client对象;
**/
func Post(url string, headers map[string]string, bodyByte []byte, client *http.Client) ([]byte, error) {
	return request(url, headers, bodyByte, http.MethodPost, client)
}

/*
* GET请求
* @headers http请求头;可以使用NewJsonHeader()生成json请求头;
* @client 允许传入自定义的Client对象;传入nil，每次请求将会临时生成默认Client对象;
**/
func Get(url string, headers map[string]string, client *http.Client) ([]byte, error) {
	return request(url, headers, []byte(""), http.MethodGet, client)
}

func NewJsonHeader() map[string]string {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	return header
}

func request(url string, headers map[string]string, bodyByte []byte, method string, client *http.Client, ) ([]byte, error) {
	var err error

	var req *http.Request
	req, err = http.NewRequest(method, url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	if client == nil {
		client = &http.Client{}
	}
	var rsp *http.Response
	rsp, err = client.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var rspBody []byte
	rspBody, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	return rspBody, nil
}
