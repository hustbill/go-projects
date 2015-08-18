// Name: webScrapter.go
// Use this code to fetch a web page by URL
package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "github.com/opesun/goquery"
    "os"
   //"strings"
)



// Helper function to pull the product-name attribute from a Token
func getProductName(t html.Token) (ok bool, href string) {
    // Iterate over all of the Token's attributes until we find an "strong class"
    for _, a := range t.Attr {
        if a.Key == "strong class" {
            href = a.Val
            ok = true
        }
    }
    
    return 
}

// use Goquery to fetech product-name from html
func useGoQuery(url string){
    p, err := goquery.ParseUrl(url)
    if err != nil {
        panic(err)
    } else {
        pTitle := p.Find("title").Text()//直接提取title的内容
        fmt.Println(pTitle)

        productList := p.Find(".product-name")
        for i := 0; i < productList.Length(); i++ {
            product:= productList.Eq(i).Text()
            fmt.Println(product)
        }
     
    }
}



// Make an HTTP request //
func makeHttpReq(url string) {
    resp, _ := http.Get(url)
    // bytes, _ := ioutil.ReadAll(resp.Body)
 //
 //
 //
 //    fmt.Println("HTML:\n\n", string(bytes))
    resp.Body.Close()
}


// Extract all http**  links from a given webpage
func crawl(url string, ch chan string, chFinished chan bool) {
    fmt.Println(url)
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
        case  tt == html.ErrorToken:
            // End of the document, we're done
            return
        case tt == html.StartTagToken:
            t := z.Token()
            
            // Check if the token is an <strong class> tag
            isAnchor := t.Data == "strong class"
            if !isAnchor {
                continue
            }
            
            // Extract the href value, if there is one
            ok, url := getProductName(t)
            fmt.Println(url)
            if !ok {
                continue
            }
            
            ch <- url
            
            // // Make sure the url begines in http**
   //          hasProto := strings.Index(url, "http") == 0
   //          if hasProto {
   //              ch <- url
   //          }
            
        }
    }
}




func main() {
    testUrl := "http://www.stelladot.com/shop/en_us/jewelry/rings"
    makeHttpReq(testUrl)
    useGoQuery(testUrl)
    
    foundUrls := make(map[string]bool)
    seedUrls := os.Args[1:]
    
    //  Channels 
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
    
    
    // We're done! Print the results...
    
    fmt.Println("\nFound", len(foundUrls), "unique urls:\n")
    
    for url, _ := range foundUrls {
        fmt.Println(" - " + url)
    }
    close(chUrls)
}