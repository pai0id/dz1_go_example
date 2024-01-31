package main

func addElement(array []int, elem int) []int {
	return append(array, elem)
}

func removeElement(array []int) []int {
	if len(array) > 0 {
		return array[:len(array)-1]
	}
	return array
}

func addOneToArray(array []int) []int {
	for i := range array {
		array[i]++
	}
	return array
}
