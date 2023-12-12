package abuseipdb

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-abuseipdb",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
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
