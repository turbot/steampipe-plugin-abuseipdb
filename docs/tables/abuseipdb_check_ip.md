---
title: "Steampipe Table: abuseipdb_check_ip - Query AbuseIPDB IP Reports using SQL"
description: "Allows users to query IP Reports in AbuseIPDB, specifically the IP address details, providing insights into IP reputation and potential abuse activities."
---

# Table: abuseipdb_check_ip - Query AbuseIPDB IP Reports using SQL

AbuseIPDB is a service that helps system administrators, IT security analysts, and developers to understand whether a specific IP address has been reported for suspicious activity. It provides an API to check and report IP addresses associated with malicious activities like spamming, brute-force attacks, and other forms of abuse. AbuseIPDB helps in identifying potential threats and taking necessary actions to prevent cyber attacks.

## Table Usage Guide

The `abuseipdb_check_ip` table provides insights into IP addresses reported for suspicious activities in AbuseIPDB. As a system administrator or IT security analyst, explore IP-specific details through this table, including abuse reports, reputation score, and associated metadata. Utilize it to uncover information about reported IPs, such as their reputation, the number of times they've been reported, and the categories of abuse they've been associated with.

## Examples

### Get information about an IP
Analyze the settings to understand the potential abuse history and confidence score for a specific IP address. This query is useful for identifying potential security risks and recent abuse reports associated with that IP.

```sql+postgres
select
  ip_address,
  abuse_confidence_score,
  last_reported_at
from
  abuseipdb_check_ip
where
  ip_address = '76.76.21.21';
```

```sql+sqlite
select
  ip_address,
  abuse_confidence_score,
  last_reported_at
from
  abuseipdb_check_ip
where
  ip_address = '76.76.21.21';
```

### List all reports for a given IP over the last year
Determine the reports associated with a specific IP address over the past year. This can be useful for understanding the history of suspicious activities and identifying recurring patterns or categories of abuse.

```sql+postgres
select
  report ->> 'reportedAt' as reported_at,
  report ->> 'comment' as comment,
  report ->> 'categories' as categories
from
  abuseipdb_check_ip,
  jsonb_array_elements(reports) as report
where
  ip_address = '76.76.21.21'
  and max_age_in_days = 365;
```

```sql+sqlite
select
  json_extract(report.value, '$.reportedAt') as reported_at,
  json_extract(report.value, '$.comment') as comment,
  json_extract(report.value, '$.categories') as categories
from
  abuseipdb_check_ip,
  json_each(reports) as report
where
  ip_address = '76.76.21.21'
  and max_age_in_days = 365;
```

### Top categories for reports against this IP in the last year
Determine the most frequently reported categories associated with a specific IP address over the past year. This can be useful for identifying trends or patterns in abusive behavior for cybersecurity purposes.

```sql+postgres
select
  category_id.value as category,
  c.title,
  count(*)
from
  abuseipdb_check_ip as ch,
  jsonb_array_elements(ch.reports) as report,
  jsonb_array_elements(report->'categories') as category_id,
  abuseipdb_category as c
where
  ip_address = '76.76.21.21'
  and max_age_in_days = 365
  and c.id = category_id.value::int
group by
  category_id.value,
  c.title
order by
  count desc;
```

```sql+sqlite
select
  json_extract(category_id.value, '$') as category,
  c.title,
  count(*)
from
  abuseipdb_check_ip as ch,
  json_each(json_extract(ch.reports, '$')) as report,
  json_each(json_extract(report.value, '$.categories')) as category_id,
  abuseipdb_category as c
where
  ch.ip_address = '76.76.21.21'
  and ch.max_age_in_days = 365
  and c.id = cast(json_extract(category_id.value, '$') as int)
group by
  json_extract(category_id.value, '$'),
  c.title
order by
  count(*) desc;
```

