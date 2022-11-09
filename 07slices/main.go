package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	//if we use this syntax we need to initialize values
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruit list is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// sublist from 1 to last
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// sublist from 1 to 3(non-inclusive)
	//fruitList = append(fruitList[1:3])
	//fmt.Println(fruitList)

	// sublist from 0 to 3(non-inclusive)
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	// use make to define slice
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 236
	highScores[3] = 465
	//we will get an error as we defined the size as 4 ans these are already used
	//highScores[4] = 238

	//It will reallocate the memory if we use append for slice
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//How to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "java"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
