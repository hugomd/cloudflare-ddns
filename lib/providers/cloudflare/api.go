package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type CloudflareAPI struct {
	Zone       string
	Host       string
	APIKey     string
	Email      string
	BaseURL    string
	httpClient *http.Client
}

type Zone struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ZoneResponse struct {
	Result []Zone `json:"result"`
}

type Record struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

type RecordResponse struct {
	Result []Record `json:"result"`
}

func NewCloudflareClient(key string, email string, zone string, host string) (*CloudflareAPI, error) {
	api := CloudflareAPI{
		Zone:    zone,
		Host:    host,
		APIKey:  key,
		Email:   email,
		BaseURL: "https://api.cloudflare.com/client/v4",
	}

	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	}

	return &api, nil
}

func (api *CloudflareAPI) UpdateRecord(ip string) error {
	zones, err := api.ListZones()
	if err != nil {
		panic(err)
	}

	var zone Zone

	for i := range zones {
		if zones[i].Name == api.Zone {
			zone = zones[i]
		}
	}

	if zone == (Zone{}) {
		panic("Zone not found")
	}

	records, err := api.ListDNSRecords(zone)
	if err != nil {
		panic(err)
	}

	var record Record
	for i := range records {
		if records[i].Name == api.Host {
			record = records[i]
		}
	}

	if record == (Record{}) {
		panic("Host not found")
	}

	if ip != record.Content {
		record.Content = ip
		err = api.UpdateDNSRecord(record, zone)
		if err != nil {
			panic(err)
		}
		log.Printf("Updated IP to %s", ip)
	} else {
		log.Print("No change in IP, not updating record")
	}

	return nil
}

func (api *CloudflareAPI) ListZones() ([]Zone, error) {
	uri := fmt.Sprintf("/zones?name=%s", api.Zone)
	resp, err := api.request("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	var r ZoneResponse
	err = json.Unmarshal(resp, &r)

	if err != nil {
		return nil, err
	}
	return r.Result, nil
}

func (api *CloudflareAPI) ListDNSRecords(zone Zone) ([]Record, error) {
	uri := fmt.Sprintf("/zones/%s/dns_records?name=%s", zone.ID, api.Host)
	resp, err := api.request("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var r *RecordResponse
	err = json.Unmarshal(resp, &r)

	if err != nil {
		return nil, err
	}

	return r.Result, nil
}

func (api *CloudflareAPI) UpdateDNSRecord(record Record, zone Zone) error {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zone.ID, record.ID)

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(record)

	_, err := api.request("PUT", uri, payload)
	if err != nil {
		return err
	}

	return nil
}

func (api *CloudflareAPI) request(method string, uri string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, api.BaseURL+uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Email", api.Email)
	req.Header.Set("X-Auth-Key", api.APIKey)

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		panic("Status code not 200")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
