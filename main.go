package main

import (
	"bytes"
	Cloudflare "github.com/hugomd/cloudflare-ddns/lib"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func checkIP() (string, error) {
	rsp, err := http.Get("http://checkip.amazonaws.com")
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
	APIKEY := os.Getenv("APIKEY")
	if APIKEY == "" {
		log.Fatal("APIKEY env. variable is required")
	}

	ZONE := os.Getenv("ZONE")
	if APIKEY == "" {
		log.Fatal("ZONE env. variable is required")
	}

	HOST := os.Getenv("HOST")
	if HOST == "" {
		log.Fatal("HOST env. variable is required")
	}

	EMAIL := os.Getenv("EMAIL")
	if EMAIL == "" {
		log.Fatal("EMAIL env. variable is required")
	}

	api, err := Cloudflare.New(APIKEY, EMAIL, ZONE, HOST)
	if err != nil {
		panic(err)
	}

	zones, err := api.ListZones()
	if err != nil {
		panic(err)
	}

	var zone Cloudflare.Zone

	for i := range zones {
		if zones[i].Name == ZONE {
			zone = zones[i]
		}
	}

	if zone == (Cloudflare.Zone{}) {
		panic("Zone not found")
	}

	records, err := api.ListDNSRecords(zone)
	if err != nil {
		panic(err)
	}

	var record Cloudflare.Record
	for i := range records {
		if records[i].Name == HOST {
			record = records[i]
		}
	}

	ip, err := checkIP()
	if err != nil {
		panic(err)
	}

	if record == (Cloudflare.Record{}) {
		panic("Host not found")
	}

	record.Content = ip

	err = api.UpdateDNSRecord(record, zone)
	if err != nil {
		panic(err)
	}
}
