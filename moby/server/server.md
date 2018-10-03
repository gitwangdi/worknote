# struct Server
```
{  
    config  
    servers         []*HTTPServer  
    routers         []router.Router  
    routerSwaper    *routerSwapper  
    middlewares     []middleware.Middleware  
}
```
  
New(config) *Server  

## (s *Server)  
>**s.UserMiddleware(m)** add m to middlewares  
**s.Accept(addr, listeners...)** add servers  
**s.Close()** loop and close servers  
**s.serveAPI()** listen to all the servers  
**s.makeHTTPHandler(handler)** ??  
**s.InitRouter** add routers  
**s.createMux()** initalizes the main router ??  
**s.wait()** Wait for servers end  
**s.handlerWithGlobalMiddlewares(hanler)** ??

## HTTPServer (s *HTTPServer)  
```
{  
    srv    *http.Server  
    l      net.Listener  
}
```
>**s.Serve()** start listening  
**s.Close()** close the server from listening  
