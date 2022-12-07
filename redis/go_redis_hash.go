package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Model struct {
	Str1    string   `redis:"str1"`
	Str2    string   `redis:"str2"`
	Int     int      `redis:"int"`
	Bool    bool     `redis:"bool"`
	Ignored struct{} `redis:"-"`
}


func HSetModel() {

	_, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, "key", "str1", "hello")
		rdb.HSet(ctx, "key", "str2", "world")
		rdb.HSet(ctx, "key", "int", 123)
		rdb.HSet(ctx, "key", "bool", 1)
		return nil
	});

	if err != nil {
		panic(err)
	}
}

func HGetToModel()  {

	var model1 Model
	// Scan all fields into the model.
	if err := rdb.HGetAll(ctx, "key").Scan(&model1); err != nil {
		panic(err)
	}

	fmt.Printf(" str1=%s,str2=%s,int=%d,bool=%t",model1.Str1,model1.Str2,model1.Int,model1.Bool)

}

