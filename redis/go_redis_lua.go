package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func IncrByLua(key string) (int, error) {

	var incrBy = redis.NewScript(`
	local key = KEYS[1]
	local change = ARGV[1]
	
	local value = redis.call("GET", key)
	if not value then
	  value = 0
	end
	
	value = value + change
	redis.call("SET", key, value)
	
	return value
	`)

	keys := []string{key}
	values := []interface{}{+1}
	var num, err = incrBy.Run(ctx, rdb, keys, values...).Int()

	return num , err

}


/**
  商品销售
 */
func DecrProductByLua(key string) (int, error) {

	var incrBy = redis.NewScript(`
	--商品KEY
	local key = KEYS[1]
	--购买数
	local val = ARGV[1]
	--现有总库存
	local stock = redis.call("GET", key)
	if (tonumber(stock)<=0) 
	then
		--没有库存
		print("没有库存")
		return -1
	else
		--获取扣减后的总库存=总库存-购买数
		local decrstock=redis.call("DECRBY", key, val)
		if(tonumber(decrstock)>=0)
		then
			--扣减购买数后没有超卖，返回现库存
			print("没有超卖，现有库存数"..decrstock)
			return decrstock
		else
			--超卖了，把扣减的再加回去
			redis.call("INCRBY", key, val)
			print("超卖了，现有库存"..stock.."不够购买数"..val)
			return -2
		end
	end
	`)

	keys := []string{key}
	values := []interface{}{1}
	var num, err = incrBy.Run(ctx, rdb, keys, values...).Int()

	return num , err

}


func SetToUpper(key string, val string) (string, error) {

	var incrBy = redis.NewScript(`
	local key = KEYS[1]
	local value = ARGV[1]
	
	local value = string.upper(value)
	
	redis.call("SET", key, value)

	local value = redis.call("GET", key)
	
	return value
	`)

	keys := []string{key}
	values := []interface{}{val}

	var val2,err = incrBy.Run(ctx, rdb, keys, values...).Text()

	print(val2)
	return val2,err

}

/**
 只用于例子演示
 同个线程可多次重入
 */
func DistributeLock(key string, threadId int) (int, error) {

	var incrBy = redis.NewScript(`
	local key = KEYS[1]
	local value = ARGV[1]
	
	if (redis.call('exists', KEYS[1]) == 0) then 
	  	redis.call('set', key, value)
	  	redis.call('expire', key, ARGV[2])
	  	return 1
    else
		local loadValue = redis.call("GET", key)
		if loadValue == nil then
			return 0
		end
		
		if tonumber(value) == tonumber(loadValue) then
			return 1
		else
			return 0
		end

  	end
	

	return 0
	`)

	keys := []string{key}
	values := []interface{}{threadId,1000}

	var val2,err = incrBy.Run(ctx, rdb, keys, values...).Int()

	return val2,err
}

func IterAllKeys()  {
	//iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	iter := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

/**
	sadd set1 one one two three two
	遍历某个set的值
 */
func IterSetKeys()  {
	//iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	//iter := rdb.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
	iter := rdb.SScan(ctx, "set1", 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

/**
	hmset hashmap1 one 1 two 2 three 3 five 5 six 6
 */
func IterHashKeys()  {
	//iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	//iter := rdb.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
	iter := rdb.HScan(ctx, "hashmap1", 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}


/**
   ZADD zset1 1 "z-one" 2 "z-two" 3 "z-three"
*/
func IterZSetKeys()  {
	//iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	//iter := rdb.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
	iter := rdb.ZScan(ctx, "zset1", 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}


//SCAN遍历数据库中的键
//SSCAN遍历set中的键
//HSCAN遍历hash中的键
//ZSCAN遍历zset中的键盘