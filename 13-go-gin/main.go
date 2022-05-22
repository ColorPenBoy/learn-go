package main

/*
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"test-go/model"
	"time"
)

// 新版GORM写法
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // 1、默认gorm会在表名后面加复数，设置这个参数后，即可关闭
			//NoLowerCase:   true,                              // skip the snake_casing of names
			//NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		), //加SQL的执行Log
	})
	if db == nil || err != nil {
		fmt.Println("数据库连接错误: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 单条查询
	tc := model.TopicClass{}
	// 3、此处也可以直接指定表名
	db.Table("topic_class").First(&tc, 2)
	fmt.Println(tc)

	// 多条查询
	var tcs []model.TopicClass
	db.Find(&tcs)
	fmt.Println(tcs)

	// where查询
	var tcs2 []model.TopicClass
	// db.Where("class_name=?", "技术类").Find(&tcs2)
	db.Where(&model.TopicClass{ClassName: "技术类"}).Find(&tcs2)
	fmt.Println(tcs2)

	// 新增数据
	topic := model.Topic{
		TopicTitle:      "新增标题",
		TopicShortTitle: "新增短标题",
		TopicUrl:        "https://new.topic.url",
		UserIP:          "127.0.0.10",
		TopicScore:      11,
		TopicDate:       time.Now(),
	}
	fmt.Println(db.Create(&topic).RowsAffected)
	fmt.Println(topic.TopicID)
}

*/

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"test-go/config"
	"test-go/dao"
	"test-go/myvalidator"
	"time"
)

/*
	参数验证框架：
		https://github.com/go-playground/validator
	文档：
		https://godoc.org/gopkg.in/go-playground/validator.v9
*/
//func main() {
//	count := 0
//	go func() {
//		for {
//			fmt.Println("执行", count)
//			count++
//			time.Sleep(time.Second * 1)
//		}
//	}()
//
//	channel := make(chan os.Signal)
//	// 延迟5秒发送 interrupt 停止信号
//	go func() {
//		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
//		select {
//		case <-ctx.Done():
//			channel <- os.Interrupt // 向通道发送信号
//		}
//	}()
//
//	signal.Notify(channel)
//	s := <-channel // 通道接收信号
//	fmt.Println(s)
//}

func main() {
	conn := config.RedisDefaultPool.Get()

	reply, err := conn.Do("get", "name")
	if err != nil {
		log.Println(err)
	}
	ret, err := redis.String(reply, err)
	log.Println(ret)
}

func main2() {
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

	//router.Run()
	server := &http.Server{
		Addr:    "8080",
		Handler: router,
	}
	go func() { // 启动web服务
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("服务器启动失败")
		}
	}()
	go func() { // 通过协程的方式，启动DB
		config.InitDB()
	}()

	// 监听异常信号，并且关闭服务
	config.ServerNotify()
	// 这里还可以做一些释放连接或者善后工作
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("服务器关闭")
	}
	log.Fatal("服务器优雅退出")
}
