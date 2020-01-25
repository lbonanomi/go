-bash-4.3$ cat *.go
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os/user"
    "os"
    "regexp"
    "strings"
)

func append_netrc(machine string, login string, password string)() {
        usr,_ := user.Current()
        netrc_file := string(usr.HomeDir) + "/.netrc"

        // *IF* there is no ~/.netrc, create one
        _,err := os.Open(netrc_file)

        if (err != nil) {
                f,_ := os.Create(netrc_file)
                f.WriteString("#")
                f.Close()
        } else {
                netrc_data,_ := ioutil.ReadFile(netrc_file)

                if !strings.Contains(string(netrc_data), machine) {
                        netrc_entry := "\nmachine\t" + machine + "\nlogin\t" + login + "\npassword\t" + password + "\n\n"

                        f,_ := os.OpenFile(netrc_file, os.O_APPEND|os.O_WRONLY, 0600)
                        f.WriteString(netrc_entry)
                        f.Close()
                }
        }
}

func curl_u(cmdline string)(after string) {
    curl := regexp.MustCompile(`:\S+?\s+\b`)

    token := strings.Trim(curl.FindString(cmdline), ":")

    if len(token) > 0 {
        // If a valid token is found, try and extract username and host
        // and append to ~/.netrc

        user_regex := regexp.MustCompile(`\S+?:\b`)
        user := strings.Trim(user_regex.FindString(cmdline), ":")

        url_regex := regexp.MustCompile(`:\/\/(\S+\.\w*)`)
        url := strings.Trim(url_regex.FindString(cmdline), ":/")

        append_netrc(url, user, token)
    }

    after = curl.ReplaceAllString(cmdline, ":REDACTED ")
    return
}

func https_creds(cmdline string)(after string) {
    pattern := regexp.MustCompile(`https://(\S+?):\S+?@`)
    after = pattern.ReplaceAllString(cmdline, "https://REDACTED:REDACTED@")
    return
}

func header_creds(cmdline string)(after string) {
    pattern := regexp.MustCompile(`(-H|--header)\s.*?(token|auth.*?)\s\S+?\s`)
    after = pattern.ReplaceAllString(cmdline, "-H AUTH_HEADER_REDACTED ")
    return
}


func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        text, _ := reader.ReadString('\n')

        if (text == "") {
            break
        }

        newtext := ""

        for _, word := range(strings.Fields(text)[1:]) {    // Remove history line number
            newtext = newtext + " " + word                  //
        }

        // Redact credential patterns
        //

        text = newtext

        text = curl_u(text)
        text = https_creds(text)
        text = header_creds(text)

        fmt.Println(text)
    }
}
-bash-4.3$
