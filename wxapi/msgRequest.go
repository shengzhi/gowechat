package wxapi

import(
	_ "fmt"
	"encoding/xml"
	"strings"
)

const(
	MsgTypeText = "text"
	MsgTypeImage = "image"
	MsgTypeVoice = "voice"
	MsgTypeVideo = "video"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLink = "link"
	MsgTypeLocation = "location"
	MsgTypeEvent = "event"
)

type MsgRequestor interface{
	GetMsgType() string
}

type BaseMsgRequest struct{
	ToUserName string
	FromUserName string
	CreateTime int64
	MsgType string
	MsgId int64
}

func(r BaseMsgRequest) GetMsgType() string{
	return strings.ToLower(r.MsgType)
}

type TextMsgRequest struct{
	Content string
}

type MultiMediaMsgRequest struct{
	MediaId string
}

type ImageMsgRequest struct{
	PicUrl string
}

type VoiceMsgRequest struct{
	Format string
}

type VideoMsgRequest struct{
	ThumbMediaId string
}

type LocationMsgRequest struct{
	Location_X, Location_Y float64
	Scale int
	Label string
}

type LinkMsgRequest struct{
	Title, Description,Url string
}

type WXMsgRequest struct{
	XMLName xml.Name `xml:"xml"`
	BaseMsgRequest
	MultiMediaMsgRequest
	TextMsgRequest
	VoiceMsgRequest
	ImageMsgRequest
	VideoMsgRequest
	LocationMsgRequest
	LinkMsgRequest
}