package main

import (
  "crypto/sha256"
  "flag"
  "fmt"
  "io"
  "log"
  "net/url"
  "net/http"
  "os"
  "strings"
  "time"
)

func main() {
  // Primary Variables
  key := os.Getenv("STATS_KEY")
  secret := os.Getenv("STATS_SECRET")
  now := time.Now()

  if(len(key) == 0 || len(secret) == 0) {
    log.Fatal("STATS_KEY and STATS_SECRET environment variables are not set")
  }

  // Endpoint Flag
  var endpoint string
  flag.StringVar(&endpoint, "endpoint", "REQUIRED", "the url endpoint for the API call")
  flag.Parse()

  if endpoint == "REQUIRED" {
    flag.Usage()
    os.Exit(1)
  }

  // Start the process
  sig := fmt.Sprintf("%s%s%d", key, secret, now.Unix())
  sigsum := sha256.Sum256([]byte(sig))

  urlString := fmt.Sprintf("http://api.stats.com/v1/stats/%s", strings.Trim(endpoint, "/"))
  u, _ := url.Parse(urlString)

  q := u.Query()
  q.Add("accept", "json")
  q.Add("api_key", key)
  q.Add("sig", fmt.Sprintf("%x", sigsum))
  u.RawQuery = q.Encode()

  resp, err := http.Get(u.String())
  if err != nil { log.Fatal(err) }

  defer resp.Body.Close()
  io.Copy(os.Stdout, resp.Body)
}
