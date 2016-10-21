package main

// https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
// Web scraper from above tutorial

import (
  "net/http"
  //"io/ioutil"
  "golang.org/x/net/html"
  "fmt"
  //"os"
  "strings"
)

func main() {

  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
  // Make http request
  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

  // Change url to scrape different websites
  url := "https://www.golang.org"

  Crawl(url)
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

  return
}

func Crawl(url string) {

  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
  // Parse HTML for Anchor Tags
  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

  resp, _ := http.Get(url)

  htmlTokenizer := html.NewTokenizer(resp.Body)

  for {
    tt := htmlTokenizer.Next()

    switch {
    case tt == html.ErrorToken:
      //End when end of document
      return
    case tt == html.StartTagToken:
      t := htmlTokenizer.Token()

      // Finds a tags in html
      isAnchor := t.Data == "a"
      if !isAnchor {
        continue
      }

      // Get href attribute from t if present
      ok, val := GetHref(t)
      if !ok {
        continue
      }

      // Makes sure the url begins with "http"
      hasProto := strings.Index(val, "http") == 0
      if hasProto {
        fmt.Println("Found a link:", val)
      }
    }
  }

  resp.Body.Close()
}
