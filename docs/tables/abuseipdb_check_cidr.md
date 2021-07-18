# Table: abuseipdb_check_cidr

Query the abuse confidence score for all IPs in a CIDR range.

## Examples

### List information about IPs in a CIDR range

```sql
select
  ip_address,
  abuse_confidence_score,
  num_reports,
  last_reported_at
from
  abuseipdb_check_cidr
where
  cidr = '76.76.21.20/30'
```
