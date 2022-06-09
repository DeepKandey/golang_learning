package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomatao", "Peach"}
	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("FruitList: ", fruitList)

	var fruitListSliced = append(fruitList[1:])
	fmt.Println("fruitListSliced: ", fruitListSliced)

	fruitList = append(fruitList[1:3])
	fmt.Println("FruitList: ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 564
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
