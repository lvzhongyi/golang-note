package main

import (
	"reflect"
	"fmt"
)
//反射	reflect //////////////////
/**

先放一放,后面再理解

 */

type  ty struct {
	a1 string
	a2 string
}

func main() {
	ty := ty{"aa", "bb"}
	t := reflect.TypeOf(ty)//得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	//我的理解是:得到类型
	fmt.Println(t)
	v := reflect.ValueOf(ty)//得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	//我的理解是:得到值
	fmt.Println(v)

	tag := t.Elem().Field(0).Tag //获取定义在struct里面的标签
	name := v.Elem().Field(0).String()//获取存储在第一个字段里面的值
	fmt.Println(tag, name)

}
