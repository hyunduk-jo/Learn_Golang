package main

import "fmt"

func main() {
	//Array - 선언한 길이로 고정됨
	//선언을 하고 초기화는 안함. 정의한 타입의 zero-value로 초기화가 된다.
	var arr1 [3]int
	//선언과 초기화를 동시에 하는데 초기화에 아무것도 넣어주지 않으면 정의한 타입의 zero-value로 초기화 된다.
	arr2 := [3]int{}
	fmt.Println(arr1)
	fmt.Println(arr2)
	//Slice - 길이가 동적임
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	//append를 통해 slice에 워소를 추가할 수 있다. Array에는 append 불가능하다.
	slice1 = append(slice1, 5, 6, 7)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	//[a:b] a 이상 b 미만까지 자른 배열을 보여준다. 원본 slice는 그대로 유지
	fmt.Println(slice1[2:4])
}
