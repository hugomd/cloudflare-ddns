package main

import (
	"bytes"
        "flag"
	"github.com/hugomd/cloudflare-ddns/lib/providers"
	_ "github.com/hugomd/cloudflare-ddns/lib/providers/_all"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func checkIP() (string, error) {
	rsp, err := http.Get("https://checkip.amazonaws.com")
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(buf)), nil
}

func main() {
	var runonce bool
	var ticker *time.Ticker

	CheckDuration := flag.Duration("duration", 0, "update interval (ex. 15s, 1m, 6h); if not specified or set to 0s, run only once and exit")
	flag.Parse()

	if *CheckDuration == time.Duration(0) {
		runonce = true
	} else {
		ticker = time.NewTicker(*CheckDuration)
	}

	runddns()

	if runonce {
		os.Exit(0)
	}

	for range ticker.C {
		runddns()
	}

	return
}

func runddns() {
	PROVIDER := os.Getenv("PROVIDER")
	if PROVIDER == "" {
		log.Fatal("PROVIDER env. variable is required")
	}

	provider, err := providers.Providers[PROVIDER]()
	if err != nil {
		panic(err)
	}

	ip, err := checkIP()
	if err != nil {
		panic(err)
	}
	log.Printf("IP is %s", ip)

	err = provider.UpdateRecord(ip)
	if err != nil {
		panic(err)
	}
}
