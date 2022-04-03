package main

import (
	"fmt"
	"sort"
)

//TODO https://www.hackerrank.com/challenges/one-week-preparation-kit-plus-minus/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-one
/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	//var res []float32
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var negInt, posInt, zero, size float64
	size = float64(len(arr))
	for _, val := range arr {
		if val < 0 {
			negInt++
		} else if val > 0 {
			posInt++
		} else {
			zero++
		}
	}
	if zero == 0 {
		zero = 1
	}
	fmt.Printf("%.6f\n", posInt/size)
	fmt.Printf("%.6f\n", negInt/size)
	fmt.Printf("%.6f\n", zero/size)
}

func main() {
	arr := []int32{-4, -5, 1, 2}
	plusMinus(arr)
}
