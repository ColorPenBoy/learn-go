package main

import (
	"github.com/gin-gonic/gin"
	"test-go/dao"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/topics")

	// 此处可以插入各种逻辑，只是使用代码块增加可读性
	{
		// localhost:8080/v1/topics?username=colorpen
		v1.GET("", func(c *gin.Context) {
			if c.Query("username") == "" {
				c.String(200, "获取帖子列表")
			} else {
				c.String(200, "获取用户名=%s的帖子列表", c.Query("username"))
			}
		})

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
