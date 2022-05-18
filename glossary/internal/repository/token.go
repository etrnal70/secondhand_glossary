package repository

import (
	"context"
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/middleware/jwt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type cacheDriver struct {
	Cache *redis.Client
}

func (d *cacheDriver) GetToken(td jwt.TokenDetails) (uint64, error) {
	userIdStr, err := d.Cache.Get(context.Background(), td.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	return userId, nil
}

func (d *cacheDriver) SetToken(userId uint, token jwt.TokenDetails) (err error) {
	at := time.Unix(token.AtExpires, 0)
	rt := time.Unix(token.RtExpires, 0)
	curTime := time.Now()

	err = d.Cache.Set(context.Background(), token.AccessUuid, strconv.Itoa(int(userId)), at.Sub(curTime)).Err()

	if err != nil {
		return
	}

	err = d.Cache.Set(context.Background(), token.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(curTime)).Err()

	return
}

func (d *cacheDriver) DeleteToken(id string) (int64, error) {
	deleted, err := d.Cache.Del(context.Background(), id).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

func (d *cacheDriver) SetRefreshToken() {
	panic("unimplemented")
}

func NewCacheDriver(c *redis.Client) domain.TokenRepository {
	return &cacheDriver{
		Cache: c,
	}
}
