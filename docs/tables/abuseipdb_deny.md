# Table: abuseipdb_deny

List IPs with an abuse confidence score above a minimum. This list is often
used as a deny list of IPs.

Notes:

- If not specified, the default `confidence_minimum` is 90.
- The free tier limit is 5 requests per day.

## Examples

### Basic deny list (default confidence minimum of 90)

```sql
select
  *
from
  abuseipdb_deny
```

### List IPs with a confidence minimum of 95

```sql
select
  *
from
  abuseipdb_deny
where
  confidence_minimum = 95
```
