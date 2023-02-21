// This is a script to practice the basic grammar of the Golang language.
// Ref: https://www.w3schools.com/go/

package main

import (
	"fmt"
)

var glob1 int

// glob2 := 2 // this is a illegal syntax

func main() {
	/*		This comment style is same as C++ 	*/

	/*
		Variable define
		1. var variablename type = value
		2. variablename := value
		for example:
	*/
	var variablename1 string = " :)" // type is string (var variablename type = value)
	var variablename2 = "Jwl"        // type is inferred
	variablename3 := 4234            // type is inferred
	var x1, x2, x3, x4 int = 1, 3, 5, 7
	var (
		y1 int
		y2 int = 2
		y3 int = 4
		y4     = 6
	)

	/*
		Constants
	*/
	const PI = 3.14159265
	// for readability
	const (
		A int = 1
		B     = 2.21234
		C     = "Hi!"
	)

	/*
		Data Types
			bool
			int, int8, int16, int32, int64
			uint, uint8, uint16, uint32, uint64
			float32, float64
			string
	*/
	glob1 = variablename3

	fmt.Print("Hello World!")
	fmt.Print(variablename1)
	fmt.Print(" ", variablename2, "\n")
	fmt.Printf("glob1 has value: %#x and type: %T\n", glob1, glob1) // %T is indicating the type of a variable.
	fmt.Println(x1, x2, x3, x4)
	fmt.Println(y1, y2, y3, y4)
	fmt.Println()

	/*
		Array
		1. var array_name = [length]datatype{values} // here length is defined
		2. var array_name = [...]datatype{values} // here length is inferred
		3. array_name := [length]datatype{values} // here length is defined
		4. array_name := [...]datatype{values} // here length is inferred
		For example:
	*/
	var arr1 = [5]string{"Volvo", "BMW", "Ford", "Mazda", "Benz"}
	var arr2 = [...]int{1, 2, 3, 4, 5}
	arr3 := [5]int{} //not initialized
	arr4 := [...]int{1, 2, 3, 4}
	arr5 := [5]int{1: 10, 2: 40} // Initialize Only Specific Elements

	fmt.Println(len(arr1), arr1)
	fmt.Println(arr2[0])
	arr3[4] = 4
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println()
	/*
		Slices
		1. slice_name := []datatype{values}
		For example:
	*/
	myslice1 := []int{}
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}

	myslice1 = append(myslice1, 2)
	fmt.Println("len of myslice1: ", len(myslice1))
	fmt.Println("cap of myslice1: ", cap(myslice1))
	fmt.Println(myslice1)

	myslice2 = append(myslice2, "!") // slice_name = append(slice_name, element1, element2, ...)
	fmt.Println("len of myslice2: ", len(myslice2))
	fmt.Println("cap of myslice2: ", cap(myslice2))
	fmt.Println(myslice2)

	// myslice := myarray[start:end] // A slice made from the array
	myslice3 := arr1[2:4]
	fmt.Println("len of myslice3: ", len(myslice3))
	fmt.Println("cap of myslice3: ", cap(myslice3))
	fmt.Println(myslice3)

	// slice_name := make([]type, length, capacity)
	myslice4 := make([]int, 5, 10)
	myslice4 = append(myslice4, myslice1...)
	fmt.Println("len of myslice4: ", len(myslice4))
	fmt.Println("cap of myslice4: ", cap(myslice4))
	fmt.Println(myslice4)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers) // copy(dest, src)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

	/*
		Arithmetic Operators: 		+, -, *, /, %, ++, --
		Assignment Operators: 		=, +=, -=, *=, /=, %=, &=, |=, ^=, >>=, <<=
		Comparison Operators:		==, !=, <, >, >=, <=
		Logical Operators:			&&, ||, !
		Bitwise Operators: 			&, |, ^, <<, >>
	*/
	/*
		Conditions
	*/
	var a int = 5
	b := 3
	if a > b {
		fmt.Printf("%v is greater than %v\n", a, b)
	}

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}

	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	/*
		Switch
	*/
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Work day")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	/*
		Loops
	*/
	// for statement1; statement2; statement3
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		} else if i == 4 {
			break
		}
		fmt.Println(i)
	}
	// for index, value := array|slice|map
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	/*
		Function
	*/
	myMessage()
	familyName("Hilla", 24)
	fmt.Println(sum1(1, 2), sum2(1, 2))
	a1, b1 := myFunction(5, "Hello")
	_, b2 := myFunction(3, "Hi")
	fmt.Println(a1, b1, b2)

	/*
		Struct
	*/
	var pers1 Person

	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	printPerson(pers1)

	/*
		Maps
			var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
			b := map[KeyType]ValueType{key1:value1, key2:value2,...}

			var a = make(map[KeyType]ValueType)
			b := make(map[KeyType]ValueType)

			var a map[KeyType]ValueType
	*/
	var a3 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b3 := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Printf("a\t%v\n", a3)
	fmt.Printf("b\t%v\n", b3)

	var a4 = make(map[string]string) // The map is empty now
	a4["brand"] = "Ford"
	a4["model"] = "Mustang"
	a4["year"] = "1964"
	fmt.Println(a4["brand"])
	delete(a4, "year")
	fmt.Println(a4)
	a4["year"] = "1964"
	fmt.Println(a4)

	var b4 map[string]string
	fmt.Println(b4 == nil)

	// Check For Specific Elements in a Map
	val1, ok1 := a4["brand"]
	_, ok4 := a4["day"]
	fmt.Println(val1, ok1)
	fmt.Println(ok4)

	// Maps Are References
	b5 := a4
	b5["year"] = "1970"
	fmt.Println("After change to b5:")
	fmt.Println(a4)

	for k, v := range a4 {
		fmt.Printf("%v : %v, ", k, v)
	}
}

func myMessage() {
	fmt.Println("I just got executed!")
}

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func sum1(x int, y int) int {
	return x + y
}

func sum2(x int, y int) (result int) {
	result = x + y
	return result
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
