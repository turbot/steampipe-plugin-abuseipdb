package abuseipdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func (c *abuseipdbClient) Get(ctx context.Context, path string, params url.Values) ([]byte, error) {

	bytes := []byte{}

	u := url.URL{
		Scheme:   "https",
		Host:     "api.abuseipdb.com",
		Path:     fmt.Sprintf("/api/v2%s", path),
		RawQuery: params.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return bytes, err
	}

	req.Header.Add("User-Agent", "Steampipe")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Key", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return bytes, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func combineErrors(errs []errorObj) error {
	errStrings := []string{}
	for _, i := range errs {
		errStrings = append(errStrings, fmt.Sprintf("HTTP %d: %s", i.Status, i.Detail))
	}
	return fmt.Errorf(strings.Join(errStrings, "\n"))
}

func (c *abuseipdbClient) CheckIP(ctx context.Context, ipAddress string, maxAgeInDays int) (abuseipdbCheckData, error) {
	v := url.Values{}
	v.Set("ipAddress", ipAddress)
	v.Set("maxAgeInDays", fmt.Sprintf("%d", maxAgeInDays))
	v.Set("verbose", "1")
	bytes, err := c.Get(ctx, "/check", v)
	if err != nil {
		return abuseipdbCheckData{}, err
	}
	result := abuseipdbCheckResponse{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return abuseipdbCheckData{}, err
	}
	if result.Errors != nil {
		return abuseipdbCheckData{}, combineErrors(result.Errors)
	}
	// Set the max age according to the request, for easier matching in the result row
	result.Data.MaxAgeInDays = &maxAgeInDays
	return result.Data, nil
}

func (c *abuseipdbClient) CheckCidr(ctx context.Context, cidr string, maxAgeInDays int) ([]abuseipdbCheckReportedAddress, error) {
	v := url.Values{}
	v.Set("network", cidr)
	v.Set("maxAgeInDays", fmt.Sprintf("%d", maxAgeInDays))
	bytes, err := c.Get(ctx, "/check-block", v)
	if err != nil {
		return []abuseipdbCheckReportedAddress{}, err
	}
	result := abuseipdbCheckCidrResponse{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return []abuseipdbCheckReportedAddress{}, err
	}
	if result.Errors != nil {
		return []abuseipdbCheckReportedAddress{}, combineErrors(result.Errors)
	}
	// Set the max age according to the request, for easier matching in the result row
	return result.Data.ReportedAddress, nil
}

func (c *abuseipdbClient) DenyList(ctx context.Context, confidenceMinimum int, limit *int64) ([]abuseipdbDenyListItem, error) {
	v := url.Values{}
	v.Set("confidenceMinimum", fmt.Sprintf("%d", confidenceMinimum))
	if limit != nil {
		v.Set("limit", fmt.Sprintf("%d", *limit))
	}
	bytes, err := c.Get(ctx, "/blacklist", v)
	if err != nil {
		return []abuseipdbDenyListItem{}, err
	}
	result := abuseipdbDenyListResponse{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return []abuseipdbDenyListItem{}, err
	}
	if result.Errors != nil {
		return []abuseipdbDenyListItem{}, combineErrors(result.Errors)
	}
	return result.Data, nil
}

func connect(_ context.Context, d *plugin.QueryData) (*abuseipdbClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "abuseipdb"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*abuseipdbClient), nil
	}

	// Default to the env var settings
	apiKey := os.Getenv("ABUSEIPDB_API_KEY")

	// Prefer config settings
	abuseipdbConfig := GetConfig(d.Connection)
	if &abuseipdbConfig != nil {
		if abuseipdbConfig.APIKey != nil {
			apiKey = *abuseipdbConfig.APIKey
		}
	}

	// Defaults
	timeout := 30

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, errors.New("api_key must be configured")
	}

	httpClient := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	conn := &abuseipdbClient{
		HTTPClient: httpClient,
		APIKey:     apiKey,
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}
