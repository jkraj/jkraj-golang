package main

import (
	"fmt"
	"reflect"
)

type DatabaseSpec struct {
	Users      []DatabaseUser `json:"users"`
}

type DatabaseUser struct {
	Name         string       `json:"name"`
	Password     string       `json:"password"`
	//PasswordFrom *ValueSource `json:"passwordFrom"`
	Role         string       `json:"role"`
}

func main()  {

	fmt.Println("Hello")

	u1 := DatabaseUser{"John", "pass1", "admin"}
	u2 := DatabaseUser{"koil", "pass1","read"}
	u3 := DatabaseUser{"Raj", "pass1","write"}

	u := []DatabaseUser{u1, u2, u3}

	U := DatabaseSpec{u}


	for _, user := range U.Users {
		//fmt.Println(user.Name)
		fmt.Println(reflect.TypeOf(user))
		fmt.Println(user.Name)
	}

	// understand basics
	p := [6]int{1,2,3,4,5,6}
	fmt.Println(p)

	var s = p[1:4]
	fmt.Println(s)

	p1 := []int{1,2,3,4}
	fmt.Println(p1)

	p2 := make([]int, 4, 4)
	p2 = append(p2, 1,2,3,4,5,6,7)
	fmt.Println(p2)
}