package main

import (
	"fmt"
	"net/url"
)

func main() {
	testUrl := "https://example.com/todos/1?test1=1&test2=2"

	parsedUrl, err := url.Parse(testUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Parsed URL:", parsedUrl)
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Query:", parsedUrl.RawQuery)
	
	parsedUrl.RawQuery += "&test3=3"
	fmt.Println("Updated Query:", parsedUrl.RawQuery)
}
