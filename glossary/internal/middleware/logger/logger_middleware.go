package logger

import (
	"context"
	"fmt"
	"secondhand_glossary/platform"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoLogger struct {
	client *mongo.Client
}

// io.Writer implementation
func (m *mongoLogger) Write(p []byte) (n int, err error) {
	c := m.client.Database("secondhand").Collection("logs")
	fmt.Println(string(p))
	_, err = c.InsertOne(context.TODO(), bson.M{
		"msg": string(p),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	return len(p), nil
}

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2021-05-18 20:00:00"),
		})
	}
	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2021-05-18 20:00:00"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func NewLogger(p *platform.Connection) *mongoLogger {
	return &mongoLogger{
		client: p.Logger,
	}
}
