package main

import (
	"fmt"
)

func main() {
	list1 := []string{"a","b","c"}
	list2 := []string{"first","second","third"}

	m := make(map[string]string)

	for i, _ := range list1 {
		m[list1[i]] = list2[i]
	}

	fmt.Println(m)
}

