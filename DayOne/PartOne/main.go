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
	path, err := filepath.Abs("./DayOne/PartOne/input.txt")
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

		o := math.Floor(numb/3) - 2
		log.Print(o)
		t = t + o
	}

	log.Printf("%.6f\n", t)
}
