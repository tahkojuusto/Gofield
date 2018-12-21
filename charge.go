package main

import (
	"fmt"
	"math"
)

// ElectricConstant describes the electric constant e.
const ElectricConstant float64 = 8.854187817 * 10e-12

// ElementaryCharge is the elementary charge e.
const ElementaryCharge float64 = 1.602176634 * 10e-19

// PointCharge describes a Vector charge in 2d space.
type PointCharge struct {
	rVec Vector
	q    float64
}

func (Q PointCharge) String() string {
	return fmt.Sprintf("(%s): %vC", Q.rVec, Q.q)
}

// GetEtotFn returns the total net electric field vector function Etot(rVec).
func GetEtotFn(Qs []*PointCharge) func(rVec *Vector) *Vector {
	return func(rVec *Vector) *Vector {
		E := &Vector{0, 0}
		for _, Q := range Qs {
			Efn := Q.GetEFn()
			E = Add(E, Efn(rVec))
		}

		return E
	}
}

// GetEFn returns the electric field vector function E(rVec).
func (Q *PointCharge) GetEFn() func(rVec *Vector) *Vector {
	k := 1 / (4 * math.Pi * ElectricConstant)

	return func(rVec *Vector) *Vector {
		dVec := Substract(&Q.rVec, rVec)
		dVec0 := Normalize(dVec)
		distance := Magnitude(dVec)
		E := k * Q.q / math.Pow(distance, 2)
		return Scale(E, dVec0)
	}
}

// PrettifyResult will format the electric field to right unit.
func PrettifyResult(result float64) string {
	fResult := result
	unit := "V/m"

	if result < 1e-3 {
		fResult *= 1e3
		unit = "mV/m"
	} else if result > 1e3 && result < 1e6 {
		fResult *= 1e-3
		unit = "kV/m"
	} else if result >= 1e6 && result < 1e9 {
		fResult *= 1e-6
		unit = "MV/m"
	} else if result >= 1e9 {
		fResult *= 1e-9
		unit = "GV/m"
	}

	return fmt.Sprintf("%v [%s]", fResult, unit)
}
