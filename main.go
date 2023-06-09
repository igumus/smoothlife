package main

import (
	"flag"
	"math"
	"math/rand"
	"time"
)

const (
	WIDTH  = 50
	HEIGHT = 50
)

var (
	ra      float64
	b1      float64
	b2      float64
	d1      float64
	d2      float64
	alpha_n float64
	alpha_m float64
	dt      float64
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

func nextStep(grid *[HEIGHT][WIDTH]float64, paperDiff bool) {
	var (
		ra_square float64 = ra * ra
		ri        float64 = ra / 3.0
		ri_square float64 = ri * ri
		x         int     = 0
		y         int     = 0
		m         float64 = 0.0
		M         float64 = 0.0
		n         float64 = 0.0
		N         float64 = 0.0
		diff      [HEIGHT][WIDTH]float64
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

			if paperDiff {
				diff[cy][cx] = dt * (2.0*s(n, m) - 1.0) * grid[cy][cx]
			} else {
				diff[cy][cx] = dt * (2.0*s(n, m) - 1.0)
			}
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
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(0)

	flagRa := flag.Float64("ra", 11.0, "")
	flagDt := flag.Float64("dt", 0.05, "")
	flagB1 := flag.Float64("b1", 0.278, "")
	flagB2 := flag.Float64("b2", 0.365, "")
	flagD1 := flag.Float64("d1", 0.267, "")
	flagD2 := flag.Float64("d2", 0.445, "")
	flagAlphaN := flag.Float64("alpha-n", 0.028, "")
	flagAlphaM := flag.Float64("alpha-m", 0.147, "")
	flagInt := flag.Duration("interval", 300*time.Millisecond, "")
	flagStep := flag.Int("step", 200, "")
	flagPaperDiff := flag.Bool("with-paper-diff", true, "")
	flag.Parse()

	grid := [HEIGHT][WIDTH]float64{}
	initGrid(&grid)

	dt = *flagDt
	ra = *flagRa
	b1 = *flagB1
	b2 = *flagB2
	d1 = *flagD1
	d2 = *flagD2
	alpha_m = *flagAlphaM
	alpha_n = *flagAlphaN

	count := 0
	for {
		displayGrid(&grid)
		nextStep(&grid, *flagPaperDiff)
		count += 1
		if count > *flagStep {
			break
		}
		time.Sleep(*flagInt)
		clearScreen()
	}
}
