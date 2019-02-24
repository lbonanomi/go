package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var list1  map[string]int			    // Vars must be defined
var list2 map[string]int			    //
var longlist map[string]int		    //
						                      //
func main() {					            //
						                      //
	list1 = make(map[string]int)		// and instanced seperately
	list2 = make(map[string]int)		//
	longlist = make(map[string]int)


	for argc,argv := range(os.Args) {				             // foreach argv
		if argc > 0 {						                           // skip argv[0]
			if _,x := os.Stat(argv); !os.IsNotExist(x) {	   // if existing file
				dat,_ := ioutil.ReadFile(argv)		             // read into memory

				arr := strings.Split(string(dat), "\n")

				for _,line := range(arr) {
					if argc == 1 {
						v := string(line)
						list1[v] = 1
						longlist[v] = 1
					}

					if argc == 2 {
						v := string(line)
						list2[v] = 1
						longlist[v] = 1
					}
				}
			}
		}
	}

	intersection := 0
	union := 0

	for item := range(longlist) {
		_,in_one := list1[item]
		_,in_two := list2[item]

		if in_one && in_two {
			intersection += 1
		} 

		union += 1
	}

	jack := float32(intersection) / float32(union)

	fmt.Println(jack)
}
