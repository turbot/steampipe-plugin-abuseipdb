---
title: "Steampipe Table: abuseipdb_check_cidr - Query AbuseIPDB CIDR Reports using SQL"
description: "Allows users to query CIDR Reports in AbuseIPDB, specifically the reports of IP addresses that have been reported as abusive, providing insights into potential security threats."
---

# Table: abuseipdb_check_cidr - Query AbuseIPDB CIDR Reports using SQL

AbuseIPDB is a service that allows users to report and check IP addresses that have been involved in malicious activities like hacking attempts, spamming, and brute force attacks. It offers a comprehensive database of internet protocol addresses that have been reported as abusive. It helps in identifying potential security threats and aids in taking appropriate preventative measures.

## Table Usage Guide

The `abuseipdb_check_cidr` table provides insights into CIDR reports within AbuseIPDB. As a security analyst, explore CIDR-specific details through this table, including the number of reports, abuse confidence score, and associated metadata. Utilize it to uncover information about potential security threats, such as the most reported IP addresses, the nature of the reported abuse, and the geographical location of the reported IPs.

## Examples

### List information about IPs in a CIDR range
Identify instances where specific IP addresses within a certain range have been reported for abuse. This can help in assessing potential security threats and taking necessary preventive actions.

```sql+postgres
select
  ip_address,
  abuse_confidence_score,
  num_reports,
  last_reported_at
from
  abuseipdb_check_cidr
where
  cidr = '76.76.21.20/30';
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```