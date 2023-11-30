---
title: "Steampipe Table: abuseipdb_category - Query AbuseIPDB Categories using SQL"
description: "Allows users to query AbuseIPDB Categories, specifically to retrieve details about the categories of IP addresses reported for abusive behavior."
---

# Table: abuseipdb_category - Query AbuseIPDB Categories using SQL

AbuseIPDB is a service that allows users to report and check IP addresses for known malicious activity. It categorizes IP addresses based on the types of abuse reported, such as fraud orders, DDoS attacks, spam emails, etc. This categorization helps in identifying the nature of the threat posed by a particular IP address.

## Table Usage Guide

The `abuseipdb_category` table provides insights into the categories of IP addresses reported for abusive behavior in the AbuseIPDB. As a security analyst, explore category-specific details through this table, including the types of abuse associated with each category. Utilize it to enhance your understanding of the threat landscape and to aid in decision-making for threat mitigation strategies.

## Examples

### List the categories
Explore all the categories available in the AbuseIPDB to understand the types of abusive behavior that are tracked, allowing for more efficient and targeted security measures.

```sql
select
  *
from
  abuseipdb_category
order by
  id
```