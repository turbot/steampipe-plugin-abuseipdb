package abuseipdb

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type abuseipdbConfig struct {
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
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
