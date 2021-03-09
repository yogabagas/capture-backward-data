package command

import (
	m "my-github/capture-backward-data/datastore/mongo"
	p "my-github/capture-backward-data/datastore/postgre"
	"my-github/capture-backward-data/datastore/redis"
	"my-github/capture-backward-data/domain/repository"
)

var (
	postgre repository.PostgreRepository
	mongo   repository.MongoRepository
	rdb     redis.InternalRedis
)

func init() {
	postgre = InitPostgre()
	mongo = InitMongo()
	rdb = InitRedis()
}

func InitPostgre() repository.PostgreRepository {
	return p.NewPostgreConn("", "", "", "", 0)
}

func InitMongo() repository.MongoRepository {
	return m.MongoMustConnect("", "")
}

func InitRedis() redis.InternalRedis {
	host := make(map[string]string)
	host[""] = ""
	return redis.NewRedisClient(host, "", 0)
}
