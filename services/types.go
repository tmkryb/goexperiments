package services

import (
	"errors"
)

//Go exported functions always start with big letters!
//Slices expand when appending new values to them
func CreateSlice(elemNum int) []int {
	intSlice := make([]int, 0, 10)
	for i := 0; i < elemNum; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}

//Array is always of fixed size
func CreateArray(elemNum int) ([10]int, error) {
	intArray := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if elemNum > 10 {
		return intArray, errors.New("elemNum is to big! limit it to 10")
	}
	for i := 0; i < elemNum; i++ {
		intArray[i] = i
	}
	return intArray, nil
}
