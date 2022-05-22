package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gomodule/redigo/redis"
	"log"
	"test-go/config"
	"test-go/dao"
	"test-go/model"
	"test-go/myvalidator"
)

/*
	参数验证框架：
		https://github.com/go-playground/validator
	文档：
		https://godoc.org/gopkg.in/go-playground/validator.v9
*/

/**
redis
*/
func main3() {
	conn := config.RedisDefaultPool.Get()

	reply, err := conn.Do("get", "name")
	if err != nil {
		log.Println(err)
		return
	}
	ret, err := redis.String(reply, err)
	log.Println(ret)
}

func main() {
	router := gin.Default()
	// 注册参数验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err1 := v.RegisterValidation("topicurl", myvalidator.TopicUrl)
		err2 := v.RegisterValidation("topiclist", myvalidator.TopicList)
		if err1 != nil || err2 != nil {
			return
		}
	}

	// 单条帖子
	v1 := router.Group("/v1/topics")

	// 此处可以插入各种逻辑，只是使用代码块增加可读性
	{
		// localhost:8080/v1/topics?username=colorpen
		v1.GET("", dao.GetTopicList)

		// localhost:8080/v1/topics/13
		v1.GET("/:topic_id", config.CacheDecorator(dao.GetTopicDetail, "topic_id", "topic_%s", model.Topic{}))

		// 如果使用了中间件，下面的所有路由都需要中间件，进行token判断
		v1.Use(dao.MustLogin())
		{
			// 新增Topic
			v1.POST("/add", dao.NewTopic)

			// 删除Topic
			v1.DELETE("/del/:topic_id", dao.DelTopic)
		}
	}

	// 多条帖子批量处理
	v2 := router.Group("/v1/multi/topics")
	// 此处可以插入各种逻辑，只是使用代码块增加可读性
	{
		// 如果使用了中间件，下面的所有路由都需要中间件，进行token判断
		v2.Use(dao.MustLogin())
		{
			// 新增多条Topics
			v2.POST("/add", dao.NewTopicBatch)
		}
	}
	router.Run()
}
