package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("numbers")
	check(err)

	fmt.Println(string(dat))

	num := strings.Split(string(dat), ", ")

	var pNum []int

	for i := 0; i < len(num); i++ {
		j, _ := strconv.Atoi(num[i])

		pNum = append(pNum, j)
	}

	fmt.Println(pNum)

	sort(pNum)

	fmt.Println(pNum)

}

func sort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			c := slice[i]
			if slice[i] < slice[j] {
				slice[i] = slice[j]
				slice[j] = c
			}
		}
	}
}
