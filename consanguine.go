package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)


func cosine(a, b []string) (cosineSimilarity float64) {
	//
	// Calculate intersection
	//

	var intersection map[string]int
	intersection = make(map[string]int)

	var uniqA map[string]int
	uniqA = make(map[string]int)

	var uniqB map[string]int
	uniqB = make(map[string]int)


	for _,z := range(a) {
		intersection[z] = intersection[z] + 1
		uniqA[z] = uniqA[z] + 1
	}

	for _,z := range(b) {
		intersection[z] = intersection[z] + 1
		uniqB[z] = uniqB[z] + 1
	}


	// Count duplicated vaues as 'numerator'
	//

	numerator := 0

	for _,z := range(intersection) {
		if z > 1 {
			numerator = numerator + 1
		}
	}

	aSum := 0
	bSum := 0

	for _,cat := range(uniqA) {
		powered := math.Pow(float64(cat), 2)
		aSum += int(powered)
	}

	for _,cat := range(uniqB) {
		powered := math.Pow(float64(cat), 2)
		bSum += int(powered)
	}

	aFloat := float64(aSum)
	bFloat := float64(bSum)

	denominator := math.Sqrt(aFloat) * math.Sqrt(bFloat)

	cosineSimilarity = float64(numerator) / denominator

	return
}

func main() {
	aDat,_ := ioutil.ReadFile(os.Args[1])
	a := strings.Split(string(aDat), "\n")
	
	bDat,_ := ioutil.ReadFile(os.Args[2])
	b := strings.Split(string(bDat), "\n")

	cozy := cosine(a, b)

	if cozy > 0.75 {
		fmt.Println(cozy)
	}
}
