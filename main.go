package main

import (
	"fmt"
	"os"
	"net/http"
)

func main () {
	// check that the user supplied a url
	if len(os.Args) < 2 {
		fmt.Println("Please supply a url to check")
		return
	}

	url := os.Args[1]
	fmt.Println("Checking " + url)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("User-Agent", "() { :;}; echo \"Warning: Server Vulnerable\"")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Header["Warning"] != nil && resp.Header["Warning"][0] == "Server Vulnerable" {
		fmt.Println("warning: Your server is vulnerable!\nFor more information about this bug and steps to take to fix it visit https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2014-6271")
	} else {
		fmt.Println("Your server does not seem to be vulnerable but it may still be.\nVisit the following links for more information and other tests to be sure\nhttp://bashsmash.ccsir.org/\nhttp://www.shellshocktest.com/")
	}
}
