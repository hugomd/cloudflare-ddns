package cloudflare

import (
	"errors"
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
		log.Fatal("CLOUDFLARE_APIKEY env. variable is required")
	}

	ZONE = os.Getenv("CLOUDFLARE_ZONE")
	if APIKEY == "" {
		log.Fatal("CLOUDFLARE_ZONE env. variable is required")
	}

	HOST = os.Getenv("CLOUDFLARE_HOST")
	if HOST == "" {
		log.Fatal("CLOUDFLARE_HOST env. variable is required")
	}

	EMAIL := os.Getenv("CLOUDFLARE_EMAIL")
	if EMAIL == "" {
		log.Fatal("CLOUDFLARE_EMAIL env. variable is required")
	}

	// Check for use of any deprecated variables first, point to how to update
	if os.Getenv("CLOUDFLARE_APIKEY") != "" || os.Getenv("CLOUDFLARE_EMAIL") != "" || os.Getenv("CLOUDFLARE_ZONE") != "" {
		log.Print("WARNING: CLOUDFLARE_APIKEY, CLOUDFLARE_EMAIL and CLOUDFLARE_ZONE are deprecated environment variables and are unsupported. Please see https://github.com/hugomd/cloudflare-ddns#deprecated-environment-variables for more information")
	}

	api, err := NewCloudflareClient(APIKEY, EMAIL, ZONE, HOST)
	if err != nil {
		return nil, err
	}

	provider := &Cloudflare{
		client: api,
	}

	return provider, nil
}

func (api *Cloudflare) UpdateRecord(ip string) error {
	zones, err := api.client.ListZones()
	if err != nil {
		return err
	}

	var zone Zone

	for i := range zones {
		if zones[i].Name == ZONE {
			zone = zones[i]
		}
	}

	if zone == (Zone{}) {
		return errors.New("Zone not found")
	}

	records, err := api.client.ListDNSRecords(zone)
	if err != nil {
		return err
	}

	var record Record
	for i := range records {
		if records[i].Name == HOST {
			record = records[i]
		}
	}

	if record == (Record{}) {
		return errors.New("Host not found")
	}

	if ip != record.Content {
		record.Content = ip
		err = api.client.UpdateDNSRecord(record, zone)
		if err != nil {
			return err
		}
		log.Printf("IP changed, updated to %s", ip)
	} else {
		log.Print("No change in IP, not updating record")
	}

	return nil
}
