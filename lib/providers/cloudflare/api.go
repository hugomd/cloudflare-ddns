package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CloudflareAPI struct {
	ZoneID     string
	Host       string
	APIToken   string
	BaseURL    string
	httpClient *http.Client
}

type Record struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Name    string `json:"name"`
	Proxied bool   `json:"proxied"`
}

type RecordResponse struct {
	Result []Record `json:"result"`
}

func NewCloudflareClient(token string, zoneID string, host string) (*CloudflareAPI, error) {
	api := CloudflareAPI{
		ZoneID:   zoneID,
		Host:     host,
		APIToken: token,
		BaseURL:  "https://api.cloudflare.com/client/v4",
	}

	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	}

	return &api, nil
}

func (api *CloudflareAPI) ListDNSARecords() ([]Record, error) {
	uri := fmt.Sprintf("/zones/%s/dns_records?type=A&name=%s", api.ZoneID, api.Host)
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

func (api *CloudflareAPI) ListDNSAAAARecords() ([]Record, error) {
	uri := fmt.Sprintf("/zones/%s/dns_records?type=AAAA&name=%s", api.ZoneID, api.Host)
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

func (api *CloudflareAPI) UpdateDNSRecord(record Record) error {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", api.ZoneID, record.ID)

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

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.APIToken))

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Status code not 200")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
