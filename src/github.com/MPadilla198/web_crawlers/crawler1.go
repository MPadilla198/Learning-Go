package main

import (
  "net/http"
  //"io/ioutil"
  "golang.org/x/net/html"
  "fmt"
  //"os"
  //"strings"
)

func main() {

  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
  // Make http request
  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

  url := "https://www.golang.org"

  resp, _ := http.Get(url)
  //bytes, _ := ioutil.ReadAll(resp.Body)

  //fmt.Println("HTML:\n\n", string(bytes))

  parseHTML(resp)

  resp.Body.Close()
}

func parseHTML(resp *http.Response) {

  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
  // Parse HTML for Anchor Tags
  //~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

  htmlTokenizer := html.NewTokenizer(resp.Body)

  for {
    tt := htmlTokenizer.Next()

    switch {
    case tt == html.ErrorToken:
      //End when end of document
      return
    case tt == html.StartTagToken:
      t := htmlTokenizer.Token()

      isAnchor := t.Data == "a"
      if isAnchor {
        fmt.Println("We found a link!")
      }
    }
  }
}
