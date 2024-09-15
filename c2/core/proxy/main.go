package proxy

import (
	"c2/models"
	proxyRepo "c2/repos/proxy"
	"crypto/rsa"
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/elazarl/goproxy"
	cry "github.com/zarkones/xena-crypto"
)

func Start(addr, port, pathToDerCert string, privKey *rsa.PrivateKey) (err error) {
	rawDer, err := os.ReadFile(pathToDerCert)
	if err != nil {
		return err
	}

	privKeyPEM, err := cry.PrivKeyToPEM(privKey)
	if err != nil {
		return err
	}

	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: rawDer,
	}

	pemBytes := pem.EncodeToMemory(pemBlock)

	goproxyCA, err := tls.X509KeyPair(pemBytes, []byte(privKeyPEM))
	if err != nil {
		return err
	}

	goproxy.GoproxyCa = goproxyCA
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCA)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCA)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCA)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCA)}

	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	// proxy.Verbose = true

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			go func() {
				rawReq, _ := httputil.DumpRequest(r, true)
				req := models.ProxyReq{
					SessionID: ctx.Session,
					Method:    r.Method,
					Host:      r.Host,
					Path:      r.URL.Path,
					Query:     r.URL.RawQuery,
					ReqLength: len(rawReq),
					RawReq:    string(rawReq),
					Time:      time.Now(),
				}
				if err := proxyRepo.Insert(&req); err != nil {
					fmt.Println("proxy: failed to insert request:", req.SessionID)
					return
				}
			}()

			return r, nil
		})

	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		rawResp, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println("proxy: httputil.DumpResponse:", err)
		}
		go func() {
			sResp := string(rawResp)
			if err := proxyRepo.UpdateRawResp(ctx.Session, resp.StatusCode, &sResp); err != nil {
				fmt.Println("proxy: failed to insert request:", ctx.Session)
				return
			}
		}()
		return resp
	})

	return http.ListenAndServe(net.JoinHostPort(addr, port), proxy)
}
