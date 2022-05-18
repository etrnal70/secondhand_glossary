package platform

import (
	"context"
	"secondhand_glossary/internal/config"
	"secondhand_glossary/platform/cache"
	"secondhand_glossary/platform/logger"
	"secondhand_glossary/platform/persistence"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Connection struct {
	Cache       *redis.Client
	Logger      *mongo.Client
	Persistence *gorm.DB
}

func InitPlatform(c config.Config) *Connection {
	cache := cache.InitCacheDB(c)
	logger := logger.InitLoggerDB(c)
	persistence := persistence.InitPersistenceDB(c)
	return &Connection{
		Cache:       cache,
		Logger:      logger,
		Persistence: persistence,
	}
}

func (c *Connection) Close() error {
	cacheErr := c.Cache.Close()
	if cacheErr != nil {
		return cacheErr
	}

	loggerErr := c.Logger.Disconnect(context.Background())
	if loggerErr != nil {
		return loggerErr
	}

	_, persistenceErr := c.Persistence.DB()
	if persistenceErr != nil {
		return persistenceErr
	}

	return nil
}
