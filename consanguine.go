package main

import (
        "fmt"
        "math"
        "io/ioutil"
        "os"
        "strings"
)

func vectorize(a []string) (words map[string]int) {
        // Convert a string into a map of words:word-counts

        var returned_words map[string]int
        returned_words = make(map[string]int)

        for _,word := range(a) {
                aCount := 0

                for _,subA := range(a) {
                        if subA == word {
                                aCount = aCount + 1
                        }
                }
                returned_words[word] = aCount
                returned_words[word] = aCount
        }

        words = returned_words

        return
}

func cosine(a, b []string) (cosineSimilarity float64) {
        aa := vectorize(a)
        bb := vectorize(b)

        // All of this to make a set...
        //

        var keys []string

        for k := range aa {
                pf := 1

                for _,y := range(keys) {
                        if k == y {
                                pf = 0
                        }
                }

                if pf == 1 {
                        keys = append(keys, k)
                }
        }

        for k := range bb {
                pf := 1

                for _,y := range(keys) {
                        if k == y {
                                pf = 0
                        }
                }

                if pf == 1 {
                        keys = append(keys, k)
                }
        }

        // Here's our set

        numerator := 0
        denominatorA := 0.0
        denominatorB := 0.0

        for _,y := range(keys) {
                powerA := math.Pow(float64(aa[y]), 2)
                powerB  := math.Pow(float64(bb[y]), 2)

                denominatorA = denominatorA + powerA
                denominatorB = denominatorB + powerB

                numerator = numerator + (aa[y] * bb[y])
        }

        denominator := math.Sqrt(denominatorA) * math.Sqrt(denominatorB)

        cosineSimilarity = float64(numerator) / denominator

        return
}

func main() {
        _,a_presence := os.Stat(os.Args[1])
        if a_presence != nil {
                fmt.Println("No such file: ", os.Args[1])
                os.Exit(1)
        }

        _,b_presence := os.Stat(os.Args[2])
        if b_presence != nil {
                fmt.Println("No such file: ", os.Args[2])
                os.Exit(1)
        }

        aDat,_ := ioutil.ReadFile(os.Args[1])
        a := strings.Fields(strings.ToLower(string(aDat)))

        bDat,_ := ioutil.ReadFile(os.Args[2])
        b := strings.Fields(strings.ToLower(string(bDat)))

        cozy := int(cosine(a, b) * 100)
        fmt.Println(cozy)
}
