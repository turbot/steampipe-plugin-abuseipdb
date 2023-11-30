---
title: "Steampipe Table: abuseipdb_deny - Query AbuseIPDB Denylists using SQL"
description: "Allows users to query AbuseIPDB Denylists, specifically the IP addresses that have been reported for abusive activities, providing insights into potential security threats."
---

# Table: abuseipdb_deny - Query AbuseIPDB Denylists using SQL

AbuseIPDB is a project dedicated to helping combat the spread of hackers, spammers, and abusive activity on the internet. Its primary function is to provide a platform for internet service providers, network administrators, and other interested parties to share and access data about abusive IP addresses. It aids in the identification of sources of malicious activities and helps in implementing preventive measures.

## Table Usage Guide

The `abuseipdb_deny` table provides insights into IP addresses that have been reported for abusive activities on the AbuseIPDB platform. As a network administrator or security analyst, explore details about these IP addresses through this table, including their abuse confidence score, country of origin, and associated reports. Utilize it to uncover information about potential security threats, such as those from known malicious sources, for effective threat intelligence and preventive measures.

## Examples

### Basic deny list (default confidence minimum of 90)
Explore which IP addresses are considered malicious based on a default confidence score of 90 or above. This helps in enhancing your network's security by blocking potentially harmful traffic.

```sql
select
  *
from
  abuseipdb_deny
```

### List IPs with a confidence minimum of 95
Discover the segments that have a high confidence level of 95, allowing you to focus on the most reliable data for your security analysis. This is particularly useful when you need to prioritize actions based on the degree of certainty in the data.

```sql
select
  *
from
  abuseipdb_deny
where
  confidence_minimum = 95
```