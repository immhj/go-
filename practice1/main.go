package main

import (
	"fmt"
)

func main() {
	var sum float64
	var x float64
	var a [1000]int
	var mid int
	var l int = 1
	var p string
	var target string
	var fd int = 1
	var o int
	fmt.Scanf("%s", &target)
	for o = 0; o < len(target); o++ {
		if target[o] == '=' {
			break
		}
	}
	for i := 0; i < len(target); i++ {
		if i < o {
			if string(target[i]) == string('+') {
				fd = 1
				l++
			} else if string(target[i]) == string('-') {
				fd = -1
				l++
			} else if string(target[i]) >= string('0') && string(target[i]) <= string('9') {
				if a[l] == 0 {
					a[l] = (int(target[i]) - int('0')) * fd
				} else {
					a[l] = a[l]*10 + (int(target[i])-int('0'))*fd
				}
			} else if target[i] >= byte('a') && target[i] <= byte('z') {
				p = string(target[i])
				if a[l] != 0 {
					x += float64(a[l])
					a[l] = 0
				} else {
					x += float64(fd)
				}
				l--
			}
		} else if i == o {
			mid = l
			l++
			fd = 1
		} else {
			if target[i] == byte('+') {
				fd = 1
				l++
			} else if target[i] == byte('-') {
				fd = -1
				l++
			} else if target[i] >= byte('0') && target[i] <= byte('9') {
				if a[l] == 0 {
					a[l] = (int(target[i]) - int('0')) * fd
				} else {
					a[l] = a[l]*10 + (int(target[i])-int('0'))*fd
				}
			} else if target[i] >= byte('a') && target[i] <= byte('z') {
				p = string(target[i])
				if a[l] != 0 {
					x -= float64(a[l])
					a[l] = 0
				} else {
					x -= float64(fd)
				}
				l--
			}
		}
	}
	for j := 0; j <= l; j++ {
		if j <= mid {
			sum -= float64(a[j])
		} else {
			sum += float64(a[j])
		}
	}
	if float64(sum)/x == 0 {
		fmt.Printf("%s=0.000", p)
	} else {
		fmt.Printf("%s=%.3f", p, sum/x)
	}
}
