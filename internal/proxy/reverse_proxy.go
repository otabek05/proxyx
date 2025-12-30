package proxy

import (
	"crypto/tls"
	"log"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var reqPool = sync.Pool{
	New: func() interface{} { return new(fasthttp.Request) },
}

var resPool = sync.Pool{
	New: func() interface{} { return new(fasthttp.Response) },
}

func (p *ProxyServer) reverseProxyxHandler(ctx *fasthttp.RequestCtx, matched *routeInfo) {
	target := matched.loadBalancer.Next()
	if target == nil {
		ServeProxyHomepage(ctx)
		return
	}

	proxyKey := target.String()
	proxy, ok := p.proxies[proxyKey]
	if !ok {

		proxy = &fasthttp.HostClient{
			Addr:                          target.Host,
			IsTLS:                         target.Scheme == "https",
			MaxConns:                      2000,
			MaxIdleConnDuration:           30 * time.Second,
			ReadTimeout:                   10 * time.Second,
			WriteTimeout:                  10 * time.Second,
			DisableHeaderNamesNormalizing: true,
			TLSConfig: &tls.Config{
				ServerName:         target.Hostname(),
				InsecureSkipVerify: false,
				MinVersion:         tls.VersionTLS12,
			},
		}

		p.proxies[proxyKey] = proxy
	}

	req := reqPool.Get().(*fasthttp.Request)
	defer func() { req.Reset(); reqPool.Put(req) }()
	ctx.Request.CopyTo(req)

	resp := resPool.Get().(*fasthttp.Response)
	defer func() { resp.Reset(); resPool.Put(resp) }()

	clientIP := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if clientIP == "" {
		clientIP = ctx.RemoteAddr().String()
	}

	if xff := req.Header.Peek("X-Forwarded-For"); len(xff) > 0 {
		req.Header.Set("X-Forwarded-For", string(xff)+", "+clientIP)
	} else {
		req.Header.Set("X-Forwarded-For", clientIP)
	}

	//req.Header.Set("X-Forwarded-For", clientIP)
	req.Header.Set("X-Forwarded-Host", string(ctx.Host()))
	req.Header.Set("X-Forwarded-Proto", map[bool]string{true: "https", false: "http"}[ctx.IsTLS()])

	if err := proxy.DoTimeout(req, resp, 5 *time.Second); err != nil {
		log.Println(err)
		ctx.Error("Failed to reach backend", fasthttp.StatusBadGateway)
		return
	}

	resp.CopyTo(&ctx.Response)
}
