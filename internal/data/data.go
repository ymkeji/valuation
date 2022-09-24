package data

import (
	"valuation/internal/conf"
	"valuation/pkg/storage"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v9"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGoodRepo, NewUserRepo, storage.NewDB, storage.NewRedis)

// Data .
type Data struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := storage.NewDB(c)
	if err != nil {
		return nil, nil, err
	}
	rdb, err := storage.NewRedis(c)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		rdb.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, redis: rdb}, cleanup, nil
}
