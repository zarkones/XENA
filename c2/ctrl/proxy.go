package ctrl

import (
	"c2/core"
	"c2/models"
	proxyRepo "c2/repos/proxy"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DEFAULT_PROXIED_REQ_PER_PAGE = 28
)

func GetProxiedRequest(c *gin.Context) {
	reqID, err := strconv.ParseInt(c.Param("reqID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	request, err := proxyRepo.Get(reqID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"err": err})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusOK, request)
}

func GetProxiedRequests(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	offset := DEFAULT_PROXIED_REQ_PER_PAGE * page
	orderBy := q.Get("orderBy")
	orderDirection := q.Get("order")
	search := q.Get("search")

	var requests []models.ProxyReq
	var err error

	if len(search) == 0 {
		requests, err = proxyRepo.GetMultiple(offset, DEFAULT_PROXIED_REQ_PER_PAGE, orderBy, orderDirection)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			return
		}
	} else {
		rMap := map[int64]models.ProxyReq{}
		wg := &sync.WaitGroup{}
		mux := &sync.Mutex{}
		offset := 0
		done := false
		const ROUTINES = 30

		for {
			if done {
				break
			}

			for i := 0; i < ROUTINES; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					r, err := proxyRepo.GetMultiple(offset*i, DEFAULT_PROXIED_REQ_PER_PAGE, orderBy, orderDirection)
					if err != nil {
						return
					}

					if len(r) == 0 {
						done = true
					}

					for _, req := range r {
						if strings.Contains(req.Host, search) {
							mux.Lock()
							rMap[req.ID] = req
							mux.Unlock()
							continue
						}
						if strings.Contains(req.Path, search) {
							mux.Lock()
							rMap[req.ID] = req
							mux.Unlock()
							continue
						}
						if strings.Contains(req.RawReq, search) {
							mux.Lock()
							rMap[req.ID] = req
							mux.Unlock()
							continue
						}
						if strings.Contains(req.RawResp, search) {
							mux.Lock()
							rMap[req.ID] = req
							mux.Unlock()
							continue
						}
						if strings.Contains(req.Query, search) {
							mux.Lock()
							rMap[req.ID] = req
							mux.Unlock()
							continue
						}
					}
				}()
			}

			offset += ROUTINES

			wg.Wait()
		}

		requests = make([]models.ProxyReq, len(rMap))
		i := 0
		for _, req := range rMap {
			requests[i] = req
			i++
		}
	}

	if len(requests) == 0 {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, requests)
}

func GetCertificate(c *gin.Context) {
	cert, err := os.ReadFile(core.CertPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.Writer.Header().Add("Content-Disposition", "attachment; filename=\"xena-cert.der\"")

	c.Writer.Write(cert)
}
