package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type myError struct {
	err string
	num int
}

func (e *myError) Error() string {
	return fmt.Sprintf("Rand number: %d Error: %s", e.num, e.err)
}

func main() {
	pls := time.Now()
	var b = 10 // capacity
	var a = make([]int, 0, b)
	c := make(chan int)
	for i := 0; i < b; i++ {
		num := rand.Intn(b) + 1
		chrand(a, num, c)
		a = append(a, <-c)
	}
	fmt.Println("Random result: ", a)
	defer fmt.Println("Process time: ", time.Since(pls))
	err := qsort(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sorting result: ", a)
	}
	defer close(c)
}

func chrand(a []int, num int, c chan int) {
	go func() {
		for i := 0; i < len(a); i++ {
			if a[i] == num {
				e := myError{err: "random repeated on element: " + strconv.Itoa(i), num: num}
				fmt.Println(e.Error())
				rand.Seed(time.Now().UnixNano()) //new random pool
				num = rand.Intn(cap(a)) + 1
				i = 0
			}
		}
		c <- num
	}()
}

func qsort(a []int) error {
	if len(a) < 2 {
		return errors.New("Slice have less than 2 elements")
	}
	mid := a[len(a)/2]
	left := 0
	right := len(a) - 1
	for left <= right {
		for a[left] < mid {
			left++
		}
		for a[right] > mid {
			right--
		}
		if left <= right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}
	fmt.Println("After sort: ", a)
	if left <= len(a)-1 {
		qsort(a[:left])
		qsort(a[left:])
	}
	return nil
}
