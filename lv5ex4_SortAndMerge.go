package main

import (
	"fmt"
	"slices"
)

func SortAndMerge(first, second []int) []int {
	left := make([]int, len(first))
	copy(left, first)
	slices.Sort(left)

	right := make([]int, len(second))
	copy(right, second)
	slices.Sort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	left := []int{4, 1, 5, 0}
	right := []int{-1, 4, 5, 10}
	fmt.Println(SortAndMerge(left, right))

	// 	left:     []int{4, 1, 5, 0},
	// right:    []int{-1, 4, 5, 10},
	// expected: []int{-1, 0, 1, 4, 4, 5, 5, 10},
}

// Даны два слайса. Напишите программу с функцией SortAndMerge(left, right []int) []int,
// которая объединит слайсы в один отсортированный в два этапа:

//     отсортировать каждый слайс
//     объединить полученные слайсы в один

// Кстати, именно так работает алгоритм сортировки слиянием (merge sort).
