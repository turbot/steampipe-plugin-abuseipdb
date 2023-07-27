![image](https://hub.steampipe.io/images/plugins/turbot/abuseipdb-social-graphic.png)

# AbuseIPDB Plugin for Steampipe

Use SQL to query AbuseIPDB configuration and more from AbuseIPDB.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/abuseipdb)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/abuseipdb/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-abuseipdb/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install abuseipdb
```

Run a query:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-abuseipdb.git
cd steampipe-plugin-abuseipdb
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/abuseipdb.spc
```

Try it!

```
steampipe query
> .inspect abuseipdb
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-abuseipdb/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [AbuseIPDB Plugin](https://github.com/turbot/steampipe-plugin-abuseipdb/labels/help%20wanted)
