package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int64 = 10
	reflectSetValue(&a)
	fmt.Println(a)
}

func reflectSetValue(x interface{}) {
	/*v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(120)
	}*/

	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	fmt.Println(v.Elem().Kind())
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(120)
	}

}
