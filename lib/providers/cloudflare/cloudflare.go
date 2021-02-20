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

var ZONEID, HOST string

func NewProvider() (providers.Provider, error) {
	// Check for use of any deprecated variables first, point to how to update
	if os.Getenv("CLOUDFLARE_APIKEY") != "" {
		log.Fatal("Do not use CLOUDFLARE_APIKEY, see https://github.com/hugomd/cloudflare-ddns#deprecated-environment-variables")
	}
	if os.Getenv("CLOUDFLARE_EMAIL") != "" {
		log.Fatal("Do not use CLOUDFLARE_EMAIL, see https://github.com/hugomd/cloudflare-ddns#deprecated-environment-variables")
	}
	if os.Getenv("CLOUDFLARE_ZONE") != "" {
		log.Fatal("Do not use CLOUDFLARE_ZONE, see https://github.com/hugomd/cloudflare-ddns#deprecated-environment-variables")
	}

	APITOKEN := os.Getenv("CLOUDFLARE_APITOKEN")
	if APITOKEN == "" {
		log.Fatal("CLOUDFLARE_APITOKEN env. variable is required")
	}

	ZONEID = os.Getenv("CLOUDFLARE_ZONEID")
	if ZONEID == "" {
		log.Fatal("CLOUDFLARE_ZONEID env. variable is required")
	}

	HOST = os.Getenv("CLOUDFLARE_HOST")
	if HOST == "" {
		log.Fatal("CLOUDFLARE_HOST env. variable is required")
	}

	api, err := NewCloudflareClient(APITOKEN, ZONEID, HOST)

	if err != nil {
		return nil, err
	}

	provider := &Cloudflare{
		client: api,
	}

	return provider, nil
}

func (api *Cloudflare) UpdateRecord(ipv4 string, ipv6 string) error {
	records, err := api.client.ListDNSARecords()
	if err != nil {
		return err
	}

    if (ipv6 != "") {
        a4records, err := api.client.ListDNSAAAARecords()
        if err != nil {
            log.Println("cant get AAAA records from Cloudflare ", err)
        }
        records = append(records, a4records...)
    }

	var arecord Record
    var a4record Record

	for i := range records {
		if records[i].Name == HOST && records[i].Type == "A" {
			arecord = records[i]
		}
        if records[i].Name == HOST && records[i].Type == "AAAA" {
            a4record = records[i]
        }
	}

	if arecord == (Record{}) && a4record == (Record{}) {
		return errors.New("Host not found")
	}

	if ipv4 != arecord.Content {
		arecord.Content = ipv4
		err = api.client.UpdateDNSRecord(arecord)
		if err != nil {
			return err
		}
		log.Printf("IPv4 changed, updated to %s", ipv4)
	} else {
		log.Print("No change in IPv4, not updating A record")
	}

	if ipv6 != "" && ipv6 != arecord.Content {
		a4record.Content = ipv6
		err = api.client.UpdateDNSRecord(a4record)
		if err != nil {
			return err
		}
		log.Printf("IPv6 changed, updated to %s", ipv6)
	} else {
		log.Print("No change in IPv6, not updating AAAA record")
	}

	return nil
}
