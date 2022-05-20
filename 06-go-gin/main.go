package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"test-go/dao"
	"test-go/myvalidator"
)

/*
	参数验证框架：
		https://github.com/go-playground/validator
	文档：
		https://godoc.org/gopkg.in/go-playground/validator.v9
*/

func main() {
	router := gin.Default()
	// 注册参数验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("topicurl", myvalidator.TopicUrl)
		if err != nil {
			return
		}
	}

	v1 := router.Group("/v1/topics")

	// 此处可以插入各种逻辑，只是使用代码块增加可读性
	{
		// localhost:8080/v1/topics?username=colorpen
		v1.GET("", dao.GetTopicList)

		// localhost:8080/v1/topics/13
		v1.GET("/:topic_id", dao.GetTopicDetail)

		// 如果使用了中间件，下面的所有路由都需要中间件，进行token判断
		v1.Use(dao.MustLogin())
		{
			// 新增Topic
			v1.POST("/add", dao.NewTopic)

			// 删除Topic
			v1.DELETE("/del/:topic_id", dao.DelTopic)
		}
	}

	router.Run()
}
