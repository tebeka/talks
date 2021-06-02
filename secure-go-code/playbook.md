- terminator
- vscode

slides up to "Code"

    - journal/db.go
    - journal/httpd.go

- gosec .
- gosec --help

- go run .

# Input:

## A1: Injection

- curl -d@add-1.json http://localhost:8080/add
- http://localhost:8080/last
- curl -d@add-2.json http://localhost:8080/add
- http://localhost:8080/last
./journal/db.go:49

restart the server

- database/sql

## A8: Insecure Deserialization

- ./journal/httpd.go:45
    - io.LimitReader

# Output:

## A7: Cross-Site Scripting (XSS)

- curl -d@add-3.json http://localhost:8080/add
- http://localhost:8080/last
- ./journal/httpd.go:38
- html/template

## A3: Sensitive Data Exposure

- ./journal/httpd.go:45 send error to client
- no return after error 

# Authentication

## A2: Broken Authentication

- ./journal/httpd.go:41 
    - middleware
    - listen on all interfaces

## A5: Broken Access Control

- ./journal/httpd.go:41 
    - ACL, RBAC

# Infrastruture

## A6: Security Misconfiguration

- ./journal/httpd.go
- http.ListenAndServeTLS
    - HTTP/2
    - letsencrypt

## A9: Using Components with Known Vulnerabilities

- go mod
- vendor
- replace
- own hosting
- depdenabot

## A10: Insufficient Logging & Monitoring

- log, zap
- expvar, prometheus
- alerting

## Beyond
