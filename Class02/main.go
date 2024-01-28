package main

import "fmt"

type RollNumber int 

func (r RollNumber) String() string {
	return fmt.Sprintf("jnsdajkajsd jkbakbhb %d",r)
}

type Password string

func (p Password) isValid() bool {
	if len(p) < 8 {
		return false
	}
	return true
}

// func (<var> <type>) 

type student struct {
	name string
	password Password
	age int
	rollno RollNumber
}

func (s student) isAdult() bool {
	if s.age > 18 {
		return true
	}
	return false
}


func main() {
	mymap := make(map[int]string)
	mymap[1] = "one"
	mymap[2] = "two"
	mymap[3] = "three"
	mymap[4] = "four"

	for key,value:= range mymap {
		fmt.Println("key is ", key, "value is ", value)
	}

	var s = student{"Raj", "sdakn", 1,1}

	if !s.password.isValid() {
		fmt.Println("password is not valid")
		return
	}

	fmt.Printf("roll number is %v\n",s.rollno.String())
}
