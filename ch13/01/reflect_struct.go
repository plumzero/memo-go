
package main

import (
	"fmt"
	"reflect"
)

type ss struct {
	int
	string
	bool
	float64
}

func (s ss) Method1(i int) string {
	 return "结构体方法1"
}

func (s *ss) Method2(i int) string {
	return "结构体方法2"
}

var (
	structValue = ss{
		20,
		"结构体",
		false,
		64.0,
	}
)

func main() {
    // 反射结构体
    fmt.Println("========== 引用 ==========")
	v := reflect.ValueOf(&structValue)
	fmt.Println("String:  ", v.String())
	fmt.Println("Type:    ", v.Type())
	fmt.Println("Kind:    ", v.Kind())
	fmt.Println("CanAddr: ", v.CanAddr())
	fmt.Println("CanSet:  ", v.CanSet())
	if v.CanAddr() {
		fmt.Println("Addr      :", v.Addr())		// 获取地址
		fmt.Println("UnsafeAddr:", v.UnsafeAddr())	// 获取自由地址
	}
    // 获取方法数量
    fmt.Println("可用方法数量:", v.NumMethod())
    if v.NumMethod() > 0 {
        i := 0
        for ; i < v.NumMethod(); i++ {
            fmt.Printf("  -> %v\n", v.Method(i).String())
        }
        // 通过名称获取方法
        fmt.Println("Method1 MethodByName:", v.MethodByName("Method1").String())
        fmt.Println("Method2 MethodByName:", v.MethodByName("Method2").String())
    }
    // 值变量
    fmt.Println("========== 值变量 ==========")
	v = reflect.ValueOf(structValue)
	fmt.Println("String:  ", v.String())
	fmt.Println("Type:    ", v.Type())
	fmt.Println("Kind:    ", v.Kind())
	fmt.Println("CanAddr: ", v.CanAddr())
	fmt.Println("CanSet:  ", v.CanSet())
	if v.CanAddr() {
		fmt.Println("Addr      :", v.Addr())		// 获取地址
		fmt.Println("UnsafeAddr:", v.UnsafeAddr())	// 获取自由地址
	}
    // 获取方法数量
    fmt.Println("可用方法数量:", v.NumMethod())
    if v.NumMethod() > 0 {
        i := 0
        for ; i < v.NumMethod(); i++ {
            fmt.Printf("  -> %v\n", v.Method(i).String())
        }
        // 通过名称获取方法
        fmt.Println("Method1 MethodByName:", v.MethodByName("Method1").String())
        fmt.Println("Method2 MethodByName:", v.MethodByName("Method2").String())
    }

    switch v.Kind() {
        // 结构体
    case reflect.Struct:
        fmt.Println("========== 结构体 ==========")
        fmt.Println("NumField: ", v.NumField())
        if v.NumField() > 0 {
            // 遍历字段
            for i := 0; i < v.NumField(); i++ {
                field := v.Field(i)
                fmt.Printf("  -> %-8v %v\n", field.Type(), field.String())
            }
            // 通过名称查找字段
            if v := v.FieldByName("ptr"); v.IsValid() {
                fmt.Println("FieldByName(ptr):", v.Type().Name())
            }
            // 通过函数查找字段
            v := v.FieldByNameFunc(func(s string) bool { return len(s) > 3})
            if v.IsValid() {
                fmt.Println("FieldByNameFunc:", v.Type().Name())
            }
        }
    }
}