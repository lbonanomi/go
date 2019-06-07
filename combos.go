package main

import (
    "fmt"
    "os"
)

func main() {
    arguments := os.Args[1:]

    for _,topitem := range(arguments) {
        for _,bottomitem := range(arguments) {

            if topitem != bottomitem {
                fmt.Println(topitem, bottomitem)
            }
        }
        new := arguments[1:]
        arguments = new
    }
}
