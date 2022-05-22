package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"test-go/config"
	"test-go/model"
)

func GetTopicDetail(c *gin.Context) {
	tidstr := c.Param("topic_id")
	topicId, _ := strconv.Atoi(tidstr)

	topic := model.Topic{}
	config.DbHelper.Find(&topic, topicId) // 从数据库取
	log.Println("GetTopicDetail - 从数据库取")
	c.Set("dbResult", topic) // 高阶函数 - 先执行缓存装饰器，判断是否有缓存，以及是否需要走业务中前几行查询数据库代码
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
