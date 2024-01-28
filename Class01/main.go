// 1 step package main  main or <folder name>
package main

import (
	"errors"
	"fmt"
)

// import

//3 step func main

//scopes

//public and private
// public matalab sabke liye Heck
// private matalab sirf usi file me heck

//comma,ok syntax

func main() {
	var meow string
	meow = "cat"
	var age = 10      //strictly typed language (GO - Lexer) ke jagah se
	name := "Devansh" //memory allocate and initialize
	fmt.Printf("My age is %d %s and my name is %s\n ", age, meow, name)
	area, perimeter := getAreaAndPerimerter(10, 20)
	fmt.Printf("Area is %d and perimeter is %d \n", area, perimeter)

	comma, ok := checkAge(1)
	if ok != nil {
		panic(ok)
	}

	fmt.Println(comma)
}

// func  <func name> (<param name> <param type>) <return type> {}

func checkAge(age int) (string, error) {
	if age < 18 {
		return "Not allowed", errors.New("CHOTA HAI TU BHAI DOOD PI")
	}
	return "Allowed", nil
}

func getAreaAndPerimerter(a int, b int) (int, int) { //rectangle
	area := a * b
	perimeter := 2 * (a + b)
	return area, perimeter
}
