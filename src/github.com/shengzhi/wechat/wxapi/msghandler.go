package wxapi

import(
	"net/http"
	"encoding/xml"
	"bytes"
	"fmt"
	"github.com/shengzhi/wechat/auth"
	"strings"
	"log"
	"github.com/shengzhi/wechat/util"
)

type MsgHandlerFunc func(*WXMsgRequest) interface{}

func (f MsgHandlerFunc)Process(m *WXMsgRequest) interface{}{
	return f(m)
}

type MsgHandler interface{
	Process(*WXMsgRequest) interface{}
}
//Define WXApiHandler
type WXApiHander struct{
	handlers map[string]MsgHandler
	DefaultHandler MsgHandler
}
//Create a new instance for WXApiHandler
func NewHandler() *WXApiHander{
	h := &WXApiHander{handlers:make(map[string]MsgHandler)}
	return h
}
//Get WX message handler as per specified message type
func (h *WXApiHander)GetHandler(msgType string)MsgHandler{
	if handler, ok := h.handlers[msgType]; ok{
		return handler
	}
	return h.DefaultHandler
}
//Reister a function as message handler
func(h *WXApiHander)Register(t string,f func(*WXMsgRequest) interface{}){
	if _,exist := h.handlers[t];!exist{
		h.handlers[t] = MsgHandlerFunc(f)
	}
}

//Implements ServeHTTP as http handler
func (h *WXApiHander)ServeHTTP(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	if isValid := urlValidate(r);!isValid{
		log.Println("Invalid request",r.RequestURI)
		http.Error(w,"Invalid request",http.StatusBadRequest)
		return
	}
	method := strings.ToLower(r.Method)
	if method == "get"{
		fmt.Fprint(w,strings.Join(r.Form["echostr"],""))
		return
	}
	
	if method == "post"{
		buf := util.ShortTextBufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		buf.ReadFrom(r.Body)
		defer r.Body.Close()
		defer util.ShortTextBufferPool.Put(buf)
		var wxRequest WXMsgRequest
		xml.Unmarshal(buf.Bytes(), &wxRequest)
		handler := h.GetHandler(wxRequest.MsgType)
		response := handler.Process(&wxRequest)
		data,_ := xml.MarshalIndent(response," ","  ")
		w.Header().Set("Content-Type","text/xml; charset=utf-8")
		fmt.Fprint(w,string(data))
		return
	}
}

func urlValidate(r *http.Request) bool{
	nonce := strings.Join(r.Form["nonce"],"")
	timestamp := strings.Join(r.Form["timestamp"],"")
	clientSignature := strings.Join(r.Form["signature"],"")
	return auth.ValidateUrl(nonce,timestamp,clientSignature)
}