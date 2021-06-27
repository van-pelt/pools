package handlers

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/van-pelt/pools/internal/user/service"
	"github.com/van-pelt/pools/pkg/cache"
	"github.com/van-pelt/pools/pkg/database"
	"net/http"
	"strings"
	"time"
)

const basicAuthPrefix string = "Basic "

type ResponseAuth struct {
	Status  string
	Code    int
	Message string
}

func Auth(ctx context.Context, storage *database.Storage, cache *cache.Cache) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		email := r.PostFormValue("email")
		//pass := r.PostFormValue("pass")
		userSrv := service.New(storage)
		fmt.Println(email)
		user, err := userSrv.CheckUserData(email)
		fmt.Println(user)
		if err != nil {
			p := ResponseAuth{
				Status:  "ERROR",
				Code:    200,
				Message: err.Error(),
			}
			data, err := json.Marshal(p)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			}
			fmt.Fprintf(w, string(data))

			return
		}
		if AuthB64(r, user.Email, user.Password, user.Hash) {
			fmt.Println("AUTH64")
			//sessID := uuid.New().String()
			cache.Add(user.Hash, user.FullName, user.Email, user.Password)
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "bauth", Value: user.Hash, Expires: expiration}
			http.SetCookie(w, &cookie)
			//SetObj(ctx, sessID)

			data, err := json.Marshal(ResponseAuth{
				Status:  "OK",
				Code:    200,
				Message: user.Hash,
			})
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			}

			/*r.Method = "POST"
			w.Header().Set("Authorization", "Basic "+user.Hash)
			w.Header().Set("WWW-Authenticate", "")
			http.Redirect(w, r, "http://"+user.Email+":"+user.Password+"@127.0.0.1:8080/protected", http.StatusSeeOther)*/
			//http.Redirect(w, r, "protected", http.StatusSeeOther)
			//http.Redirect(w, r, "protected", http.StatusFound)
			fmt.Fprintf(w, string(data))
			return
			//http.Redirect(w, r, "http://"+user.Email+":"+user.Password+"@127.0.0.1:8080/protected", http.StatusSeeOther)
		}

		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}

}

func AuthB64(r *http.Request, user, pass, hash string) bool {
	auth := r.Header.Get("Authorization")
	fmt.Println("AUTH", auth)
	if strings.HasPrefix(auth, basicAuthPrefix) {
		payload, err := b64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])
		fmt.Println("payload auth=", string(payload))
		if err == nil {

			pair := bytes.SplitN(payload, []byte(":"), 2)

			if len(pair) == 2 && string(pair[0]) == user && string(pair[1]) == pass {
				fmt.Println("hash=", hash, " ", auth[len(basicAuthPrefix):])
				if hash == auth[len(basicAuthPrefix):] {
					return true
				}
			}
		}
	}
	return false
}
