package ding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var token = ""
var mu sync.Mutex

//授权接口获取token
func authorize() string {
	log.SetPrefix("authorize ")
	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", appKey, appSecret)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return ""
	}
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	d := authorizeResp{}
	err = json.Unmarshal(bs, &d)
	log.Println(string(bs))
	if err != nil {
		log.Println(err)
		return ""
	}
	if d.ErrCode != 0 {
		log.Println(err)
		return ""
	}
	log.Printf("success [token:%s]\n", token)
	return d.AccessToken
}

//钉钉异步通知
func AsyncSend(contentType, template string) {
	mu.Lock()
	defer mu.Unlock()
	for {
		asr := syncSendHandler(contentType, template)
		if asr == nil {
			break
		}
		switch asr.ErrCode {
		case 88:
			//token失效
		case 0:
			//success
			log.Printf("success [message:\n%s]\n", template)
			return
		default:
			//success
			log.Printf("fail [errmsg:%s]\n", asr.ErrMsg)
			return
		}
		time.Sleep(time.Second)
	}
}

func syncSendHandler(contentType, template string) *asyncSendResp {
	log.SetPrefix("AsyncSend ")
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", authorize())
	res, err := http.Post(url, contentType, bytes.NewBufferString(template))
	if err != nil {
		log.Println(err)
		return nil
	}
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	d := asyncSendResp{}
	err = json.Unmarshal(bs, &d)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(string(bs))
	return &d
}
