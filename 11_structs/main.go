package main

import (
	"fmt"
)

//User struct를 만듬
type User struct {
	city      string
	age       int
	name      string
	isMarried bool
	gender    string
	data      map[string]int
}

//* 있으면 pointer receiver 없으면 value receiver
//pointer receiver는 변경을 하면 원래의 struct도 변경 되고 value receiver는 변경을 못하고 읽는것만 가능
func (u *User) moveCity(newCity string) {
	u.city = newCity
}

func changeName(u *User, newName string) {
	u.name = newName
}

func (u *User) getMarried() {
	u.isMarried = true
}

func main() {
	user1 := User{isMarried: false, name: "Tommy", city: "Sejong", age: 27, gender: "Male", data: map[string]int{"a": 1}}
	fmt.Println(user1)
	user1.moveCity("seoul")
	fmt.Println(user1)
	changeName(&user1, "hyunduk")
	fmt.Println(user1)
	user1.getMarried()
	fmt.Println(user1)
}
