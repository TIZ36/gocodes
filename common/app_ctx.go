package common

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type AppContext struct {
	MysqlClient   *sql.DB
	HttpClient    *resty.Client
	InMemoryCache *bigcache.BigCache
	RedisClient   *redis.Client
}

var (
	datetimePrecision = 2
	AppCtx            = &AppContext{}
)

func NewAppContext(config Config) {
	setUpMysql(config)
	setUpCache()
	//setUpRedis()
	return
}

func setUpMysql(config Config) {
	db, err := sql.Open("mysql", config.Mysql.Dsn)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	AppCtx.MysqlClient = db
}

func setUpCache() {
	cache, e := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	if e != nil {
		panic(e)
	}

	AppCtx.InMemoryCache = cache
}

func (ctx *AppContext) DestroyAppCtx() {
	ctx.MysqlClient.Close()
	fmt.Println("destroy ok")
}
