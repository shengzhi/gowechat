package wxapi

import (
	"testing"
	"encoding/xml"
)

func TestTextMsgRequest(t *testing.T){
	data := `<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName> 
 <CreateTime>1348831860</CreateTime>
 <MsgType><![CDATA[text]]></MsgType>
 <Content><![CDATA[this is a test]]></Content>
 <MsgId>1234567890123456</MsgId>
 </xml>`
var request WXMsgRequest
	xml.Unmarshal([]byte(data),&request)
	if request.GetMsgType() != MsgTypeText{
		t.Errorf("Msg type: %s",request.GetMsgType())
		t.Errorf("request:%+v",request)
	} else{
		t.Logf("context:%s",request.Content)
	}	
}