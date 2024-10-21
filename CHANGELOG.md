## v1.0.0 [2024-10-22]

There are intentionally no significant changes in this plugin version, but it has been released to coincide with the [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin follows [semantic versioning's specification](https://semver.org/#semantic-versioning-specification-semver) and preserves backward compatibility in each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#42](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/42))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#42](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/42))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#29](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/29))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#29](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/29))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-abuseipdb/blob/main/docs/LICENSE). ([#29](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/29))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#28](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/28))

## v0.5.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#21](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/21))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#19](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/19))

- Recompiled plugin with Go version `1.21`. ([#19](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/19))

## v0.4.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#15](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/15))

## v0.3.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#12](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/12))
- Recompiled plugin with Go version `1.19`. ([#12](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/12))

## v0.2.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#9](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/9))

## v0.2.0 [2022-04-27]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#6](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/6))
- Added support for native Linux ARM and Mac M1 builds. ([#7](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/7))

## v0.1.0 [2021-12-15]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#3](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/3))
- Recompiled plugin with Go version 1.17 ([#3](https://github.com/turbot/steampipe-plugin-abuseipdb/pull/3))

## v0.0.2 [2021-10-18]

_Enhancements_

- Update plugin SDK to v1.7.0

_Bug fixes_

- Handle null for last_reported_at column [#1](https://github.com/turbot/steampipe-plugin-abuseipdb/issues/1)


## v0.0.1 [2021-07-20]

_What's new?_

- New tables added
  - [abuseipdb_category](https://hub.steampipe.io/plugins/turbot/abuseipdb/tables/abuseipdb_category)
  - [abuseipdb_check_cidr](https://hub.steampipe.io/plugins/turbot/abuseipdb/tables/abuseipdb_check_cidr)
  - [abuseipdb_check_ip](https://hub.steampipe.io/plugins/turbot/abuseipdb/tables/abuseipdb_check_ip)
  - [abuseipdb_deny](https://hub.steampipe.io/plugins/turbot/abuseipdb/tables/abuseipdb_deny)
