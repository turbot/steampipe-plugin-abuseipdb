package main

import (
	"github.com/turbot/steampipe-plugin-abuseipdb/abuseipdb"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: abuseipdb.Plugin})
}
