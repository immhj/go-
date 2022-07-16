package main

import "fmt"

type shiti struct {
	x   int
	y   int
	dir int
}

var m [15][15]int
var cow shiti
var farmer shiti
var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}
var mark [11][11][4][11][11][4]int

func move() {
	fx := farmer.x + dx[farmer.dir]
	fy := farmer.y + dy[farmer.dir]
	cx := cow.x + dx[cow.dir]
	cy := cow.y + dy[cow.dir]
	if m[fx][fy] == 1 {
		farmer.x = fx
		farmer.y = fy
	} else {
		farmer.dir = (farmer.dir + 1) % 4
	}
	if m[cx][cy] == 1 {
		cow.x = cx
		cow.y = cy
	} else {
		cow.dir = (cow.dir + 1) % 4
	}
}
func main() {
	var c byte
	var time int
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 12; j++ {
			fmt.Scanf("%c", &c)
			if c == '.' {
				m[i][j] = 1
			} else if c == 'C' {
				m[i][j] = 1
				cow.x = i
				cow.y = j
				cow.dir = 0
			} else if c == 'F' {
				m[i][j] = 1
				farmer.x = i
				farmer.y = j
				farmer.dir = 0
			}
		}
	}
	for farmer.x != cow.x || farmer.y != cow.y {
		time++
		move()
		if mark[farmer.x][farmer.y][farmer.dir][cow.x][cow.y][cow.dir] == 1 {
			fmt.Println(0)
			return
		} else {
			mark[farmer.x][farmer.y][farmer.dir][cow.x][cow.y][cow.dir] = 1
		}
	}
	fmt.Println(time)
}
