package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age  int32
}

var m map[string]interface{}

func Get(name string, pval interface{}) error {
	//先判断是否传入的指针
	if reflect.TypeOf(pval).Kind() != reflect.Ptr {
		return errors.New("pval is not a ptr")
	}

	for k, v := range m {
		if k == name {
			//判断类型是否一致
			if reflect.TypeOf(pval).Elem() == reflect.TypeOf(v) {
				//获取Ptr元素值，并Set
				reflect.ValueOf(pval).Elem().Set(reflect.ValueOf(v))
				return nil
			} else {
				return errors.New("value type not equal")
			}
		}
	}
	return errors.New("not find the key" + name)
}

func main() {
	m = make(map[string]interface{})
	m["aaa"] = "abc"
	m["bbb"] = 123
	m["xiaoming"] = Human{name: "xiaoming", age: 32}
	var val Human
	if err := Get("xiaoming", &val); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Get xiaoming=", val)
	}

	if err := Get("aaa", &val); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Get aaa=", val)
	}
}
