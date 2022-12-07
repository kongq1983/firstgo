package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "172.16.5.19:6379",
	Password: "jkct-redis-123456", // no password set
	DB:       12,  // use default DB
})

func init() {

	//rdb = redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//fmt.Println(res)

	//err := rdb.Set(ctx, "", "", 0).Err()

}

func Set(key string, val string)  {

	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		//panic(err)
		fmt.Printf(" set key=%s fail ! err=%s \n",key,err)
	}
}

func Get(key string) string  {

	var val, err = rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("%s , key does not exist \n",val)
	} else if err != nil {
		fmt.Printf(" get key=%s fail ! err=%s \n",key,err)
	} else {
		return val
	}

	return "";
}

func GetByCommand(key string) string {

	var val, err = rdb.Do(ctx, "get", key).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return ""
		}
	}
	fmt.Println(val.(string))

	return val.(string)

}

func SetNX(key string, val string, expiration time.Duration)  {

	err := rdb.SetNX(ctx, key, val, expiration).Err()
	if err != nil {
		//panic(err)
		fmt.Printf(" set key=%s fail ! err=%s \n",key,err)
	}
}