package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [20]int
	rand.Seed(time.Now().UnixNano())
	Fill(&arr)
	Print(arr)
	Sort(&arr)
	Print(arr)
}

func Fill(arr *[20]int) {
	for i, _ := range arr {
		arr[i] = rand.Intn(100+50+1) - 50
	}
}

func Print(arr [20]int) {
	for i, _ := range arr {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("")
}

func Sort(arr *[20]int) {
	var i, j, m, d int
	for i = 0; i < len(arr); i++ {
		m = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		d = arr[i]
		arr[i] = arr[m]
		arr[m] = d
	}
}
