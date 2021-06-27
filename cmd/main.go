package main

import (
	"github.com/van-pelt/pools/pkg/cache"
	"github.com/van-pelt/pools/pkg/cfgreader"
	"github.com/van-pelt/pools/pkg/config"
	"github.com/van-pelt/pools/pkg/database"
	"github.com/van-pelt/pools/pkg/logger"
	"github.com/van-pelt/pools/pkg/router"
	"github.com/van-pelt/pools/pkg/server"
	"go.uber.org/fx"
	"os"
)

func main() {

	fx.New(
		fx.Provide(
			config.NewConfig,
			interface{}(func() (*logger.Logger, error) {
				return logger.New("test", 1, os.Stdout)
			}),
			database.NewDBInstance,
			server.NewServer,
			router.NewRouter,
			cache.NewCache,
		),
		fx.Invoke(router.RegisterRoutes),
		fx.Invoke(database.NewDB),
		//fx.Invoke(Pr),
	).Run()
}

func Pr(lc fx.Lifecycle, cf *cfgreader.CfgReader, ff *logger.Logger) {
	cf.R()
	ff.Critical("This is Critical!")
	// You can also use fmt compliant naming scheme such as log.Criticalf, log.Panicf etc
	// with small 'f'

	// Debug
	// Since default logging level is Info this won't print anything
	ff.Debug("This is Debug!")

	ff.Warning("This is Warning!")

	ff.Error("This is Error!")

	// Notice
	ff.Notice("This is Notice!")

	ff.Info("This is Info!")
}
