package menu

import(
	_ "encoding/json"
	"gowechat/util"
	"gowechat/wxapi"
	"fmt"
	"errors"
)

const(
	ButtonLimit = 3 	//一级菜单三个
	SubButtonLimit = 5 	//二级菜单5个
)

const(
	ButtonNameLimit = 16	//菜单标题，不超过16个字节
	SubButtonNameLimit = 40 //子菜单不超过40个字节
)

const(
	retCode_Success = 0
)

const(
	url_createMenu=" https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
)

const(
	BtnTypeClick = "click"	//点击推事件
	BtnTypeView = "view"	//跳转URL
	BtnTypeScanPush = "scancode_push"		//扫码推事件
	BtnTypeScanWait = "scancode_waitmsg" 	//扫码推事件且弹出“消息接收中”提示框
	BtnTypePhoto = "pic_sysphoto"			//弹出系统拍照发图
	BtnTypePhotoOrAlbum="pic_photo_or_album"//弹出拍照或者相册发图
	BtnTypePicWX = "pic_weixin"			//弹出微信相册发图器
	BtnTypeLocation = "location_select" //弹出地理位置选择器
	BtnTypeMediaId = "media_id"			//下发消息（除文本消息）
	BtnTypeViewLimit = "view_limited"	//跳转图文消息URL
)

type Menu struct{
	Buttons []Button 	`json:"button,omitempty"`
}

type Button struct{
	Type       string   `json:"type,omitempty"`       // 非必须; 菜单的响应动作类型
	Name       string   `json:"name,omitempty"`       // 必须;  菜单标题，不超过16个字节，子菜单不超过40个字节
	Key        string   `json:"key,omitempty"`        // 非必须; 菜单KEY值，用于消息接口推送，不超过128字节
	URL        string   `json:"url,omitempty"`        // 非必须; 网页链接，用户点击菜单可打开链接，不超过256字节
	MediaId    string   `json:"media_id,omitempty"`   // 非必须; 调用新增永久素材接口返回的合法media_id
	SubButtons []Button `json:"sub_button,omitempty"` // 非必须; 二级菜单数组，个数应为1~5个
}

type Result struct{
	Code int32 `json:"errcode,omitempty"`
	Message string `json:"errmsg,omitempty"`
}

func Create(m *Menu) error{
	token,err := wxapi.GetToken(false)
	if err != nil{
		return err
	}
	url := fmt.Sprintf(url_createMenu,token)
	var result Result
	if err = util.PostJSON(url,m,&result);err != nil{return nil}
	if result.Code != retCode_Success{
		return errors.New("wechat server return error:"+result.Message)
	}
	return nil
}