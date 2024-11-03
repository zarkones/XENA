package analyzer

import (
	"c2/models"
	findingsRepo "c2/repos/findings"
	proxyRepo "c2/repos/proxy"
	"common/sec"
	"fmt"
	"log"
	"strings"
	"time"
)

func Start() {
	for range time.Tick(time.Second / 2) {
		traffic, err := proxyRepo.GetMultipleNonAnalyzed(10)
		if err != nil {
			log.Println("analyzer: proxyRepo:", err)
			continue
		}

		for _, req := range traffic {
			func() {
				secrets, err := sec.FindSecret(&req.RawResp)
				if err != nil {
					log.Println("analyzer: sec: FindSecret:", err)
					return
				}
				if len(secrets) == 0 {
					return
				}

				if err := findingsRepo.Insert(&models.Finding{
					RequestID: req.ID,
					Tag:       "SECRETS",
					Data:      strings.Join(secrets, "\n"),
				}); err != nil {
					fmt.Println("analyzer: findingsRepo: Insert:", err)
					return
				}
			}()

			func() {
				nodeModuleRefs := sec.FindNodeModulesReference(&req.RawResp)
				if len(nodeModuleRefs) == 0 {
					return
				}

				if err := findingsRepo.Insert(&models.Finding{
					RequestID: req.ID,
					Tag:       "NODE_MODULE_REFERENCES",
					Data:      strings.Join(nodeModuleRefs, "\n"),
				}); err != nil {
					fmt.Println("analyzer: findingsRepo: Insert:", err)
					return
				}
			}()

			if err := proxyRepo.SetAnalyzed(req.ID); err != nil {
				fmt.Println("analyzer: proxyRepo: SetAnalyzed:", err)
				continue
			}
		}
	}
}
