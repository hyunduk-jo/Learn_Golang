package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//type을 쓰지 않아도 초기화를 해주면 그에 맞는 type으로 된다.
	//초기화를 하지 않으면 각 타입에 맞는 zero-value로 초기화 된다.(int -> 0, float -> 0.0, string -> "" etc)
	var str string
	var name string = "Hyunduk"
	//선언과 초기화 동시에 하는 변수 선언 := 초기화를 꼭 해주어야 한다.
	age := 27

	fmt.Println(str, name, age)

	//변수의 타입은 reflect.TypeOf(age) 또는 Printf("%T", age)이고 크기는 unsafe.Sizeof(age)이다.
	fmt.Println(reflect.TypeOf(age))
	fmt.Printf("%T\n", age)
	fmt.Println(unsafe.Sizeof(age))
}
