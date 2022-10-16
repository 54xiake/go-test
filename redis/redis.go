package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// redis

// 定义一个全局变量
var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed! err : %v\n", err)
		return
	}
	fmt.Println("redis连接成功！")
	opt()
	zset()
	set_rand()
}

func opt() {
	// 存普通string类型，10分钟过期
	redisdb.Set("test:name", "科科儿子", time.Minute*10)
	// 存hash数据
	redisdb.HSet("test:class", "521", 42)
	// 存list数据
	redisdb.RPush("test:list", 1) // 向右边添加元素
	redisdb.LPush("test:list", 2) // 向左边添加元素
	// 存set数据
	redisdb.SAdd("test:set", "apple")
	redisdb.SAdd("test:set", "pear")
}

// zset 添加
func zset() {
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 93, Member: "JAVA"},
		redis.Z{Score: 96, Member: "C++"},
		redis.Z{Score: 92, Member: "PYTHON"},
		redis.Z{Score: 100, Member: "GOLANG"},
	}
	// 添加元素到redis中
	_, err := redisdb.ZAdd(key, items...).Result()
	if err != nil {
		fmt.Printf("redis zadd failed! err:%v\n", err)
		return
	}
	fmt.Println("数据传输成功！！！")
}

func set_rand() {
	//返回几个中随机任意一个
	str1, err := redisdb.SRandMember("test:set").Result()
	if err != nil {
		fmt.Printf("redis set rand failed! err:%v\n", err)
		return
	}
	fmt.Println(str1)

	//返回几个中随机任意3个
	str2, err := redisdb.SRandMemberN("test:set", 3).Result()
	if err != nil {
		fmt.Printf("redis set rand failed! err:%v\n", err)
		return
	}
	fmt.Println(str2)

}
