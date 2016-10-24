package main

// https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
// Web scraper from above tutorial

import (
  "net/http"
  "golang.org/x/net/html"
  "fmt"
  "strings"
)

func main() {

  foundUrls := make(map[string]bool)
  seedUrls := [3]string{"https://www.golang.org", "https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html", "https://godoc.org/golang.org/x/net/html#TokenType"}

  // channels
  chUrls := make(chan string)
  chFinished := make(chan bool)

  // conurrently set off all seedUrls to be scraped
  for _, url := range seedUrls {
    go Crawl(url, chUrls, chFinished)
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

  // Done scraping, now to print the results
  fmt.Println("\nFound", len(foundUrls), "unique urls:\n")

  for url, _ := range foundUrls {
    fmt.Println(" - " + url)
  }

  close(chUrls)
}

// Function to find href attribute from a Token
func GetHref(t html.Token) (ok bool, href string) {
  //Iterates over the attributes in t to find "href"
  for _, a := range t.Attr {
    if a.Key == "href" {
      href = a.Val
      ok = true
      return
    }
  }

  return false, ""
}

func Crawl(url string, ch chan string, chFinished chan bool) {

  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
  // Parse HTML for Anchor Tags
  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

  resp, err := http.Get(url)

  defer func() {
    // Notify that we're done after this function
    chFinished <- true
  }()

  if err != nil {
    fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
    return
  }

  defer resp.Body.Close()

  htmlTokenizer := html.NewTokenizer(resp.Body)

  for {
    tt := htmlTokenizer.Next()

    switch {
    case tt == html.ErrorToken:
      //End when end of document
      return
    case tt == html.StartTagToken:
      t := htmlTokenizer.Token()

      // checks if t is an <a> tags in html
      isAnchor := t.Data == "a"
      if !isAnchor {
        continue
      }

      // Get href attribute from t if present
      ok, url := GetHref(t)
      if !ok {
        continue
      }

      // Makes sure the url begins with "http"
      hasProto := strings.Index(url, "http") == 0
      if hasProto {
        ch <- url
      }
    }
  }
}
