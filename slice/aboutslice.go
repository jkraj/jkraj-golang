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
}