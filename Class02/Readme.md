# Go Code Overview - Operations in Map, Functions, Strings, and Methods

## Introduction
This Go code provides a practical demonstration of various concepts, including operations in maps, functions, strings, and methods. Let's explore each aspect with theoretical explanations and example code snippets.

## Custom Types and Stringer Interface

### RollNumber Type
- **Definition**: `RollNumber` is a custom type based on the underlying type `int`.
- **Stringer Interface Implementation**: The `String` method is implemented for the `RollNumber` type, satisfying the `Stringer` interface. It formats the roll number as a string.

```go
type RollNumber int 

func (r RollNumber) String() string {
    return fmt.Sprintf("Roll Number: %d", r)
}
```

### Password Type
- **Definition**: `Password` is a custom type based on the underlying type `string`.
- **isValid Method**: The `Password` type has a method named `isValid` that checks if the password is valid by ensuring its length is at least 8 characters.

```go
type Password string

func (p Password) isValid() bool {
    return len(p) >= 8
}
```

## Student Struct and Methods

### Student Type
- **Definition**: The `student` struct represents a student with fields such as `name`, `password`, `age`, and `rollno`.

```go
type student struct {
    name     string
    password Password
    age      int
    rollno   RollNumber
}
```

### isAdult Method
- **Definition**: The `isAdult` method is associated with the `student` type. It checks if the student is an adult based on the age condition (age > 18).

```go
func (s student) isAdult() bool {
    return s.age > 18
}
```

## Main Function
- **Map Operations**: The main function demonstrates the creation of a map (`mymap`) and iterates through its key-value pairs using a `for` loop.
- **Student Instance**: A `student` instance (`s`) is created with sample data.
- **Password Validation**: The `isValid` method is called to validate the password. If invalid, an error message is printed.
- **Stringer Interface Usage**: The `Stringer` interface is utilized when printing the roll number using `s.rollno.String()`.

```go
func main() {
    // Map Operations
    mymap := make(map[int]string)
    mymap[1] = "one"
    mymap[2] = "two"
    mymap[3] = "three"
    mymap[4] = "four"

    // Iterating over Map
    for key, value := range mymap {
        fmt.Println("key is ", key, "value is ", value)
    }

    // Student Instance
    var s = student{"Raj", "sdakn", 1, 1}

    // Password Validation
    if !s.password.isValid() {
        fmt.Println("password is not valid")
        return
    }

    // Stringer Interface Usage
    fmt.Printf("roll number is %v\n", s.rollno.String())
}
```

## `make` Function in Go

### Slices
The `make` function is commonly used to create slices with a specified length and capacity.

```go
// Creating a slice of integers with length 3 and capacity 5
mySlice := make([]int, 3, 5)
```

### Maps
The `make` function is used to create a map with an initial capacity.

```go
// Creating a map with string keys and integer values
myMap := make(map[string]int, 10)
```

### Channels
The `make` function is used to create a channel.

```go
// Creating an unbuffered channel of integers
myChannel := make(chan int)
```

### Important Notes
- The `make` function initializes the data structure and returns a reference to it.
- It is used for slices, maps, and channels to allocate memory and set up the initial state.

The use of `make` ensures that the created data structure is properly initialized and ready to use. It's a convenient way to allocate memory and set up the initial state for dynamic data structures in Go.
