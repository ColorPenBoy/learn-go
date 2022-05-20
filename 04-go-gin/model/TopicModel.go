package model

type Topic struct {
	TopicID    int    `json:"id"`
	TopicTitle string `json:"title" binding:"required"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username" binding:"required"` // 必填参数
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}
