// Copyright Cosmin Cremarenco 2012
package main

import "fmt"
import "os"
import "net/http"
import "io/ioutil"

func handleError(err error) {
  if err != nil {
    panic(fmt.Sprintf("Failed to get quote, reason: %v", err))
  }
}

func main() {
  args := os.Args
  if len(args) != 2 {
    fmt.Printf("Usage: getquote <ticker>\n")
    return
  }
  ticker := args[1]

  url := fmt.Sprintf("http://quote.yahoo.com/d/quotes.csv?s=%s&f=l1", ticker)

  resp, err := http.Get(url)
  handleError(err)

  defer resp.Body.Close()

  if resp.StatusCode == 200 {
    bytes, err2 := ioutil.ReadAll(resp.Body)
    handleError(err2)

    asString := string(bytes)
    fmt.Println(asString)
  }
}

