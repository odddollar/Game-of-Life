package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const probabilityOfInitialAlive = 80

type board struct {
	b    [][]bool
	w, h int
}

func newBoard(w, h int) board {
	b := make([][]bool, h)

	for i := 0; i < h; i++ {
		b[i] = make([]bool, w)
	}

	return board{b: b, w: w, h: h}
}

func (b *board) initialiseBoard() {
	for y := 0; y < b.h; y++ {
		for x := 0; x < b.w; x++ {
			b.b[y][x] = rand.Intn(100) > probabilityOfInitialAlive
		}
	}
}

func showBoard(b board) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	for y := 0; y < b.h; y++ {
		for x := 0; x < b.w; x++ {
			if !b.b[y][x] {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	field := newBoard(30, 10)

	field.initialiseBoard()

	showBoard(field)
}
