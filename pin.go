package main

import (
	"math/rand"
	"time"
	"strconv"
	"strings"
)


func getPIN() string {
	rand.Seed(time.Now().UnixNano())
	pin := toString(complete(partial()))

	return pin
}

func randInts(min, max int) (int, int, int) {
	num := min + rand.Intn(max-min)
	if num > 99 {
		return num / 100 % 10, num / 10 % 10, num % 10
	}
	if num < 10 {
		return 0, num, 0
	}
	return num / 10 % 10, num % 10, 0
}

func partial() []int {
	y, yy, _ := randInts(0, 100)
	m, mm, _ := randInts(1, 13)
	d, dd, _ := randInts(1, 29)
	x, xx, xxx := randInts(111, 1000)
	// x, xx, xxx := randInts(980, 1000) // safe range

	return []int{y, yy, m, mm, d, dd, x, xx, xxx}
}

func complete(ints []int) []int {
	var sum int
	for i, v := range ints {
		v = v * (2 - i%2)
		if v >= 10 {
			v -= 9
		}
		sum += v
	}
	lastDigit := (100 - sum) % 10
	return append(ints, lastDigit)
}

func toString(a []int) string {
	b := make([]string, len(a)+1)

	for i, v := range a {
		if i == 6 {
			b = append(b, "-")
		}
		b = append(b, strconv.Itoa(v))
	}
	return strings.Join(b, "")
}
