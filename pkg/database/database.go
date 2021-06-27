package database

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/van-pelt/pools/pkg/config"
	"github.com/van-pelt/pools/pkg/logger"
	"go.uber.org/fx"
	"strconv"
)

type Storage struct {
	DB *pg.DB
}

func NewDBInstance(log *logger.Logger, cfg *config.Config) (*Storage, error) {

	log.InfoF("Connect to DB:%s:%s:%s:%s", cfg.Database.Host, strconv.Itoa(cfg.Database.Port), cfg.Database.User, cfg.Database.DBName)

	db := pg.Connect(&pg.Options{
		Addr:     cfg.Database.Host + ":" + strconv.Itoa(cfg.Database.Port),
		User:     cfg.Database.User,
		Password: cfg.Database.Pass,
		Database: cfg.Database.DBName,
	})
	contex := context.Background()
	err := db.Ping(contex)
	if err != nil {
		log.FatalF("DB Ping:%s", err.Error())
		return nil, err
	}
	log.Info("DB Ping:OK")

	return &Storage{DB: db}, nil

}

func NewDB(lc fx.Lifecycle, storage *Storage, log *logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			log.Info("DB.Shutdown")
			storage.DB.Close()
			return nil
		},
	})
}
