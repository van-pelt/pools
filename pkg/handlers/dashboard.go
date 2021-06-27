package handlers

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/van-pelt/pools/pkg/cache"
	"github.com/van-pelt/pools/pkg/database"
	"html/template"
	"log"
	"net/http"
)

func Dashboard(ctx context.Context, storage *database.Storage, cache *cache.Cache) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		auth := r.Header.Get("Authorization")
		fmt.Println("AUTH", auth)
		cookie, err := r.Cookie("bauth")
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		fmt.Println("bauth C=", cookie.Value)
		chl := cache.Get(cookie.Value)
		r.Header.Set("Authorization", "Basic "+cookie.Value)
		if AuthB64(r, chl.Email, chl.Pass, cookie.Value) {
			data := GetObj(ctx)
			//TODO add ok
			t, err := template.ParseFiles(data.AllowDir + "dashboard.html")
			if err != nil {
				log.Print(err.Error())
				return
			}
			err = t.Execute(w, data)
			if err != nil {
				log.Print(err.Error())
			}
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
