package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	path, err := filepath.Abs("./DayOne/PartTwo/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	t := float64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		numb, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatal(err)
		}

		f, total := calculate(numb, 0)
		log.Print(f, total)
		t = t + total
	}

	log.Print(t)
	log.Printf("%.6f\n", t)
}

func calculate(f float64, t float64) (float64, float64) {
	x := math.Floor(f/3) - 2
	if x <= 0 {
		return f, t
	}

	t = t + x
	return calculate(x, t)
}
