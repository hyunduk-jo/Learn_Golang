package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for _, value := range arr {
		fmt.Println(value)
	}

	map1 := map[string]int{"age": 10, "height": 173}
	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
