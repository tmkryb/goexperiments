package services

func CreateSlice(elemNum int) []int {
	intSlice := make([]int, 0, 10)
	for i := 0; i < elemNum; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}
