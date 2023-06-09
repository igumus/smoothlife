package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	WIDTH  = 50
	HEIGHT = 50
)

var (
	ra        float64 = 11.0
	ra_square float64 = ra * ra
	ri        float64 = ra / 3.0
	ri_square float64 = ri * ri

	b1      float64 = 0.278
	b2      float64 = 0.365
	d1      float64 = 0.267
	d2      float64 = 0.445
	alpha_n float64 = 0.028
	alpha_m float64 = 0.147
)

func sigma(x, a, alpha float64) float64 {
	return 1.0 / (1.0 + math.Pow(math.E, -(x-a)*4.0/alpha))
}

func sigma_m(x, y, m float64) float64 {
	v := sigma(m, 0.5, alpha_m)
	return x*(1.0-v) + y*v
}

func sigma_n(x, a, b float64) float64 {
	return sigma(x, a, alpha_n) * (1.0 - sigma(x, b, alpha_n))
}

func s(n, m float64) float64 {
	return sigma_n(n, sigma_m(b1, d1, m), sigma_m(b2, d2, m))
}

func nextStep(grid *[HEIGHT][WIDTH]float64, dt float64) {
	var (
		x    int     = 0
		y    int     = 0
		m    float64 = 0.0
		M    float64 = 0.0
		n    float64 = 0.0
		N    float64 = 0.0
		diff [HEIGHT][WIDTH]float64
	)

	// computing diff
	for cy := 0; cy < HEIGHT; cy += 1 {
		for cx := 0; cx < WIDTH; cx += 1 {
			for dy := -(ra - 1); dy < ra; dy += 1 {
				for dx := -(ra - 1); dx < ra; dx += 1 {
					x = emod(cx+int(dx), WIDTH)
					y = emod(cy+int(dy), HEIGHT)
					total := dx*dx + dy*dy
					if total <= ri_square {
						m += grid[y][x]
						M += 1.0
					} else if total <= ra_square {
						n += grid[y][x]
						N += 1.0
					}
				}
			}

			m /= M
			n /= N
			// diff[cy][cx] = dt * (2.0*s(n, m) - 1.0) * grid[cy][cx]
			diff[cy][cx] = dt * (2.0*s(n, m) - 1.0)
			m = 0.0
			M = 0.0
			n = 0.0
			N = 0.0
		}
	}

	// applying diff
	for cy := 0; cy < HEIGHT; cy += 1 {
		for cx := 0; cx < WIDTH; cx += 1 {
			grid[cy][cx] += diff[cy][cx]
			clamp(&grid[cy][cx], 0.0, 1.0)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flagDt := flag.Float64("dt", 0.05, "dt value")
	flagInt := flag.Duration("interval", 300*time.Millisecond, "interval between each step")
	flag.Parse()

	grid := [HEIGHT][WIDTH]float64{}
	initGrid(&grid)
	count := 0
	for {
		fmt.Printf("--[%d]-------------------------------------------------------------\n", count)
		displayGrid(&grid)
		nextStep(&grid, *flagDt)
		fmt.Printf("--[%d]-------------------------------------------------------------\n", count)
		time.Sleep(*flagInt)
		clearScreen()
		count += 1
	}
}
