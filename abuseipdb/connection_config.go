package abuseipdb

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type abuseipdbConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &abuseipdbConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) abuseipdbConfig {
	if connection == nil || connection.GetConfig() == nil {
		return abuseipdbConfig{}
	}
	config, _ := connection.GetConfig().(abuseipdbConfig)
	return config
}
