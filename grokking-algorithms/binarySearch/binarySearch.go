package main

func binarySearch(list []int, item int) pos int {
	low:= 0
	high:=len(list) - 1
	for _, low <= high, i++ {
		mid = (low+high)
		guess = list[mid]
		switch {
		case guess == item:
			return mid
		case guess > item:
			high = mid -1
		default:
			low = mid +1
		}
	}
return pos
}

func main () {
	list:= make([]int, 5)
	fmt.Scanf(&list)
	fmt.Println(binarySearch(list, 4))
}