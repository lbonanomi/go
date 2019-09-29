package main

import (
        "fmt"
        "io/ioutil"
        "os"
        "strings"
        "github.com/agnivade/levenshtein"
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

func find_possibles(have string, known []string)(candidates []string) {
        for _, known_word := range(known) {
                distance := levenshtein.ComputeDistance(known_word, have)

                if distance <= 1 {
                        if string(known_word[0]) == string(have[0]) {
                                candidates = append(candidates, known_word)
                        }
                }
        }
        return(candidates)
}

func main() {
        words := known_words()

        input,_ := ioutil.ReadFile(os.Args[1])

        for _,word := range(strings.Fields(string(input))) {

                word = strings.ReplaceAll(word, ".", "")
                if !find_word(word, words) {
                        fmt.Println("Couldn't find", word, "how-about: ", find_possibles(word, words))
                }
        }
}
