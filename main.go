package main

import (
	"fmt"
	"github.com/RookieJoel/GoAPI-essential/Jo" //import the whole folder as a package. s in the same directory, the package name must be the same
	// "github.com/RookieJoel/GoAPI-essential/Jo/internal/testInternal" // internal package, can only be imported by the same module
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	jo.SayHello() //calling the exported function from the jo package
	// jo.sayGoodbye() //this will cause an error because sayGoodbye is unexported (private)
}

func Variable() {
	//variable types
	// Boolean
	//  - true or false
	// Numeric
	//  - Integer (int, int8, int16, int32, int64)
	//  - Floating-point (float32, float64)
	//  - Complex (complex64, complex128)
	// String

	//type declaration
	// 1. Using the `var` keyword
	var name string = "Rookie Joel" // var <name> <type> = <value>
	var age int = 30                // var <name> <type> = <value>
	var isEmployed bool = true      // var <name> <type> = <value>
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Employed:", isEmployed)
	// 2. Using the `:=` shorthand (only within functions)
	city := "New York" // <name> := <value>
	fmt.Println("City:", city)

	// 3. Using the `const` keyword for constants
	const pi float64 = 3.14159 // const <name> <type> = <value>
	fmt.Println("Pi:", pi)

	// 4. Using the `var` keyword without initialization (zero value)
	var score int // score will be initialized to 0
	fmt.Println("Score:", score)

	// 5. Using the `var` keyword with multiple variables
	var x, y int = 10, 20 // var <name1>, <name2> <type> = <value1>, <value2>
	fmt.Println("X:", x, "Y:", y)	

	// 6. Using the `var` keyword with a slice
	var numbers []int = []int{1, 2, 3, 4, 5} // var <name> []<type> = []<type>{<value1>, <value2>, ...}
	fmt.Println("Numbers:", numbers)
}

func PreProcess() {
	//preprocessing if else statements
	num1 := 10
	num2 := 20

	if sumNum := num1 + num2; sumNum > 15 {
		fmt.Println("The sum is greater than 15: ", sumNum)
	}
}

func Iteration() {
	//iteration
	// 1. Using a `for` loop
	for i := 0; i < 5; i++ {
		fmt.Println("For Loop Iteration:", i)
	}

	// 2. Using a `while`-like loop (Go doesn't have a `while` keyword, but you can use `for`)
	j := 0
	for j < 5 {
		fmt.Println("While-like Loop Iteration:", j)
		j++
	}

	// 3. Using a `range` loop to iterate over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

func DataStruc(){
	//data structures
	// 1. Arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5} // fixed size array
	fmt.Println("Array:", arr) //output: Array: [1 2 3 4 5]

	// 2. Slices (dynamic arrays)
	slice := []int{1, 2, 3, 4, 5} // dynamic size slice
	slice = append(slice, 6) // appending to a slice
	fmt.Println("Slice:", slice) //output: Slice: [1 2 3 4 5 6]
	subSlice := slice[1:4] // slicing a slice
	fmt.Println("SubSlice:", subSlice) //output: SubSlice: [2 3 4]
	fmt.Println("Slice:", slice) //output: Slice: [1 2 3 4 5]

	//coverting array to slice
	arrToSlice := arr[:] // converting array to slice
	fmt.Println("Array to Slice:", arrToSlice) //output: Array to Slice: [1 2 3 4 5]


	// 3. Maps (key-value pairs)
	mm := make(map[string]int) // creating an empty map
	mm["one"] = 1               // adding key-value pairs
	mm["two"] = 2
	mm["three"] = 3
	fmt.Println("Map:", mm) //output: Map: map[one:1 three:3 two:2]

	// Creating a map with initial values
	m := map[string]int{"one": 1, "two": 2, "three": 3} // map with string keys and int values
	fmt.Println("Map:", m) //output: Map: map[one:1 three:3 two:2]

	// Adding a new key-value pair to the map
	m["four"] = 4
	fmt.Println("Updated Map:", m) //output: Updated Map: map[four:4 one:1 three:3 two:2]
	// Accessing a value by key
	value, exists := m["two"] // checking if the key exists
	if exists {
		fmt.Println("Value for 'two':", value) //output: Value for 'two': 2
	} else {
		fmt.Println("'two' does not exist in the map")
	}
	// Removing a key-value pair from the map
	delete(m, "three") // removing the key "three"
	fmt.Println("Map after deletion:", m) //output: Map after deletion: map[four:4 one:1 two:2]
	// Iterating over a map
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
		//output: Key: four, Value: 4
		//output: Key: one, Value: 1
		//output: Key: two, Value: 2
	}
	

	// 4. Structs (custom data types)
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Rookie Joel", Age: 30}
	fmt.Println("Struct:", p)

	// 5. Functions as first-class citizens
	add := func(a int, b int) int {
		return a + b
	}
	result := add(5, 10)
	fmt.Println("Function Result:", result)

	// 6. Interfaces (defining behavior)
	type Speaker interface {
		Speak() string
	}

	type Dog struct {
		Name string
	}

	func (d Dog) Speak() string {
		return d.Name + " says Woof!"
	}

	type Cat struct {
		Name string
	}

	func (c Cat) Speak() string {
		return c.Name + " says Meow!"
	}

	var s Speaker

	s = Dog{Name: "Buddy"}
	fmt.Println(s.Speak())

	s = Cat{Name: "Whiskers"}
	fmt.Println(s.Speak())
	

	// 7. Channels (for concurrent programming)

}

func ErrorHandling() {
	//error handling
	
	func divide(a, b int) (int, error) {
		if b == 0 {
			return 0, error.New("division by zero") //returning an error if b is zero
		}
		return a / b, nil
	}
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) //output: Error: division by zero
	} else {
		fmt.Println("Result:", result)
	}	
	
}