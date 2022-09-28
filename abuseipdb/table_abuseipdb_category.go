package abuseipdb

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableAbuseIPDbCategory(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "abuseipdb_category",
		Description: "Abuse categories.",
		List: &plugin.ListConfig{
			Hydrate: listCategory,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Category ID."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the category."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the category."},
		},
	}
}

type category struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func categories() []category {
	return []category{
		{ID: 1, Title: "DNS Compromise", Description: "Altering DNS records resulting in improper redirection."},
		{ID: 2, Title: "DNS Poisoning", Description: "Falsifying domain server cache (cache poisoning)."},
		{ID: 3, Title: "Fraud Orders", Description: "Fraudulent orders."},
		{ID: 4, Title: "DDoS Attack", Description: "Participating in distributed denial-of-service (usually part of botnet)."},
		{ID: 5, Title: "FTP Brute-Force", Description: ""},
		{ID: 6, Title: "Ping of Death", Description: "Oversized IP packet."},
		{ID: 7, Title: "Phishing", Description: "Phishing websites and/or email."},
		{ID: 8, Title: "Fraud VoIP", Description: ""},
		{ID: 9, Title: "Open Proxy", Description: "Open proxy, open relay, or Tor exit node."},
		{ID: 10, Title: "Web Spam", Description: "Comment/forum spam, HTTP referer spam, or other CMS spam."},
		{ID: 11, Title: "Email Spam", Description: "Spam email content, infected attachments, and phishing emails. Note: Limit comments to only relevent information (instead of log dumps) and be sure to remove PII if you want to remain anonymous."},
		{ID: 12, Title: "Blog Spam", Description: "CMS blog comment spam."},
		{ID: 13, Title: "VPN IP", Description: "Conjunctive category."},
		{ID: 14, Title: "Port Scan", Description: "Scanning for open ports and vulnerable services."},
		{ID: 15, Title: "Hacking", Description: ""},
		{ID: 16, Title: "SQL Injection", Description: "Attempts at SQL injection."},
		{ID: 17, Title: "Spoofing", Description: "Email sender spoofing."},
		{ID: 18, Title: "Brute-Force", Description: "Credential brute-force attacks on webpage logins and services like SSH, FTP, SIP, SMTP, RDP, etc. This category is seperate from DDoS attacks."},
		{ID: 19, Title: "Bad Web Bot", Description: "Webpage scraping (for email addresses, content, etc) and crawlers that do not honor robots.txt. Excessive requests and user agent spoofing can also be reported here."},
		{ID: 20, Title: "Exploited Host", Description: "Host is likely infected with malware and being used for other attacks or to host malicious content. The host owner may not be aware of the compromise. This category is often used in combination with other attack categories."},
		{ID: 21, Title: "Web App Attack", Description: "Attempts to probe for or exploit installed web applications such as a CMS like WordPress/Drupal, e-commerce solutions, forum software, phpMyAdmin and various other software plugins/solutions."},
		{ID: 22, Title: "SSH", Description: "Secure Shell (SSH) abuse. Use this category in combination with more specific categories."},
		{ID: 23, Title: "IoT Targeted", Description: "Abuse was targeted at an 'Internet of Things' type device. Include information about what type of device was targeted in the comments."},
	}
}

func listCategory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	for _, i := range categories() {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
