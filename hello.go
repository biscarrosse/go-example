package main

import (
	"fmt"
	"errors"
	"math"
)


func main() {
	// var x int // just declare
	// var y int = 4
	// z := 3 // var z, type is inferred, assigned value of 3
	// x = 0 // init a value
	// var sum int = x + y + z
	// fmt.Println("sum:", sum)

	
	// ifElse(x)

	// idk()

	// mappps()

	// looper()

	// result := suma(7, 4)
	// fmt.Println(result)

	// res, err := sqrt(-13)
	// if err != nil {
	// 	fmt.Println(err)
	// 	} else {
	// 		fmt.Println(res)
	// }

	// makeHuman("Giulia", 21)

	m := 1
	n := 1
	
	increment(m)
	fmt.Println(m)

	incrementReally(&n) // sending pointer to memory reference
	fmt.Println(n)
}

func ifElse(x int) {
	if x > 2 {
		fmt.Println("x greater 2")
	} else if x < -2 {
		fmt.Println("x smaller 2")
	} else if x == 0{
		fmt.Println("x eq 0")
	} else {
		// idk
	}
}

// array / slice:
func idk() {
	var arr [5]int

	anotherArr := [4]int{1,2,3,4}

	slice := []int{12,123,45,3,2,4,5,} // how many I want
	fmt.Println("old slice:", slice)
	
	slice = append(slice, 13)
	arr[2] = 7
	
	fmt.Println("arr:", arr)
	fmt.Println("anotherArr:", anotherArr)
	fmt.Println("updated copy of slice:", slice)
}

// maps are like objects in JS, key: value pairs
func mappps() {
	coolMap := make(map[string]int)
	coolMap["car"] = 22
	coolMap["plane"] = 44

	fmt.Println("coolMap", coolMap)

	// read specific value:
	fmt.Println(coolMap["car"])

	delete(coolMap, "car")
}

// for loops & iteration:
func looper() {
	for i := 0; i < 5; i++ {
		fmt.Println("loop", i)
	}

	// while loop:
	j := 0
	for j < 5 {
		fmt.Println("loop j", j)
		j++
	}

	// iterate key value of arr or slice:
	arr := []string{"hey", "hi", "hello"}
	for index, value := range arr {
		fmt.Println("iterating key & value:", index, value)
	}

	// iterate key value of map:
	person := make(map[string]string)
	person["name"] = "Giacomo"
	person["surname"] = "Puccini"

	for key, value := range person {
		fmt.Println("key:", key, "value:", value)
	}

}

// multiple arguments:
func suma(a, b int) int {
	return a + b
}

// multiple return values + error:
func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("no negative numbers pls")
	} else {
		return math.Sqrt(num), nil
	}
}

// structs:
type human struct {
	name string // can be another nested struct
	age int
}
func makeHuman(name string, age int) human {
	newHuman := human{name: name, age: age}
	
	fmt.Println("new human is ready:", newHuman)
	// print only name:
	// fmt.Println("new human is ready:", newHuman.name)
	
	return newHuman
}

// pointers:
func increment(num int) {
	num++ // useless, not doing anything
}
func incrementReally(num *int) {
	// de-referencing pointer here, another asterisk:

	*num++ // useless, not doing anything
}