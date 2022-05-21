package main

//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//// 新版GORM写法
//func main() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/test_go?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if db == nil || err != nil {
//		fmt.Println("数据库连接错误: ", err.Error())
//	}
//	sqlDB, _ := db.DB()
//	defer sqlDB.Close()
//
//	rows, _ := db.Raw("SELECT topic_id, topic_title FROM topics").Rows()
//	for rows.Next() {
//		var t_id int
//		var t_title string
//		rows.Scan(&t_id, &t_title)
//		fmt.Println(t_id, t_title)
//	}
//}

//旧版GORM写法
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_go?charset=utf8mb4&parseTime=True&loc=Local")
	if db == nil || err != nil {
		fmt.Println("数据库连接错误: ", err.Error())
	}
	defer db.Close()

	rows, _ := db.Raw("SELECT topic_id, topic_title FROM topics").Rows()
	for rows.Next() {
		var t_id int
		var t_title string
		rows.Scan(&t_id, &t_title)
		fmt.Println(t_id, t_title)
	}
}

/*
	参数验证框架：
		https://github.com/go-playground/validator
	文档：
		https://godoc.org/gopkg.in/go-playground/validator.v9
*/
//func main() {
//	router := gin.Default()
//	// 注册参数验证器
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		err1 := v.RegisterValidation("topicurl", myvalidator.TopicUrl)
//		err2 := v.RegisterValidation("topiclist", myvalidator.TopicList)
//		if err1 != nil || err2 != nil {
//			return
//		}
//	}
//
//	// 单条帖子
//	v1 := router.Group("/v1/topics")
//
//	// 此处可以插入各种逻辑，只是使用代码块增加可读性
//	{
//		// localhost:8080/v1/topics?username=colorpen
//		v1.GET("", dao.GetTopicList)
//
//		// localhost:8080/v1/topics/13
//		v1.GET("/:topic_id", dao.GetTopicDetail)
//
//		// 如果使用了中间件，下面的所有路由都需要中间件，进行token判断
//		v1.Use(dao.MustLogin())
//		{
//			// 新增Topic
//			v1.POST("/add", dao.NewTopic)
//
//			// 删除Topic
//			v1.DELETE("/del/:topic_id", dao.DelTopic)
//		}
//	}
//
//	// 多条帖子批量处理
//	v2 := router.Group("/v1/multi/topics")
//	// 此处可以插入各种逻辑，只是使用代码块增加可读性
//	{
//		// 如果使用了中间件，下面的所有路由都需要中间件，进行token判断
//		v2.Use(dao.MustLogin())
//		{
//			// 新增多条Topics
//			v2.POST("/add", dao.NewTopicBatch)
//		}
//	}
//
//	router.Run()
//}
