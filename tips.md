```go
type HandlerFunc func(responseWriter, *Resquest)
func makeHandler() HandlerFunc{
    ...
}
// use
makeHandler()(resp, req)
```

```go
type inter interface {}
```

annotation:
```go
// return info
// <str1, str2, err>:
// str1: ...
// ...
```

```go
w.(flusher)
// w 是一个interface， 转换数据类型为flusher类型的 interface
```

```go
interface 蛮溜的一个应用
sort.Sort(interface{}) 要求interface代表的struct具有Len()/Less()/Swap()三个函数，只要对此类型定义这几个函数，就可以使用sort
```