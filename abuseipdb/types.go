package abuseipdb

import "net/http"

type errorObj struct {
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

type abuseipdbCheckResponse struct {
	Data   abuseipdbCheckData `json:"data,omitempty"`
	Errors []errorObj         `json:"errors,omitempty"`
}

type abuseipdbCheckData struct {
	IPAddress            string                 `json:"ipAddress"`
	MaxAgeInDays         int                    `json:"maxAgeInDays,omitempty"`
	IsPublic             bool                   `json:"isPublic"`
	IPVersion            int                    `json:"ipVersion"`
	IsWhitelisted        bool                   `json:"isWhitelisted"`
	AbuseConfidenceScore int                    `json:"abuseConfidenceScore"`
	CountryCode          string                 `json:"countryCode"`
	UsageType            string                 `json:"usageType"`
	Isp                  string                 `json:"isp"`
	Domain               string                 `json:"domain"`
	TotalReports         int                    `json:"totalReports"`
	NumDistinctUsers     int                    `json:"numDistinctUsers"`
	LastReportedAt       string                 `json:"lastReportedAt"`
	Reports              []abuseipdbCheckReport `json:"reports"`
}

type abuseipdbCheckReport struct {
	ReportedAt          string `json:"reportedAt"`
	Comment             string `json:"comment"`
	Categories          []int  `json:"categories"`
	ReporterID          int    `json:"reporterId"`
	ReporterCountryCode string `json:"reporterCountryCode"`
	ReporterCountryName string `json:"reporterCountryName"`
}

type abuseipdbCheckCidrResponse struct {
	Data   abuseipdbCheckCidrData `json:"data,omitempty"`
	Errors []errorObj             `json:"errors,omitempty"`
}

type abuseipdbCheckCidrData struct {
	NetworkAddress   string                          `json:"networkAddress"`
	Netmask          string                          `json:"netmask"`
	MinAddress       string                          `json:"minAddress"`
	MaxAddress       string                          `json:"maxAddress"`
	NumPossibleHosts int                             `json:"numPossibleHosts"`
	AddressSpaceDesc string                          `json:"addressSpaceDesc"`
	ReportedAddress  []abuseipdbCheckReportedAddress `json:"reportedAddress"`
}

type abuseipdbCheckReportedAddress struct {
	IPAddress            string `json:"ipAddress"`
	NumReports           int    `json:"numReports"`
	LastReportedAt       string `json:"mostRecentReport"`
	AbuseConfidenceScore int    `json:"abuseConfidenceScore"`
	CountryCode          string `json:"countryCode"`
	MaxAgeInDays         int    `json:"maxAgeInDays,omitempty"`
}

type abuseipdbDenyListResponse struct {
	Data   []abuseipdbDenyListItem `json:"data"`
	Errors []errorObj              `json:"errors,omitempty"`
}

type abuseipdbDenyListItem struct {
	IPAddress            string `json:"ipAddress"`
	AbuseConfidenceScore int    `json:"abuseConfidenceScore"`
	LastReportedAt       string `json:"lastReportedAt"`
	ConfidenceMinimum    int    `json:"confidenceMinimum,omitempty"`
}

type abuseipdbClient struct {
	HTTPClient *http.Client
	APIKey     string
}
