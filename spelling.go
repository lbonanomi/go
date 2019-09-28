package main

import (
        "fmt"
        "io/ioutil"
        "strings"
        //"github.com/agnivade/levenshtein"
)

func known_words()(dict []string) {
        words, _ := ioutil.ReadFile("/usr/share/dict/words")
        dict = strings.Fields(string(words))
        return
}

func find_word(wanted string, known []string)(gotcha bool) {
        for _, known_word := range(known) {
                wanted = strings.ToLower(wanted)

                if known_word == wanted {
                        return(true)
                }
        }
        return(false)
}

func main() {
        //distance := levenshtein.ComputeDistance(a, b)

        words := known_words()

        input,_ := ioutil.ReadFile("/home/ec2-user/loren")

        for _,word := range(strings.Fields(string(input))) {
                if !find_word(word, words) {
                        fmt.Println("DIST", word)
                }
        }
}
