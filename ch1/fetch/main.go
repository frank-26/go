// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"io"
	"strings"
	"fmt"
	// "io/ioutil"
	"net/http"
	"os"
)

/*

net: 网络收发信息，还可以建立更底层的网络连接，编写服务器程序
*/

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url,"http://"){
			url="http://"+url
		}
		resp, err :=  http.Get(url)
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		b,err := io.Copy(os.Stdout,resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", b)
		fmt.Printf("%s", resp.Status)
	}
}

//!-
