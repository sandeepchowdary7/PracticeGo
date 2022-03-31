package day1

/*
InitDay1
	1. Given a list of numbers and a number k,
	2. return whether any two numbers from the list add up to k.
	3. For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
*/

func InitDay1() {
	a := []int{1, 2, 3, 5, 6, 9}
	println(checkSum(a, 15))
}

func checkSum(a []int, k int) bool {
	for i, x := range a {
		for j, y := range a {
			if i != j && x+y == k {
				return true
			}
		}
	}
	return false
}
