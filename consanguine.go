HAPROXY: cat cozy.go
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

    var pen []string

    arguments := os.Args[1:]

    for _,topitem := range(arguments) {
    for _,bottomitem := range(arguments) {
        if topitem != bottomitem {
        topType,_ := os.Stat(topitem)
        botType,_ := os.Stat(bottomitem)

        if topType.Mode().IsRegular() && botType.Mode().IsRegular() {
        aDat,_ := ioutil.ReadFile(topitem)
        a := strings.Fields(string(aDat)) //, "\n")

        bDat,_ := ioutil.ReadFile(bottomitem)
        b := strings.Fields(string(bDat)) //, "\n")

        cozy := int(cosine(a, b) * 100)

        if cozy > 85 {
        //fmt.Println(topitem, bottomitem)
        //pen = append(pen, topitem, bottomitem)

        topfound := 0
        botfound := 0

        for _,penned := range(pen) {
            if topitem == penned {
                topfound = 1
            }

            if bottomitem == penned {
                botfound = 1
            }
        }

        if (topfound == 0) {
            pen = append(pen, topitem)
        }

        if (botfound == 0) {
            pen = append(pen, bottomitem)
        }
        }
        }
        }
    }
    new := arguments[1:]
    arguments = new
    }

    fmt.Println(pen)
}

HAPROXY: