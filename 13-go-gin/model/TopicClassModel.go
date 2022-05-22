package model

// 与数据库表topic_class对应实体

type TopicClass struct {
	ClassId     int `gorm:"primaryKey"`
	ClassName   string
	ClassRemark string
	ClassType   string `gorm:"column:classtype"`
}

// 2、手动给topic_class指定表名称

func (TopicClass) TableName() string {
	return "topic_class"
}
