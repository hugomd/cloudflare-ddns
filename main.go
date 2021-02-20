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
	"strings"
	"time"
)

func checkIP() (string, string) {

    var ipv4 string
    var ipv6 string

	rsp, err := http.Get("https://ipv4.wtfismyip.com/text")
	if err != nil {
        log.Println("ipv4 get failed: ", err)
	}

    if rsp != nil {
        buf, err := ioutil.ReadAll(rsp.Body)
        if err != nil {
            log.Println("ipv4 read failed: ", err)
        }
        rsp.Body.Close()

        if buf != nil {
            ipv4 = string(bytes.TrimSpace(buf))
        }
    }

	rsp, err = http.Get("https://ipv6.wtfismyip.com/text")
	if err != nil {
        log.Println("ipv6 get failed: ", err)
	}

    if (rsp != nil) {
        buf, err := ioutil.ReadAll(rsp.Body)
        if err != nil {
            log.Println("ipv6 read failed: ", err)
        }
        rsp.Body.Close()

        if buf != nil {
            ipv6 = string(bytes.TrimSpace(buf))
        }
    }

    return ipv4, ipv6
}

func setEnvVarsFromConfig(filename *string) error {
	contents, err := ioutil.ReadFile(*filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if strings.Contains(line, "=") {
			values := strings.SplitN(line, "=", 2)
			os.Setenv(values[0], values[1])
		}
	}

	return nil
}

func main() {
	var runonce bool
	var ticker *time.Ticker

	CheckDuration := flag.Duration("duration", 0, "update interval (ex. 15s, 1m, 6h); if not specified or set to 0s, run only once and exit")
	ConfigFile := flag.String("config", "", "location of an (optional) config file to load environment variables from")
	flag.Parse()

	if *CheckDuration == time.Duration(0) {
		runonce = true
	} else {
		ticker = time.NewTicker(*CheckDuration)
	}

	if *ConfigFile != "" {
		err := setEnvVarsFromConfig(ConfigFile)
		if err != nil {
			panic(err)
		}
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

	ipv4, ipv6 := checkIP()
	if ipv4 == "" && ipv6 == "" {
		panic("cant get ipv4 or ipv6 addresses")
	}
	log.Printf("IPv4 is %s and IPv6 is %s", ipv4, ipv6)

	err = provider.UpdateRecord(ipv4, ipv6)
	if err != nil {
		panic(err)
	}
}
