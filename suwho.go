package main

import (
        "fmt"
        "os"
        "os/user"
        "path/filepath"
        "strconv"
        "syscall"
)

func main() {
        // Get realpath of stdin
        tty, _ := filepath.EvalSymlinks("/dev/stdin")

        // Jump through flaming hoop to get TTY-owner's UID
        statHandle, _ := os.Stat(tty)
        stat := statHandle.Sys().(*syscall.Stat_t)

        if int(stat.Uid) != int(syscall.Getuid()) {
                userID := strconv.Itoa(int(stat.Uid))

                userName,_ := user.LookupId(userID)

                fmt.Println("su-ed (", userName.Username, ")")
        } else {
                fmt.Println("plain")
        }
}
