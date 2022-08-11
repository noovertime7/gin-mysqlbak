package core

import (
	"bytes"
	"encoding/json"
	"github.com/noovertime7/gin-mysqlbak/conf"
	"io/ioutil"
	"net/http"
)

type dingSendMessage struct {
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
	Message      string `json:"message"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}

var url = "http://" + conf.GetStringConf("dingProxyAgent", "addr")

func NewDingSender(token, secret, message string) *dingSendMessage {
	return &dingSendMessage{
		AccessToken:  token,
		AccessSecret: secret,
		Message:      message + "\n" + conf.GetStringConf("dingProxyAgent", "content"),
		Title:        conf.GetStringConf("dingProxyAgent", "title"),
	}
}

func (d *dingSendMessage) SendMessage() (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(d); err != nil {
		return "", err
	}
	res, err := http.Post(url+"/ding/sendmsg", "application/json", buf)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	dingmsg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(dingmsg), err
}

func (d *dingSendMessage) SendMarkdown() (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(d); err != nil {
		return "", err
	}
	res, err := http.Post(url+"/ding/sendmd", "application/json", buf)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	dingmsg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(dingmsg), err
}
