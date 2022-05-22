package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"log"
)

// 缓存装饰器
func CacheDecorator(h gin.HandlerFunc, param string, redKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 缓存的redis判断
		getID := context.Param(param)
		redisKey := fmt.Sprintf(redKeyPattern, getID)

		log.Println("CacheDecorator - 缓存装饰器")

		conn := RedisDefaultPool.Get()
		defer conn.Close()
		ret, err := redis.Bytes(conn.Do("get", redisKey))
		if err != nil { // 缓存中没有
			h(context) // 执行目标方法
			dbResult, exists := context.Get("dbResult")
			if !exists { // 数据库无结果
				dbResult = empty
			} else { // 数据库存在结果
				log.Println("从数据库读取")
				retData, _ := ffjson.Marshal(dbResult)
				conn.Do("setex", redisKey, 20, retData)
				context.JSON(200, dbResult)
			}
		} else { // 缓存有，直接抛出结果
			log.Println("从Redis缓存读取")
			ffjson.Unmarshal(ret, &empty)
			context.JSON(200, empty)
		}
	}
}
