package main

import "fmt"

func main() {
	//make를 사용
	emails := make(map[string]string)
	emails["tom"] = "tom@email.com"
	emails["paul"] = "paul@email.com"
	fmt.Println(emails)

	//초기화 안하면 zero-value가 들어간다. 지금은 string 이기 때문에 ""가 들어간다.
	students := map[string]string{"tom": "tom", "paul": "paul"}
	fmt.Println(students)
}
