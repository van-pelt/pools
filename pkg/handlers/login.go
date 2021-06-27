package handlers

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

func Login(ctx context.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := GetObj(ctx)
		//TODO add ok
		t, err := template.ParseFiles(data.AllowDir + "login.html")
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
