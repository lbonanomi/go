package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func curl_u(cmdline string)(after string) {
    curl := regexp.MustCompile(`:\S+?\s+\b`)
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
