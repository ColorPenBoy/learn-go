package model

import "time"

type Topic struct {
	TopicID int `json:"id" gorm:"primaryKey"`
	// 标题长度必须在4-20位
	TopicTitle string `json:"title" binding:"min=4,max=20"`
	// 短标题和标题不能一样
	TopicShortTitle string `json:"shortTitle" binding:"required,nefield=TopicTitle"`
	// 可以不填，如果填写，需要符合 `topicurl` 的规则
	TopicUrl string `json:"url" binding:"omitempty,topicurl"`
	// user ip必须是ipv4的形式
	UserIP string `json:"ip" binding:"ipv4"`
	// Score要么不填，要填必须大于5分
	TopicScore int       `json:"s ore" binding:"omitempty,gt=5"`
	TopicDate  time.Time `json:"date" binding:"required"`
}

// 2、手动给topic_class指定表名称

func (Topic) TableName() string {
	return "topics"
}

type Topics struct {
	// 添加dive，进行嵌套验证
	TopicList     []Topic `json:"topics" binding:"gt=0,lt=3,topiclist,dive"`
	TopicListSize int     `json:"size"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title, "", "", "", 0, time.Time{}}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username" binding:"required"` // 必填参数
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
