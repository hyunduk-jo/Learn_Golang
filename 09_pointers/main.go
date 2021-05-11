package main

import (
	"fmt"
)

func main() {
	a := 5
	//b의 값에 a의 주소를 대입한다. a의 타입은 int 이므로 b의 타입은 *int이다.
	b := &a
	fmt.Println(a, b)
	//b는 a의 주소를 가르키고 그 주소의 값을 가져오려면 *를 써서 불러오면 된다.
	fmt.Println(*b)
	//포인터 변수 또한 자신의 주소를 갖는다.
	fmt.Println(&a)
	fmt.Println(&b)

	//*b는 a의 값이므로 a의 값이 20으로 변경된다.
	*b = 20
	fmt.Println(a)

}
