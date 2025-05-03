package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"slices"
)

func mounted()(mounts []string) {
	procMounts,_ := os.Open("/proc/mounts")
	defer procMounts.Close()

	scanner := bufio.NewScanner(procMounts)

	for scanner.Scan() {
		procLine := strings.Fields(scanner.Text())

		if(strings.HasPrefix(procLine[0], "/")) {
			mounts = append(mounts, procLine[1])
		}
	}
	return
}

func fstable()(mounts []string) {

	fstab,_ := os.Open("/etc/fstab")
	defer fstab.Close()

	scanner := bufio.NewScanner(fstab)

	for scanner.Scan() {
		fstabLine := strings.Fields(scanner.Text())

		if(!strings.HasPrefix(fstabLine[0], "#")) {
			if(strings.HasPrefix(fstabLine[1], "/")) {
				mounts = append(mounts, fstabLine[1])
			}
		}
	}
	return
}


func main() {
	for _,mount := range(mounted()) {
		if (!slices.Contains(fstable(), mount)) {
			fmt.Println(mount)
		}
	}
}
