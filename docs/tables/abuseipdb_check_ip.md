# Table: abuseipdb_check_ip

Query the abuse confidence score and other information about an IP address.

## Examples

### Get information about an IP

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

### List all reports for a given IP over the last year

```sql
select
  report ->> 'reportedAt' as reported_at,
  report ->> 'comment' as comment,
  report ->> 'categories' as categories
from
  abuseipdb_check_ip,
  jsonb_array_elements(reports) as report
where
  ip_address = '76.76.21.21'
  and max_age_in_days = 365
```

### Top categories for reports against this IP in the last year

```sql
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
  count desc
```
