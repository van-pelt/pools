package server

import (
	"github.com/van-pelt/pools/pkg/config"
	"net"
	"net/http"
	"strconv"
)

func NewServer(handler http.Handler, conf *config.Config) http.Server {

	return http.Server{
		Addr:    net.JoinHostPort(conf.Server.Url, strconv.Itoa(conf.Server.Port)),
		Handler: handler,
	}
}
