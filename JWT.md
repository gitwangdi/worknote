# JWT  

Contains three parts.

xxxxx.yyyyy.zzzzz

## Header  

Contains two fields: the type of the token, the hash algorithm.

```go
{
    "alg": "HS256",
    "typ": "JWT"
}
```

Use Base64Url to encode it.

## payload

Contains three chaims: registered chaims, public chaims and private chaims.  

```go
{
    "sub": "123456",
    "name": "John",
    "admin": true,
}
```

Use Base64Url to encode it.  

## Signature  

Use hash algrothm to hash the information above.  
For example:  

```go
HMACHASH256(base64UrlEncode(header),".",base64UrlEncode(payload),secret)
```

Put three parts(after encoding with Base64Url) seperated by dots.  
