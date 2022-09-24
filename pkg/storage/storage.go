package storage

import (
	"context"
	"fmt"
	"time"

	"valuation/internal/conf"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func NewDB(c *conf.Data) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	return DB, err
}

func NewRedis(c *conf.Data) (*redis.Client, error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password:     c.Redis.Password, // no password set
		Username:     c.Redis.Username,
		DB:           int(c.Redis.Db),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return Redis, nil
}
