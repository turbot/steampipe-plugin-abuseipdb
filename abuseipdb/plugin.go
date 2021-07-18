package abuseipdb

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-abuseipdb",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"abuseipdb_category":   tableAbuseIPDbCategory(ctx),
			"abuseipdb_check_cidr": tableAbuseIPDbCheckCidr(ctx),
			"abuseipdb_check_ip":   tableAbuseIPDbCheckIP(ctx),
			"abuseipdb_deny":       tableAbuseIPDbDeny(ctx),
		},
	}
	return p
}
