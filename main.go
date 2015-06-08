package main

import (
	_ "fmt"
	"log"
	"net/http"
	"time"
	"github.com/shengzhi/gowechat/wxapi"
	"os"
)

func main() {
	appid := os.Getenv("GOWECHAT_APP_ID")
	secret:= os.Getenv("GOWECHAT_SECRET")
	wxapi.RunTokenServer(appid, secret)
	log.Println("wechat server: start!")
	log.Printf("appid:%s,secret:%s\r\n",appid,secret)
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
	res.Content = "消息已收到，正在考虑怎么回复..."
	res.MsgType = wxapi.MsgTypeText
	return &res
}

