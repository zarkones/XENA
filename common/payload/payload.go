package payload

const TYPE_SECURITY_PIPELINE = 0
const TYPE_PORT_SCAN = 1
const TYPE_RUN_XENA_LANG = 2
const TYPE_REVERSE_DNS_LOOKUP = 3
const TYPE_SUBDOMAIN_ENUM = 4
const TYPE_TLD_ENUM = 5
const TYPE_GRAB_BANNER = 6
const TYPE_WEB_CRAWL = 7

type Generic struct {
	Type int `json:"type"`
}

type PortScanCtx struct {
	Type  int    `json:"type"`
	Ports []int  `json:"ports"`
	Host  string `json:"host"`
}

type RunXenaLangCtx struct {
	Type   int    `json:"type"`
	Script string `json:"script"`
}

type ReverseDNSLookupCtx struct {
	Type      int    `json:"type"`
	IpAddress string `json:"ipAddress"`
}

type SubdomainEnumCtx struct {
	Type     int      `json:"type"`
	Domain   string   `json:"domain"`
	Wordlist []string `json:"wordlist"`
}

type TldEnumCtx struct {
	Type     int      `json:"type"`
	Domain   string   `json:"domain"`
	Wordlist []string `json:"wordlist"`
}

type GrabBannerCtx struct {
	Type int    `json:"type"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type WebCrawlCtx struct {
	Type   int    `json:"type"`
	Domain string `json:"domain"`
}
