package handlers

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/van-pelt/pools/pkg/router/ctxparam"
	"html/template"
	"log"
	"net/http"
)

/*func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	basePath, err := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return
	}
	t, _ := template.ParseFiles(basePath + "/web/register.html")
	data := TodoPageData{
		Title: basePath,
	}
	t.Execute(w, data)
}*/

func Index(ctx context.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := GetObj(ctx)
		//TODO add ok
		fmt.Println(data.AllowDir + "register.html")
		t, err := template.ParseFiles(data.AllowDir + "register.html")
		if err != nil {
			log.Print(err.Error())
			return
		}
		err = t.Execute(w, data)
		if err != nil {
			log.Print(err.Error())
			return
		}
	}
}

type PageData struct {
	BasePath    string
	AllowDir    string
	BaseUrl     string
	UserSession string
}

func GetObj(ctx context.Context) PageData {
	return PageData{
		BasePath:    ctx.Value("params").(ctxparam.CtxParams).Get("BasePath"),
		AllowDir:    ctx.Value("params").(ctxparam.CtxParams).Get("AllowDir"),
		BaseUrl:     ctx.Value("params").(ctxparam.CtxParams).Get("BaseUrl"),
		UserSession: ctx.Value("params").(ctxparam.CtxParams).Get("UserSession"),
	}
}

func SetObj(ctx context.Context, userSessID string) {
	ctx.Value("params").(ctxparam.CtxParams).Val["UserSession"] = userSessID
}
