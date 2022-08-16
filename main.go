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
	// create and return new board struct
	b := make([][]bool, h)

	for i := 0; i < h; i++ {
		b[i] = make([]bool, w)
	}

	return board{b: b, w: w, h: h}
}

func (b *board) initialisePositions() {
	// set all positions of board to random values
	for y := 0; y < b.h; y++ {
		for x := 0; x < b.w; x++ {
			b.b[y][x] = rand.Intn(100) > probabilityOfInitialAlive
		}
	}
}

func (b *board) isAlive(x, y int) bool {
	// determine alive value at position with wrapping
	x += b.w
	x %= b.w
	y += b.h
	y %= b.h
	return b.b[y][x]
}

func (b *board) numNeighbours(x, y int) int {
	neighbours := 0

	// iterate through all spaces around x, y co-ordinates
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// if current iteration position isn't x, y and position is alive (taking in account wrapping)
			// then increment neighbours
			if (i != 0 || j != 0) && b.isAlive(x+i, y+j) {
				neighbours++
			}
		}
	}

	return neighbours
}

func (b *board) step() {
	newB := newBoard(b.w, b.h)
}

func (b *board) show() {
	// clear screen
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	// print board characters based on alive values
	for y := 0; y < b.h; y++ {
		for x := 0; x < b.w; x++ {
			if b.isAlive(x, y) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// set random seed
	rand.Seed(time.Now().UnixNano())

	// create board, initialise to random values and display
	field := newBoard(30, 10)
	field.initialisePositions()
	field.show()

	// main loop
	// steps forward generation
	// displays board
	// waits period before running again
	for {
		field.step()
		field.show()
		time.Sleep(time.Second / 2)
	}
}
