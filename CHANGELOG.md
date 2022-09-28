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
