package logger

import (
	"context"
	"fmt"
	"os"
	"secondhand_glossary/internal/config"

   log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func InitLoggerDB(c config.Config) *mongo.Client {
	connString := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		c.MONGO_USERNAME,
		c.MONGO_PASSWORD,
		c.MONGO_HOST,
		c.MONGO_PORT,
	)

	logger_db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))

  if err != nil {
    log.Error("Cannot connect to mongo: ", err)
    os.Exit(1) // TODO Handle this better
  }

  log.Info("Connected to mongo instance")

  return logger_db
}

func InitMongoLogger(){}
