package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/van-pelt/pools/internal/user/service"
	"github.com/van-pelt/pools/pkg/cache"
	"github.com/van-pelt/pools/pkg/database"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID       int64
	FullName string
	Email    string
}

type UserFull struct {
	ID       int64
	FullName string
	Email    string
	Password string
}

type UserResponse struct {
	Response []User
}

type UserDeleteResponse struct {
	Status string
	Code   int
}

func Users(ctx context.Context, storage *database.Storage, cache *cache.Cache) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//auth := r.Header.Get("Authorization")
		cookie, err := r.Cookie("bauth")
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		chl := cache.Get(cookie.Value)
		r.Header.Set("Authorization", "Basic "+cookie.Value)
		log.Println(chl.Email, chl.Pass, cookie.Value)
		if AuthB64(r, chl.Email, chl.Pass, cookie.Value) {
			userSrv := service.New(storage)
			user, err := userSrv.GetUsersData()
			if err != nil {
				fmt.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			var usr []User
			for i, v := range *user {
				log.Println(i, " ", v.Email)
				usr = append(usr, User{
					ID:       v.Id,
					FullName: v.FullName,
					Email:    v.Email,
				})
			}
			var p = UserResponse{Response: usr}
			data, err := json.Marshal(p)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			}
			fmt.Fprintf(w, string(data))
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
}

func UsersDelete(ctx context.Context, storage *database.Storage, cache *cache.Cache) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("DELETE")
		//auth := r.Header.Get("Authorization")
		cookie, err := r.Cookie("bauth")
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		chl := cache.Get(cookie.Value)
		fmt.Println("CACHE:", chl)
		r.Header.Set("Authorization", "Basic "+cookie.Value)
		fmt.Println("header=:", r.Header.Get("Authorization"))
		if AuthB64(r, chl.Email, chl.Pass, cookie.Value) {
			id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
			if err != nil {
				fmt.Println(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			userSrv := service.New(storage)
			if err = userSrv.DelUsersData(id); err != nil {
				fmt.Println(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			cookie, err := r.Cookie("bauth")
			cache.Del(cookie.Value)
			cookie.Value = "000"
			//fmt.Fprintf(w, "DELETE")
			p := UserDeleteResponse{
				Status: "OK",
				Code:   http.StatusOK,
			}
			data, err := json.Marshal(p)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			}
			fmt.Fprintf(w, string(data))
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
}

func GetUser(ctx context.Context, storage *database.Storage, cache *cache.Cache) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//auth := r.Header.Get("Authorization")
		cookie, err := r.Cookie("bauth")
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		chl := cache.Get(cookie.Value)
		r.Header.Set("Authorization", "Basic "+cookie.Value)
		log.Println(chl.Email, chl.Pass, cookie.Value)
		if AuthB64(r, chl.Email, chl.Pass, cookie.Value) {
			userSrv := service.New(storage)
			id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
			if err != nil {
				fmt.Println(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			user, err := userSrv.GetUserByID(id)
			if err != nil {
				fmt.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			usr := UserFull{
				ID:       user.Id,
				Email:    user.Email,
				FullName: user.FullName,
				Password: user.Password,
			}

			data, err := json.Marshal(usr)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			}
			fmt.Fprintf(w, string(data))
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
}
