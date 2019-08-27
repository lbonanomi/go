package main

import (
        "fmt"
        "io/ioutil"
        "math"
        "os"
        "strings"
)

func vectorize(a []string) (words map[string]int) {
        // Convert a string into a map of words:word-counts

        var returned_words map[string]int
        returned_words = make(map[string]int)

        for _, word := range a {
                aCount := 0

                for _, subA := range a {
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

                for _, y := range keys {
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

                for _, y := range keys {
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

        for _, y := range keys {
                powerA := math.Pow(float64(aa[y]), 2)
                powerB := math.Pow(float64(bb[y]), 2)

                denominatorA = denominatorA + powerA
                denominatorB = denominatorB + powerB

                numerator = numerator + (aa[y] * bb[y])
        }

        denominator := math.Sqrt(denominatorA) * math.Sqrt(denominatorB)

        cosineSimilarity = float64(numerator) / denominator

        return
}

func main() {

        arguments := os.Args[1:]

        for _, topitem := range arguments {

                var piggy []string

                //fmt.Print(topitem, " ")
                piggy = append(piggy, topitem)

                for _, bottomitem := range arguments {
                        if topitem != bottomitem {
                                topType, _ := os.Stat(topitem)
                                botType, _ := os.Stat(bottomitem)

                                if topType.Mode().IsRegular() && botType.Mode().IsRegular() {
                                        aDat, _ := ioutil.ReadFile(topitem)
                                        a := strings.Fields(string(aDat)) //, "\n")

                                        bDat, _ := ioutil.ReadFile(bottomitem)
                                        b := strings.Fields(string(bDat)) //, "\n")

                                        cozy := int(cosine(a, b) * 100)

                                        //fmt.Println(topitem, bottomitem, cozy)

                                        if cozy > 80 {
                                            //fmt.Println(topitem, bottomitem, cozy)
                                            fmt.Print(bottomitem, " ")
                                            piggy = append(piggy, bottomitem)
                                        }
                                }
                        }
                }
                new := arguments[1:]
                arguments = new

                if len(piggy) > 1 {
                    fmt.Println(strings.Join(piggy, " "), "has", len(piggy), "things")
                }
        }
}

