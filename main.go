package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalln("Usage: ./gofield <config-file-path>")
	}

	rVecs, Qs, minVec, maxVec, N := ReadConfigFromFile(os.Args[1])
	for i, Q := range Qs {
		log.Printf("r(Q%d) = %s", i+1, Q.rVec)
	}

	EtotFn := GetEtotFn(Qs)
	for _, rVec := range rVecs {
		Etot := EtotFn(rVec)

		log.Printf("E(%s) = %s", rVec, Etot)
		log.Printf("|E(%s)| = %v", rVec, PrettifyResult(Magnitude(Etot)))
	}

	Draw(Qs, EtotFn, minVec, maxVec, N)
}
