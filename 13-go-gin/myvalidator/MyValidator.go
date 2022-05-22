package myvalidator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"test-go/model"
)

//func TopicUrl(
//	v *validator.Validate,
//	topicStruct reflect.Value,
//	currentStructOrField reflect.Value,
//	field reflect.Value,
//	fieldType reflect.Type,
//	fieldKind reflect.Kind,
//	param string
//	) bool
func TopicUrl(fl validator.FieldLevel) bool {
	// 判断model类型是否为Topic
	_, ok1 := fl.Top().Interface().(model.Topics)
	_, ok2 := fl.Top().Interface().(model.Topic)
	if ok1 || ok2 {
		getValue := fl.Field().String()
		fmt.Println(getValue)
		// URL只能是数字，字母，下划线，且必须在4--10字符
		if ret, _ := regexp.MatchString(`^\w{4,10}$`, getValue); ret {
			return true
		}
	}
	return false
}

func TopicList(fl validator.FieldLevel) bool {
	// 判断model类型是否为Topic
	topics, ok := fl.Top().Interface().(model.Topics)
	if ok && topics.TopicListSize == len(topics.TopicList) {
		return true
	}
	return false
}
