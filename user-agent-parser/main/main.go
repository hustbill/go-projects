package main

import (
	"fmt"
	"log"

	"github.com/ua-parser/uap-go/uaparser"
)

func main() {
	uagent := "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-80) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16 Silk-Accelerated=true"

	parser, err := uaparser.New("./regexes.yaml")
	if err != nil {
		log.Fatal(err)
	}

	client := parser.Parse(uagent)

	fmt.Println(client.UserAgent.Family) // "Amazon Silk"
	fmt.Println(client.UserAgent.Major)  // "1"
	fmt.Println(client.UserAgent.Minor)  // "1"
	fmt.Println(client.UserAgent.Patch)  // "0-80"
	fmt.Println(client.Os.Family)        // "Android"
	fmt.Println(client.Os.Major)         // ""
	fmt.Println(client.Os.Minor)         // ""
	fmt.Println(client.Os.Patch)         // ""
	fmt.Println(client.Os.PatchMinor)    // ""
	fmt.Println(client.Device.Family)    // "Kindle Fire"
}
