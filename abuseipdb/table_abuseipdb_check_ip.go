package abuseipdb

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableAbuseIPDbCheckIP(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "abuseipdb_check_ip",
		Description: "List all checks for a given IP address.",
		List: &plugin.ListConfig{
			Hydrate: listCheckIP,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "ip_address"},
				{Name: "max_age_in_days", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip_address", Type: proto.ColumnType_IPADDR, Description: "IP address to check."},
			{Name: "abuse_confidence_score", Type: proto.ColumnType_INT, Description: "Abuse confidence score."},
			// Other columns
			{Name: "country_code", Type: proto.ColumnType_STRING, Description: "Country code where the IP server is located."},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain name found at the IP."},
			{Name: "ip_version", Type: proto.ColumnType_INT, Description: "IP address version."},
			{Name: "is_public", Type: proto.ColumnType_BOOL, Description: "True if the IP address is public."},
			{Name: "is_whitelisted", Type: proto.ColumnType_BOOL, Description: "True if the IP address has been whitelisted."},
			{Name: "isp", Type: proto.ColumnType_STRING, Description: "ISP hosting the IP."},
			{Name: "last_reported_at", Type: proto.ColumnType_TIMESTAMP, Description: "Last time when the IP was reported."},
			{Name: "max_age_in_days", Type: proto.ColumnType_INT, Description: "Max age in days of the report data."},
			{Name: "num_distinct_users", Type: proto.ColumnType_INT, Description: "Number of users reporting the IP."},
			{Name: "reports", Type: proto.ColumnType_JSON, Description: "Report details."},
			{Name: "total_reports", Type: proto.ColumnType_INT, Description: "Total number of reports for this IP."},
			{Name: "usage_type", Type: proto.ColumnType_STRING, Description: "Usage type, e.g. Commercial."},
		},
	}
}

func listCheckIP(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("abuseipdb_check.listCheckIP", "connection_error", err)
		return nil, err
	}

	equalQuals := d.KeyColumnQuals
	ipAddress := equalQuals["ip_address"].GetInetValue().GetAddr()

	// Default to 90 days max age, but use key column if provided
	maxAge := 90
	if equalQuals["max_age_in_days"] != nil {
		i := equalQuals["max_age_in_days"].GetInt64Value()
		maxAge = int(i)
	}

	data, err := conn.CheckIP(ctx, ipAddress, maxAge)
	if err != nil {
		plugin.Logger(ctx).Error("abuseipdb_check.listCheckIP", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, data)
	return nil, nil
}
