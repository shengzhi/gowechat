package main

import (
	_ "fmt"
	"log"
	"net/http"
	"time"
	"github.com/shengzhi/gowechat/wxapi"
)

const (
	appid  = "wx5ba7d468fb057942"
	secret = "dcf52903fd32e8908ae6cccb4a0b1c4b"
)

func main() {
	wxapi.RunTokenServer(appid, secret)
	log.Println("wechat server: start!")
	wxHandler := wxapi.NewHandler()
	wxHandler.DefaultHandler = wxapi.MsgHandlerFunc(defaultMsgHandler)
	//wxHandler.Register(wxapi.MsgTypeText, textMsgHandler)
	http.Handle("/", wxHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("Wechat server launch failed, error:", err)
	}
	log.Println("Wechat service: stopped!")
}

func defaultMsgHandler(r *wxapi.WXMsgRequest) interface{} {
	res := wxapi.TextMsgResponse{}
	res.FromUserName = r.ToUserName
	res.ToUserName = r.FromUserName
	res.CreateTime = time.Now().Unix()
	res.Content = "这是一个悲伤的故事"
	res.MsgType = wxapi.MsgTypeText
	return &res
}

