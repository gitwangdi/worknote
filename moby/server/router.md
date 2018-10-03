# struct routerSwapper
```
{
    mu       sysnc.Mutex
    router   *mux.Router
}
```

**rs.Swap(\*mux.Router)** set router  
**rs.ServerHTTP()** implement the handler interface ??  

# struct Route 
```go
type Router interface {
	Routes() []Route
}
type Route interface {
	Handler() httputils.APIFunc
	Method() string
	Path() string
}
```

两种route：  
lcoalRoute   
experimentalRoute

# 未完待续