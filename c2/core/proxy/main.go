package proxy

import (
	"crypto/rsa"
	"crypto/tls"
	"encoding/pem"
	"net"
	"net/http"
	"os"

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

	// proxy.OnRequest().DoFunc(
	// 	func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// 		h, _, _ := time.Now().Clock()
	// 		if h >= 8 && h <= 17 {
	// 			return r, goproxy.NewResponse(r,
	// 				goproxy.ContentTypeText, http.StatusForbidden,
	// 				"Don't waste your time!")
	// 		} else {
	// 			ctx.Warnf("clock: %d, you can waste your time...", h)
	// 		}
	// 		return r, nil
	// 	})

	return http.ListenAndServe(net.JoinHostPort(addr, port), proxy)
}
