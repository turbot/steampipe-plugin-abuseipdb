package abuseipdb

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAbuseIPDbDeny(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "abuseipdb_deny",
		Description: "List all IP addresses over a given abuse confidence score (default 90).",
		List: &plugin.ListConfig{
			Hydrate:    listDeny,
			KeyColumns: plugin.OptionalColumns([]string{"confidence_minimum"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip_address", Type: proto.ColumnType_IPADDR, Description: "IP address to deny."},
			{Name: "abuse_confidence_score", Type: proto.ColumnType_INT, Description: "Abuse confidence score."},
			{Name: "last_reported_at", Type: proto.ColumnType_TIMESTAMP, Description: "Last time when the IP was reported."},
			// Other columns
			{Name: "confidence_minimum", Type: proto.ColumnType_INT, Transform: transform.FromQual("confidence_minimum"), Description: "Minimum confidence score."},
		},
	}
}

func listDeny(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("abuseipdb_deny.listDeny", "connection_error", err)
		return nil, err
	}

	equalQuals := d.KeyColumnQuals

	min := 90
	if equalQuals["confidence_minimum"] != nil {
		min = int(equalQuals["confidence_minimum"].GetInt64Value())
	}

	var limit *int64
	if d.QueryContext.Limit != nil {
		limit = d.QueryContext.Limit
	}

	data, err := conn.DenyList(ctx, min, limit)
	if err != nil {
		plugin.Logger(ctx).Error("abuseipdb_deny.listDeny", "query_error", err)
		return nil, err
	}

	for _, i := range data {
		i.ConfidenceMinimum = min
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
