package helpers

import (
	"fmt"
	"strings"
	"time"
)

// Dga stands for Domain Generation Algorithm and it returns a list of domains
// which are potentially in possesion of the bot herder.
func Dga(dgaSeed int) []string {
	domains := []string{}
	for _, topLevelDOmain := range topLevelDomains {
		for index := 0; index <= 50; index++ {
			domain := strings.ToLower(fmt.Sprint(time.Now().Month())+fmt.Sprint(time.Now().Day()*(dgaSeed+time.Now().Year())*index)+fmt.Sprint(time.Now().Month())) + topLevelDOmain
			domains = append(domains, domain)
		}
	}
	return domains
}

var topLevelDomains = []string{
	".com",
	".co",
	".studio",
	".tech",
	".solutions",
	".services",
	".org",
	".net",
	".int",
	".xyz",
	".network",
	".africa",
	".art",
	".io",
	".church",
	".faith",
	".co.uk",
	".best",
	".biz",
	".blog",
	".club",
	".earth",
	".global",
	".love",
	".market",
	".pizza",
	".site",
	".social",
	".store",
	".online",
	".click",
	".digital",
	".world",
	".email",
	".ru",
	".rs",
	".al",
	".ch",
	".ad",
	".af",
	".am",
	".bo",
	".br",
	".bs",
	".cd",
	".cn",
	".et",
	".kr",
	".kz",
	".ly",
	".me",
	".ng",
	".pk",
	".tr",
	".za",
	".zm",
	".zw",
}
