package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) ChangeName(newName string) {
	u.Name = newName
}

func (u *User) ChangeAge(newAge int) {
	u.Age = newAge

}

func main() {
	u := User{
		Name: "TOm",
		Age:  18,
	}

	u.ChangeName("Jerry")
	u.ChangeAge(17)

	fmt.Println(u.Name)
	fmt.Println(u.Age)

}
