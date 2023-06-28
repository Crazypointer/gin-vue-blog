package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initClient()
	if err != nil {
		panic(err)
	}
	err = rdb.Set("score", 100, 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	r := rdb.Keys("*")
	fmt.Println(r)
	keys, err := r.Result()
	fmt.Println(keys)

}
