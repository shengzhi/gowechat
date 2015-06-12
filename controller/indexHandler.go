package controller

import(
	"net/http"
	"html/template"
	_ "fmt"
	"github.com/shengzhi/gowechat/config"
)

var(
	layoutTmp = 
	templates = template.Must(template.ParseFiles(config.View_Base_Folder+`/home/index.html`))
)
type HomeHandler struct{
	
}

type homeViewModel struct{
	StaticWebSite string
	Title string
}

func(this *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	templates.Execute(w,&homeViewModel{
	"http://localhost:8080","菜单创建"})
}