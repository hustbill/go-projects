// Name: webScrapter.go
// Use this code to fetch a web page by URL
package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "github.com/opesun/goquery"
    "os"
    "strings"
)


// use Goquery to fetch product-name from html
func fetchProductName(url string){
    p, err := goquery.ParseUrl(url)
    if err != nil {
        panic(err)
    } else {
        pTitle := p.Find("title").Text()// fetch the content of title
        fmt.Println(pTitle)

        productList := p.Find(".product-name")
        for i := 0; i < productList.Length(); i++ {
            product:= productList.Eq(i).Text()
            fmt.Println(product)
        }
     
    }
}


// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
    // Iterate over all of the Token's attributes until we find an "href"
    for _, a := range t.Attr {
        if a.Key == "href" {
            href = a.Val
            ok = true
        }
    }
    return
}


// Extract all shop related links from a given webpage
func crawl(url string, ch chan string, chFinished chan bool) {
    resp, err := http.Get(url)

    defer func() {
        // Notify that we're done after this function
        chFinished <- true
    }()

    if err != nil {
        fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
        return
    }

    b := resp.Body
    defer b.Close() // close Body when the function returns

    z := html.NewTokenizer(b)

    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:
            // End of the document, we're done
            return
        case tt == html.StartTagToken:
            t := z.Token()

            // Check if the token is an <a> tag
            isAnchor := t.Data == "a"
            if !isAnchor {
                continue
            }

            // Extract the href value, if there is one
            ok, url := getHref(t)
            if !ok {
                continue
            }
            

            //Make sure the url includes in "shop"
            hasProto := strings.Count(url, "shop") > 0
            if hasProto {
                ch <- url
            }
        }
    }
}


func main() {
    //testUrl := "http://www.stelladot.com/shop/en_us/jewelry/rings"
    //fetchProductName(testUrl)
   // fetchProductName(os.Args[1])
   
   foundUrls := make(map[string]bool)
   seedUrls := os.Args[1:]

    // Channels
    chUrls := make(chan string)
    chFinished := make(chan bool) 

    // Kick off the crawl process (concurrently)
    for _, url := range seedUrls {
        go crawl(url, chUrls, chFinished)
    }

    // Subscribe to both channels
    for c := 0; c < len(seedUrls); {
        select {
        case url := <-chUrls:
            foundUrls[url] = true
        case <-chFinished:
            c++
        }
    }

    // We're done! Print the URLs and Product-name ...
    fmt.Println("\nFound", len(foundUrls), "unique urls:\n")

    for url, _ := range foundUrls {
        
        hasProto := strings.Index(url, "http") == 0
        
        if !hasProto {
            url = "http://www.stelladot.com" + url
        }
        
        fmt.Println(" - " + url)
        
        fetchProductName(url)  
    }
    
    close(chUrls)
    
}


// Reference:
// (1)  http://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
// (2)  github.com/opesun/goquery