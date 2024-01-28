# Go (Golang) - Basic Concepts

## Table of Contents
- [Introduction](#introduction)
- [History](#history)
- [Features](#features)
- [Syntax](#syntax)
- [Data Types](#data-types)
  - [Basic Data Types](#basic-data-types)
  - [More Data Types](#more-data-types)
  - [Advanced Data Types](#advanced-data-types)
- [Installation](#installation)
- [Hello World](#hello-world)
- [Variable Declaration](#variable-declaration)
- [Scopes](#scopes)

## Introduction
Go, also known as Golang, is an open-source programming language developed by Google. It is designed for simplicity, efficiency, and ease of use. Go is statically typed and compiled, with a focus on concurrency.

## History
Go was created by Robert Griesemer, Rob Pike, and Ken Thompson at Google in 2007 and was first released to the public in 2009.

## Features
- Concurrency support
- Simple and clean syntax
- Efficient compilation
- Strongly typed
- Garbage collection
- Built-in testing
- Cross-platform support

## Syntax
Go has a C-like syntax with some unique features. It uses curly braces for code blocks and relies on indentation for grouping. It has a clean and straightforward syntax, making it easy to read and write.

## Data Types

### Basic Data Types
- `int`, `float64`, `string`: Basic numeric and string types
- `bool`: Boolean type
- `array`, `slice`, `map`: Composite data types
- `struct`: User-defined composite type
- `interface`: Abstract type representing a set of methods

#### Examples
```go
package main

import "fmt"

func main() {
    // Numeric types
    var x int = 10
    var y float64 = 3.14

    // String type
    var message string = "Hello, Go!"

    // Boolean type
    var isTrue bool = true

    // Array and slice
    var arr [3]int = [3]int{1, 2, 3}
    var slice []int = []int{4, 5, 6}

    // Map
    var myMap map[string]int = map[string]int{"one": 1, "two": 2}

    // Struct
    type Person struct {
        Name string
        Age  int
    }

    var person1 Person = Person{Name: "John", Age: 30}

    // Interface
    type Shape interface {
        area() float64
    }

}
```

### More Data Types
- `complex64`, `complex128`: Complex number types
- `byte`, `rune`: Aliases for `uint8` and `int32`, respectively
- `uintptr`: Integer type that can hold the value of a pointer

#### Examples (Continued)
```go
package main

import "fmt"

func main() {
    // Complex numbers
    var complexNum complex64 = 2 + 3i
    fmt.Println("Complex Number:", complexNum)

    // Byte and Rune
    var char byte = 'A'
    var unicodeChar rune = 'Σ'
    fmt.Printf("Byte: %c, Rune: %c\n", char, unicodeChar)

    // Pointer and uintptr
    var pointerValue *int = new(int)
    *pointerValue = 42
    var uintptrValue uintptr = uintptr(unsafe.Pointer(pointerValue))
    fmt.Println("Pointer Value:", *pointerValue)
    fmt.Println("Pointer as uintptr:", uintptrValue)

}
```

### Advanced Data Types
- **Channels**: Used for communication between goroutines in Go's concurrency model.
- **Function Types**: Functions can be assigned to variables and passed as arguments.
- **Slices of Slices**: Multi-dimensional slices for more complex data structures.
- **Time**: Handling time-related operations.

#### Examples (Continued)
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Channel of integers
    ch := make(chan int)

    // Goroutine sending data to the channel
    go func() {
        ch <- 42
    }()

    // Receiving data from the channel
    value := <-ch
    fmt.Println("Received from channel:", value)

    // Function type and variable
    type MathFunction func(int, int) int
    var operation MathFunction = add
    result := operation(3, 4)
    fmt.Println("Result of operation:", result)

    // Slice of slices (2D slice)
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("Matrix:", matrix)

    // Current time and formatting
    currentTime := time.Now()
    formattedTime := currentTime.Format("Mon Jan 2 15:04:05 MST 2006")
    fmt.Println("Current Time:", currentTime)
    fmt.Println("Formatted Time:", formattedTime)
}
```

## Installation
To install Go, follow these steps:
1. Visit the official [Go download page](https://golang.org/dl/).
2. Download the installer for your operating system.
3. Follow the installation instructions provided on the website.

## Hello World
The traditional "Hello, World!" program in Go is simple:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Variable Declaration
Variables in Go are declared using the `var` keyword, followed by the variable name and type.

```go
var age int = 25
var name string = "Alice"
```

## Scopes
Go has block-level scoping, and variables declared within a block are only accessible within that block.

```go
package main

import "fmt"

func main() {
    // Variable 'x' is only accessible within this block
    var x int = 42
    fmt.Println(x)
}
