package wxapi

import "testing"
import "encoding/xml"

func TestImageMsgResponse(t *testing.T){
	res := ImageMsgResponse{BaseMsgResponse:BaseMsgResponse{FromUserName: "Gong",ToUserName: "Gao",CreateTime: 12345678,MsgType:"image"},MediaId: "54321"}
	data,err := xml.Marshal(res)
	if err != nil{
		t.Error(err)
	} else{
		t.Logf("%+v",string(data))
	}
}

func TestNewsMsgResponse(t *testing.T){
	res := NewsMsgResponse{}
	res.FromUserName = "Gong"
	res.ToUserName="Gao"
	res.ArticleCount = 2
	res.CreateTime = 12345678
	res.MsgType = "news"
	res.Articles = []Article{{"title1","description1","PicUrl1","url1"},{"title2","description2","PicUrl2","url2"}}
	data,err := xml.Marshal(res)
	if err != nil{
		t.Error(err)
	} else{
		t.Logf("%+v",string(data))
	}
}