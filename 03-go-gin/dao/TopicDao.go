package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTopicDetail(c *gin.Context) {
	c.String(200, "获取topicid=%s的帖子", c.Param("topic_id"))
}

func NewTopic(c *gin.Context) {
	// 判断登陆
	c.String(200, "新增帖子")
}

func DelTopic(c *gin.Context) {
	// 判断登陆
	c.String(200, "删除帖子")
}

func GetTopicList() {

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
