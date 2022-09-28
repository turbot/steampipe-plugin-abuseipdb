package abuseipdb

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableAbuseIPDbCheckCidr(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "abuseipdb_check_cidr",
		Description: "List check results for all IPs within a CIDR range.",
		List: &plugin.ListConfig{
			Hydrate: listCheckCidr,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "cidr"},
				{Name: "max_age_in_days", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip_address", Type: proto.ColumnType_IPADDR, Description: "IP address."},
			{Name: "abuse_confidence_score", Type: proto.ColumnType_INT, Description: "Abuse confidence score for the IP."},
			// Other columns
			{Name: "cidr", Type: proto.ColumnType_CIDR, Transform: transform.FromQual("cidr"), Description: "CIDR range to check."},
			{Name: "country_code", Type: proto.ColumnType_STRING, Description: "Country code where the IP is located."},
			{Name: "last_reported_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time of the last report for this IP."},
			{Name: "max_age_in_days", Type: proto.ColumnType_INT, Description: "Max age in days of the report data."},
			{Name: "num_reports", Type: proto.ColumnType_INT, Description: "Number of reports for this IP address."},
		},
	}
}

func listCheckCidr(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("abuseipdb_check_cidr.listCheckCidr", "connection_error", err)
		return nil, err
	}

	equalQuals := d.KeyColumnQuals
	cidr := equalQuals["cidr"].GetInetValue().GetCidr()

	// Default to 30 days max age, but use key column if provided
	maxAge := 30
	if equalQuals["max_age_in_days"] != nil {
		i := equalQuals["max_age_in_days"].GetInt64Value()
		maxAge = int(i)
	}

	data, err := conn.CheckCidr(ctx, cidr, maxAge)
	if err != nil {
		plugin.Logger(ctx).Error("abuseipdb_check_cidr.listCheckCidr", "query_error", err)
		return nil, err
	}

	for _, i := range data {
		i.MaxAgeInDays = maxAge
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
