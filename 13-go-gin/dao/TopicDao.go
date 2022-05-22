package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-go/config"
	"test-go/model"
)

func GetTopicDetail(c *gin.Context) {
	topicId, _ := strconv.Atoi(c.Param("topic_id"))

	topic := model.Topic{}
	config.DbHelper.Find(&topic, topicId)
	c.JSON(200, topic)
}

// 单帖新增
func NewTopic(c *gin.Context) {
	query := model.Topic{}
	// GET 请求绑定参数到model中
	err := c.BindJSON(&query)
	if err != nil {
		c.String(400, "参数错误:%s", err.Error())
	} else {
		c.JSON(200, query)
	}
}

// 多帖批量新增
func NewTopicBatch(c *gin.Context) {
	topics := model.Topics{}
	// GET 请求绑定参数到model中
	err := c.BindJSON(&topics)
	if err != nil {
		c.String(400, "参数错误:%s", err.Error())
	} else {
		c.JSON(200, topics)
	}
}

func DelTopic(c *gin.Context) {
	// 判断登陆
	c.String(200, "删除帖子")
}

func GetTopicList(c *gin.Context) {
	query := model.TopicQuery{}
	// GET 请求绑定参数到model中
	err := c.BindQuery(&query)
	if err != nil {
		c.String(400, "参数错误:%s", err.Error())
	} else {
		c.JSON(200, query)
	}
}

// Gin中间件，判断登陆
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断token参数是否存在
		if token, status := c.GetQuery("token"); status == false {
			c.String(http.StatusUnauthorized, "token不存在")
			c.Abort() // 不会再向下走逻辑
		} else {
			fmt.Printf("Token存在: %s \n", token)
			c.Next() // 继续往下走逻辑
		}
	}
}
