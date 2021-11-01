package main

import "fmt"

func main() {
	fmt.Println("Hello, Algorithms")
	var i int = 10
	fmt.Println(i)

	var nums = []int{2, 7, 11, 3, 13}
	var target = 9

	var result = twoSum(nums, target)
	fmt.Println(result)

}

// We're going to use a HashMap and do a single pass on our array
// while keeping track of the numbers we encounceter and their indexes
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {

		if v, found := m[target-num]; found {
			return []int{v, idx}
		}

		m[num] = idx
	}
	return nil
}
