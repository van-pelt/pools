package router

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/van-pelt/pools/pkg/cache"
	"github.com/van-pelt/pools/pkg/config"
	"github.com/van-pelt/pools/pkg/database"
	"github.com/van-pelt/pools/pkg/handlers"
	"github.com/van-pelt/pools/pkg/logger"
	"github.com/van-pelt/pools/pkg/router/ctxparam"
	"github.com/van-pelt/pools/pkg/server"
	"go.uber.org/fx"
	"net/http"
	"os"
	"strconv"
)

func NewRouter() *httprouter.Router {
	return httprouter.New()
}

func RegisterRoutes(lc fx.Lifecycle, conf *config.Config, log *logger.Logger, storage *database.Storage, cache *cache.Cache) {

	router := NewRouter()

	basePath, err := os.Getwd()
	if err != nil {
		log.ErrorF("basePath:%s", err.Error())
		return
	}
	router.ServeFiles("/"+conf.AllowDir+"*filepath", http.Dir(basePath+"/"+conf.AllowDir))

	serv := server.NewServer(router, conf)
	root := context.Background()
	ctx := context.WithValue(root, "params", ctxparam.NewCtxParam(basePath, conf.AllowDir, conf.Server.Url+":"+strconv.Itoa(conf.Server.Port)+"/"))

	/*user := []byte("u.b.c.s.bravo@gmail.com")
	pass := []byte("fff")*/

	router.GET("/", handlers.Index(ctx))
	router.GET("/login", handlers.Login(ctx))
	router.POST("/auth", handlers.Auth(ctx, storage, cache))
	router.GET("/dashboard", handlers.Dashboard(ctx, storage, cache))

	router.GET("/users/", handlers.Users(ctx, storage, cache))
	router.DELETE("/users/delete/:id", handlers.UsersDelete(ctx, storage, cache))
	router.GET("/users/:id", handlers.GetUser(ctx, storage, cache))
	//router.GET("/users/add", handlers.Dashboard(ctx, storage, cache))

	router.GET("/protected", handlers.Protected(ctx, storage, cache))

	//fmt.Println("PP=" + filepath.FromSlash(http.Dir(basePath+conf.AllowDir)))
	//router.ServeFiles(conf.AllowDir+"*filepath", http.Dir(filepath.FromSlash(basePath+conf.AllowDir)))
	//http.Handle("/", http.FileServer(http.Dir(basePath+"/web/")))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info("server.Start")

				err := serv.ListenAndServe()
				if err != nil {
					log.FatalF(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("server.Shutdown")
			return serv.Shutdown(ctx)
		},
	})

}

/*func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}*/

func BasicAuth(h httprouter.Handle, cache *cache.Cache) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials

		user, password, hasAuth := r.BasicAuth()

		fmt.Println("BB=", user, password, hasAuth)

		hash := b64.StdEncoding.EncodeToString([]byte(user + ":" + password))
		el := cache.Get(string(hash))

		if hasAuth && user == el.Email && password == el.Pass {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
