package main

import (
        "fmt"
        "io/ioutil"
        "math"
        "os"
        "sort"
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
        sort.Strings(arguments)

        var pairings map[string]string
        pairings = make(map[string]string)

        for _, topitem := range arguments {
                var matched_files []string

                matched_files = append(matched_files, topitem)

                for _, bottomitem := range arguments {
                        if topitem != bottomitem {
                                topType, _ := os.Stat(topitem)
                                botType, _ := os.Stat(bottomitem)

                                if topType.Mode().IsRegular() && botType.Mode().IsRegular() {
                                        aDat, _ := ioutil.ReadFile(topitem)
                                        a := strings.Fields(string(aDat))

                                        bDat, _ := ioutil.ReadFile(bottomitem)
                                        b := strings.Fields(string(bDat))

                                        if int(cosine(a, b)*100) > 80 {
                                                pairings[topitem] = bottomitem
                                        }
                                }
                        }
                }
        }

        var chatter []string

        // Group individual pairs into larger sets
        //

        for _, index := range pairings {
                var groups []string

                groups = append(groups, "")

                for a, b := range pairings {
                        if b == index && groups[len(groups)-1] != a {
                                groups = append(groups, a)
                        }
                        if a == index && groups[len(groups)-1] != b {
                                groups = append(groups, b)
                        }

                        sort.Strings(groups)
                }

                groups = append(groups, index)

                sort.Strings(groups)

                group := strings.Join(groups, " ")

                // This process is repetitive
                //

                chattering := 0

                for _, repeated := range chatter {
                        if strings.Contains(repeated, group) {
                                chattering = 1
                        }
                }

                chatter = append(chatter, group)

                if chattering == 0 {
                        fmt.Println(group)
                }
        }
}
