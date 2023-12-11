---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/abuseipdb.svg"
brand_color: "#4e7e14"
display_name: "AbuseIPDB"
short_name: "abuseipdb"
description: "Steampipe plugin to query IP address abuse data and more from AbuseIPDB."
og_description: "Query AbuseIPDB with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/abuseipdb-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# AbuseIPDB + Steampipe

[AbuseIPDB](https://abuseipdb.com) allows users to report and identify IP addresses that have been associated with malicious activity online.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Query IP address base data from AbuseIPDB:

```sql
select
  ip_address,
  abuse_confidence_score,
  last_reported_at
from
  abuseipdb_check_ip
where
  ip_address = '76.76.21.21'
```

```
+-------------+------------------------+---------------------+
| ip_address  | abuse_confidence_score | last_reported_at    |
+-------------+------------------------+---------------------+
| 76.76.21.21 | 73                     | 2021-07-10 15:01:31 |
+-------------+------------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/abuseipdb/tables)**

## Get started

### Install

Download and install the latest AbuseIPDB plugin:

```bash
steampipe plugin install abuseipdb
```

### Configuration

Installing the latest abuseipdb plugin will create a config file (`~/.steampipe/config/abuseipdb.spc`) with a single connection named `abuseipdb`:

```hcl
connection "abuseipdb" {
  plugin = "abuseipdb"
  api_key = "5a76843869c183a4ea901c79102bfa1f667f39a2ea0ba857c9a35a9885d91fbd9c4ae24d6a10999f"
}
```

- `api_key` - Free API key for authenticated access.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-abuseipdb
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
