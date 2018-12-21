package main

import (
	"fmt"
	"math"
)

// Vector describes (x, y) Vector in 2d space.
type Vector struct {
	x float64
	y float64
}

func (vec Vector) String() string {
	return fmt.Sprintf("%.02vi + %.02vj", vec.x, vec.y)
}

// Multiply defines the dot product v1 * v2.
func Multiply(vec1 *Vector, vec2 *Vector) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y
}

// Scale will multiply the vector by a scalar a * v.
func Scale(a float64, p *Vector) *Vector {
	return &Vector{a * p.x, a * p.y}
}

// Add defines the vector summation v1 + v2.
func Add(vec1 *Vector, vec2 *Vector) *Vector {
	return &Vector{vec1.x + vec2.x, vec1.y + vec2.y}
}

// Substract defines the vector substraction v1 + (-v2).
func Substract(vec1 *Vector, vec2 *Vector) *Vector {
	return Add(vec1, Scale(-1, vec2))
}

// Distance describes scalar norm of two vectors |v2 - v1|.
func Distance(vec1 *Vector, vec2 *Vector) float64 {
	return math.Sqrt(math.Pow(vec1.x-vec2.x, 2) + math.Pow(vec1.y-vec2.y, 2))
}

// Magnitude returns the length of the vector |v|.
func Magnitude(vec *Vector) float64 {
	return math.Sqrt(math.Pow(vec.x, 2) + math.Pow(vec.y, 2))
}

// Normalize will scale the vector to be a unit vector v0 = v/|v|.
func Normalize(vec *Vector) *Vector {
	return Scale(1/Magnitude(vec), vec)
}
