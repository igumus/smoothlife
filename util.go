package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

const level = " .-=coaA@#"

func randFloat64() float64 {
	return float64(rand.Int63()) / (1 << 63)
}

func emod(a, b int) int {
	return (a%b + b) % b
}

func clamp(x *float64, low float64, high float64) {
	if *x < low {
		*x = low
	} else if *x > high {
		*x = high
	}
}

func initGrid(grid *[HEIGHT][WIDTH]float64) {
	for dy := 0; dy < HEIGHT; dy++ {
		for dx := 0; dx < WIDTH; dx++ {
			grid[dy][dx] = randFloat64()
		}
	}
}

func displayGrid(grid *[HEIGHT][WIDTH]float64) {
	size := len(level)
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			value := grid[y][x]
			step := int(value) * (size - 1)
			v := level[emod(step, size)]
			fmt.Printf("%c", v)
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
