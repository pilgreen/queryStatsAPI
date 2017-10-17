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

var Config struct {
  Key string
  Secret string
  Endpoint string
}

func main() {
  flag.StringVar(&Config.Key, "key", os.Getenv("STATS_KEY"), "Stats API key")
  flag.StringVar(&Config.Secret, "secret", os.Getenv("STATS_SECRET"), "Stats API secret")
  flag.StringVar(&Config.Endpoint, "endpoint", "", "Stats API endpoint")
  flag.Parse()

  args := flag.Args()
  if len(Config.Endpoint) == 0 && len(args) > 0 {
    Config.Endpoint = args[0]
  }

  if len(Config.Key) == 0 || len(Config.Secret) == 0 || len(Config.Endpoint) == 0 {
    flag.Usage()
    os.Exit(1)
  }

  // Start the process
  now := time.Now()
  sig := fmt.Sprintf("%s%s%d", Config.Key, Config.Secret, now.Unix())
  sigsum := sha256.Sum256([]byte(sig))

  urlString := fmt.Sprintf("http://api.stats.com/v1/stats/%s", strings.Trim(Config.Endpoint, "/"))
  u, _ := url.Parse(urlString)

  q := u.Query()
  q.Add("accept", "json")
  q.Add("api_key", Config.Key)
  q.Add("sig", fmt.Sprintf("%x", sigsum))
  u.RawQuery = q.Encode()

  resp, err := http.Get(u.String())
  if err != nil { log.Fatal(err) }

  defer resp.Body.Close()
  io.Copy(os.Stdout, resp.Body)
}
