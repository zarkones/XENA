package proxy

import (
	"crypto/rsa"
	"crypto/tls"
	"encoding/hex"
	"encoding/pem"
	"net"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/elazarl/goproxy"
	cry "github.com/zarkones/xena-crypto"
)

type Req struct {
	ID     int64
	Host   string
	Method string
	Path   string
	Body   bool
	Length int
	RawHex string
}

var RequestsStream = []Req{}

var RequestsCh = make(chan Req, 999)

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

	go func() {
		for req := range RequestsCh {
			// TODO: Add more details about the request.
			RequestsStream = append(RequestsStream, req)
		}
	}()

	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	// proxy.Verbose = true

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			go func() {
				rawReq, _ := httputil.DumpRequest(r, true)
				req := Req{
					ID:     ctx.Session,
					Method: r.Method,
					Host:   r.Host,
					Path:   r.URL.Path,
					Length: len(rawReq),
					RawHex: hex.EncodeToString(rawReq),
				}
				RequestsCh <- req
			}()
			return r, nil
		})

	return http.ListenAndServe(net.JoinHostPort(addr, port), proxy)
}
