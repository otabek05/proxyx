package common

type RouteType string

const (
    RouteStatic       RouteType = "Static"
    RouteReverseProxy RouteType = "ReverseProxy"
    RouteRedirect     RouteType = "Redirect"
    RouteWebsocket    RouteType = "Websocket"
)


func (r RouteType) IsValid() bool {
    switch r {
    case RouteStatic, RouteReverseProxy, RouteRedirect, RouteWebsocket:
        return true
    }
    return false
}


func (r RouteType) String() string {
    return string(r)
}