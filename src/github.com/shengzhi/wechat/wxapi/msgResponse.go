package wxapi

import "encoding/xml"

type BaseMsgResponse struct{
	XMLName xml.Name `xml:"xml"`
	ToUserName string
	FromUserName string
	CreateTime int64
	MsgType string
}

type TextMsgResponse struct{
	BaseMsgResponse
	Content string
}

type ImageMsgResponse struct{
	BaseMsgResponse
	MediaId string `xml:"Image>MediaId"`
}

type VoiceMsgResponse struct{
	BaseMsgResponse
	MediaId string `xml:"Voice>MediaId"`
}

type VideoMsgResponse struct{
	BaseMsgResponse
	MediaId string `xml:"Video>MediaId"`
	Title string `xml:"Video>Title"`
	Description string `xml:"Video>Description"`
}

type MusicMsgResponse struct{
	BaseMsgResponse
	Title string `xml:"Music>Title"`
	Description string `xml:"Music>Description"`
	MusicURL string `xml:"Music>MusicURL"`
	HQMusicUrl string `xml:"Music>HQMusicUrl"`
	ThumbMediaId string `xml:"Music>ThumbMediaId"`
}

type Article struct{
	Title string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl string `xml:"PicUrl"`
	Url string `xml:"Url"`
}

type NewsMsgResponse struct{
	BaseMsgResponse
	ArticleCount int
	Articles []Article `xml:"Articles>item"`
}