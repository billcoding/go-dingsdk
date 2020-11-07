package ding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var token = ""

//授权接口获取token
func authorize() {
	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", appKey, appSecret)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	d := authorizeResp{}
	err = json.Unmarshal(bs, &d)
	log.Println(string(bs))
	if err != nil {
		log.Fatal(err)
		return
	}
	if d.ErrCode != 0 {
		log.Fatal(err)
		return
	}
	token = d.AccessToken
	log.Printf("authorize success [token:%s]\n", token)
}

//钉钉异步通知
func AsyncSend(contentType, template string) {
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", token)
	res, err := http.Post(url, contentType, bytes.NewBufferString(template))
	if err != nil {
		log.Fatal(err)
		return
	}
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	d := asyncSendResp{}
	err = json.Unmarshal(bs, &d)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(bs))
	switch d.ErrCode {
	case 88:
		//token失效
		time.Sleep(time.Second)
		authorize()
		AsyncSend(contentType, template)
	case 0:
		//success
		log.Printf("asyncSend success [message:\n%s]\n", template)
	default:
		//success
		log.Fatalf("asyncSend fail [errmsg:%s]\n", d.ErrMsg)
	}
}
