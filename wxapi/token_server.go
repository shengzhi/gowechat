package wxapi

import(
	"fmt"
	"time"
	"gowechat/util"
	"encoding/json"
	"log"
	"bytes"
	"errors"
)

const(
	grantType = "client_credential"
	UrlGetToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"
	defaultDuration = 7000*time.Second
)

var tokenServer *accTokenServer

type wxtoken struct{
	Token string `json:"access_token"`
	Duration int `json:"expires_in"`
}

type errMsg struct{
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

type accTokenServer struct{
	AppId,Secret string
	Token *wxtoken
}

func RunTokenServer(appid,secret string){
	if tokenServer== nil{
		tokenServer = &accTokenServer{AppId:appid,Secret:secret}
	}
	tokenServer.fetchToken()
	go func(s *accTokenServer){
		if s.Token != nil{
			c := time.Tick(time.Duration(s.Token.Duration-200)*time.Second)
			for _ = range c{
				s.fetchToken()
			}
		}
	}(tokenServer)
}

func (s *accTokenServer) fetchToken(){
	token,err := makeToken(s.AppId,s.Secret)
	if err != nil{
		log.Println("Get access token failed, error:",err)
		return
	}
	s.Token = token
}

func GetToken(forceRefresh bool) (string,error){
	if forceRefresh{tokenServer.fetchToken()}
	if tokenServer.Token != nil{
		return tokenServer.Token.Token,nil
	}
	return "",errors.New("No available access token")
}

func makeToken(appid,secret string)(*wxtoken,error){
	url := fmt.Sprintf(UrlGetToken,grantType,appid,secret)
	data,err := util.HttpGet(url)
	if err != nil{
		return nil,err
	}
	var success wxtoken
	var errRet errMsg
	if bytes.Contains(data,[]byte("access_token")){
		err = json.Unmarshal(data,&success)
		if err != nil{
			return nil,err
		}
		return &success,nil
	} else {
		json.Unmarshal(data,&errRet)
		return nil, errors.New(errRet.Errmsg)
	}
}