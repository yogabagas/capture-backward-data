package mongo

import (
	"context"
	"my-github/capture-backward-data/domain/model"
	"my-github/capture-backward-data/domain/repository"
	"time"

	"github.com/newrelic/go-agent/_integrations/nrmongo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoConnect(uri string) (c *mongo.Client, err error) {
	var (
		defaultConnectTimeout = 10 * time.Second
		defaultPingTimeout    = 2 * time.Second
	)
	ctx, _ := context.WithTimeout(context.Background(), defaultConnectTimeout)
	nrMon := nrmongo.NewCommandMonitor(nil)
	opts := []*options.ClientOptions{
		options.Client().SetConnectTimeout(defaultConnectTimeout).ApplyURI(uri).SetAppName("trackingsvc"),
		options.Client().SetMonitor(nrMon),
	}
	c, err = mongo.Connect(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, "failed to create mongodb client")
		return
	}
	ctx, _ = context.WithTimeout(context.Background(), defaultPingTimeout)
	if err = c.Ping(context.Background(), readpref.Primary()); err != nil {
		err = errors.Wrap(err, "failed to establish connection to mongodb server")
	}
	return
}

type mongoReadWriter struct {
	mg *mongo.Database
}

func MongoMustConnect(uri string, db string) repository.MongoRepository {
	c, err := MongoConnect(uri)
	if err != nil {
		panic(err)
	}
	return &mongoReadWriter{mg: c.Database(db)}
}

func (m *mongoReadWriter) BulkInsertDataAWB(ctx context.Context, awb []model.AWBDetailPartner) error {
	return nil
}
