package main

import (
	"fmt"
)

func main() {
	//initialize array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("print array value : ", arr)

	//loop over array
	fmt.Println("loop over array")
	for i := 0; i < 5; i++ {
		fmt.Printf("Key : %d value : %d \n", i, arr[i])
	}

	//range over array
	fmt.Println("Range over array")
	for key, value := range arr {
		fmt.Printf("key : %d value : %d \n", key, value)
	}

	//blank range over array
	fmt.Println("Blank index range over array")
	for _, value := range arr {
		fmt.Printf("blank range value : %d \n", value)
	}

	fmt.Println("length of array : ", len(arr))
}
