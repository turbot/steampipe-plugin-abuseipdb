package abuseipdb

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type abuseipdbConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &abuseipdbConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) abuseipdbConfig {
	if connection == nil || connection.Config == nil {
		return abuseipdbConfig{}
	}
	config, _ := connection.Config.(abuseipdbConfig)
	return config
}
