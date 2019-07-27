package cloudflare

import (
	"github.com/hugomd/cloudflare-ddns/lib/providers"
	"log"
	"os"
)

type Cloudflare struct {
	client *CloudflareAPI
}

func init() {
	providers.RegisterProvider("cloudflare", NewProvider)
}

var ZONE, HOST string

func NewProvider() (providers.Provider, error) {
	APIKEY := os.Getenv("CLOUDFLARE_APIKEY")
	if APIKEY == "" {
		log.Fatal("APIKEY env. variable is required")
	}

	ZONE = os.Getenv("CLOUDFLARE_ZONE")
	if APIKEY == "" {
		log.Fatal("ZONE env. variable is required")
	}

	HOST = os.Getenv("CLOUDFLARE_HOST")
	if HOST == "" {
		log.Fatal("HOST env. variable is required")
	}

	EMAIL := os.Getenv("CLOUDFLARE_EMAIL")
	if EMAIL == "" {
		log.Fatal("EMAIL env. variable is required")
	}

	api, err := NewCloudflareClient(APIKEY, EMAIL, ZONE, HOST)
	if err != nil {
		panic(err)
	}

	provider := &Cloudflare{
		client: api,
	}

	return provider, nil
}

func (api *Cloudflare) UpdateRecord(ip string) error {
	zones, err := api.client.ListZones()
	if err != nil {
		panic(err)
	}

	var zone Zone

	for i := range zones {
		if zones[i].Name == ZONE {
			zone = zones[i]
		}
	}

	if zone == (Zone{}) {
		panic("Zone not found")
	}

	records, err := api.client.ListDNSRecords(zone)
	if err != nil {
		panic(err)
	}

	var record Record
	for i := range records {
		if records[i].Name == HOST {
			record = records[i]
		}
	}

	if record == (Record{}) {
		panic("Host not found")
	}

	if ip != record.Content {
		record.Content = ip
		err = api.client.UpdateDNSRecord(record, zone)
		if err != nil {
			panic(err)
		}
		log.Printf("IP changed, updated to %s", ip)
	} else {
		log.Print("No change in IP, not updating record")
	}

	return nil
}
