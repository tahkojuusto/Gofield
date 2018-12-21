package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadConfigFromFile reads the configuration (charges, measure points) from a file.
func ReadConfigFromFile(path string) ([]*Vector, []*PointCharge) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln("Failed to open file.")
	}

	var rVecs []*Vector
	var Qs []*PointCharge

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
		} else {
			log.Fatalln("Failed to parse config file: Invalid resource type.")
		}
	}

	return rVecs, Qs
}
