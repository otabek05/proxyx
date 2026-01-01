package proxy

import (
	"log"

	wsPproxy "github.com/yeqown/fasthttp-reverse-proxy/v2"

	"github.com/valyala/fasthttp"
)


func (p *ProxyServer) configureWSProxy() {
	for _, c := range p.config {
		for _, r := range c.Spec.Routes {
			if r.Websocket != nil {
				customProxy, err :=  newWsProxyInstance(r.Websocket.URL)
				if err != nil {
					log.Fatal(err)
				}

				p.wsProxies[c.Spec.Domain] = customProxy
			}
		}
	}
}


func (p *ProxyServer) websocketProxyHandler(ctx *fasthttp.RequestCtx) {
	wsServer, ok := p.wsProxies[string(ctx.Host())]
	if !ok {
		ServeProxyHomepage(ctx)
		return 
	} 

	ctx.Request.Header.Set(wsPproxy.DefaultOverrideHeader, string(ctx.Path()))
	wsServer.ServeHTTP(ctx)
}


func newWsProxyInstance(serverURL string ) (*wsPproxy.WSReverseProxy, error ) {
	return wsPproxy.NewWSReverseProxyWith(
		wsPproxy.WithURL_OptionWS(serverURL),
		wsPproxy.WithDynamicPath_OptionWS(
			true,
			wsPproxy.DefaultOverrideHeader,
		),
	)
}