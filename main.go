package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalln("Usage: ./gofield <config-file-path>")
	}

	rVecs, Qs := ReadConfigFromFile(os.Args[1])
	log.Println("### INPUTS ###")
	for i, Q := range Qs {
		log.Printf("r(Q%d) = %s", i+1, Q.rVec)
	}

	log.Println("### OUTPUTS ###")
	for _, rVec := range rVecs {
		EtotFn := GetEtotFn(Qs)
		Etot := EtotFn(rVec)

		log.Printf("E(%s) = %s", rVec, Etot)
		log.Printf("|E| = %v", PrettifyResult(Magnitude(Etot)))
	}
}
