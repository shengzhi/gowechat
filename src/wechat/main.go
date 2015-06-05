package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wechat/wxapi"
)

const (
	appid  = "wx5ba7d468fb057942"
	secret = "dcf52903fd32e8908ae6cccb4a0b1c4b"
)

func main() {
	wxapi.RunTokenServer(appid, secret)

	go func() {
		for {
			if token,err := wxapi.GetToken(false);err == nil{
				fmt.Printf("Token: %s\r\n", token)
			}else{
				log.Println(err)
			}
			time.Sleep(time.Duration(10) * time.Second)
		}
	}()
	log.Println("wechat server: start!")
	wxHandler := wxapi.NewHandler()
	wxHandler.DefaultHandler = wxapi.MsgHandlerFunc(defaultMsgHandler)
	//wxHandler.Register(wxapi.MsgTypeText, textMsgHandler)
	http.Handle("/", wxHandler)
	err := http.ListenAndServe(":8080", nil)
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
