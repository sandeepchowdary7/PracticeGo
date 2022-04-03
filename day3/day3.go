package main

//TODO https://www.hackerrank.com/challenges/one-week-preparation-kit-mini-max-sum/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-one
import (
	"fmt"
	"sort"
)

//Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.
/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var minSum int32
	for i := 0; i < len(arr)-1; i++ {
		minSum += arr[i]
	}
	maxSum := minSum - arr[0] + arr[len(arr)-1]
	fmt.Println(minSum, maxSum)
}

func main() {
	arr := []int32{8, 6, 3, 9}
	miniMaxSum(arr)
}
