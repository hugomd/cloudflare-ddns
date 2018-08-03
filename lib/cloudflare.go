package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type API struct {
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

func New(key string, email string, zone string, host string) (*API, error) {
	api := API{
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

func (api *API) ListZones() ([]Zone, error) {
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

func (api *API) ListDNSRecords(zone Zone) ([]Record, error) {
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

func (api *API) UpdateDNSRecord(record Record, zone Zone) error {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zone.ID, record.ID)

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(record)

	_, err := api.request("PUT", uri, payload)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) request(method string, uri string, body io.Reader) ([]byte, error) {
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
