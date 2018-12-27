package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadConfigFromFile reads the configuration (charges, measure points) from a file.
func ReadConfigFromFile(path string) ([]*Vector, []*PointCharge, *Vector, *Vector, int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln("Failed to open file.")
	}

	var rVecs []*Vector
	var Qs []*PointCharge
	var minVec *Vector
	var maxVec *Vector
	var N int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		x, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			log.Fatalln("Failed to parse config file: Could not parse x value.")
		}
		y, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			log.Fatalln("Failed to parse config file: Could not parse y value.")
		}

		if line[0] == "Q" {
			q, err := strconv.ParseFloat(line[3], 64)
			if err != nil {
				log.Fatalln("Failed to parse config file: Could not parse charge value.")
			}
			Qs = append(Qs, &PointCharge{Vector{x, y}, q})
		} else if line[0] == "P" {
			rVecs = append(rVecs, &Vector{x, y})
		} else if line[0] == "D" {
			xMax, err := strconv.ParseFloat(line[3], 64)
			if err != nil {
				log.Fatalln("Failed to parse config file: Could not parse max x value.")
			}
			yMax, err := strconv.ParseFloat(line[4], 64)
			if err != nil {
				log.Fatalln("Failed to parse config file: Could not parse max y value.")
			}
			N0, err := strconv.ParseInt(line[5], 10, 0)
			if err != nil {
				log.Fatalln("Failed to parse config file: Could not parse N value.")
			}
			minVec = &Vector{x, y}
			maxVec = &Vector{xMax, yMax}
			N = int(N0)
		} else {
			log.Fatalln("Failed to parse config file: Invalid resource type.")
		}
	}

	return rVecs, Qs, minVec, maxVec, N
}
