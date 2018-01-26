
# responsiveweb [![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://travis-ci.org/tahasevim/responsiveweb)[![GoDoc](https://godoc.org/github.com/tahasevim/responsiveweb?status.svg)](https://godoc.org/github.com/tahasevim/responsiveweb)[![Coverage](https://img.shields.io/badge/coverage-85%25-orange.svg)](https://github.com/tahasevim/responsiveweb/tree/master/handlers)
![Alt text](https://image.ibb.co/k3oueG/gopherr.png)<br>
Basic HTTP server written in Go.<br>
Inspired from https://httpbin.org/.<br>
Main aims of the responsiveweb project are:<br>
  - Learning crucial features and infrastructures of Go.<br>
  - Learning basics of the HTTP applications and relations between each.<br>
Clients can make any request to this HTTP server in order to test their requests.
## To-Do List
Given below endpoints's handlers should be implemented<br>
- [x] `/`
- [x] `/ip`
- [x] `/uuid`
- [x] `/user-agent`
- [x] `/header` 
- [x] `/get` 
- [x] `/post` 
- [x] `/delete`
- [x] `/put` 
- [ ] `/patch` 
- [x] `/anything` 
- [x] `/encoding/utf8` 
- [x] `/gzip` 
- [x] `/deflate` 
- [x] `/brotli`
- [x] `/status/:code` 
- [x] `/response-headers?key=val` 
- [x] `/redirect/:n` 
- [x] `/redirect-to?url=foo`
- [x] `/redirect-to?url=foo&status_code=307` 
- [ ] `/relative-redirect/:n`
- [ ] `/absolute-redirect/:n` 
- [x] `/cookies` 
- [x] `/cookies/set?name=value`
- [x] `/cookies/delete?name`
- [x] `/basic-auth/:user/:passwd`
- [x] `/hidden-basic-auth/:user/:passwd`
- [ ] `/digest-auth/:qop/:user/:passwd/:`
- [x] `/stream/:n`
- [x] `/delay/:n`
- [ ] `/drip?numbytes=n&duration=s&delay=s&code=code`
- [ ] `/range/1024?duration=s&chunk_size=code`
- [x] `/html`
- [x] `/robots.txt`
- [x] `/deny`
- [x] `/cache` 
- [ ] `/etag/:etag` 
- [x] `/cache/:n`
- [x] `/bytes/:n`
- [ ] `/stream-bytes/:n` 
- [x] `/links/:n` 
- [x] `/image`
- [x] `/image/png`
- [x] `/image/jpeg`
- [x] `/image/webp`
- [x] `/image/svg`
- [x] `/forms/post`
- [x] `/xml`

## Install
`go get github.com/tahasevim/responsiveweb`
## Usage and Examples
#### Usage
First,web server should be simply run with below command:
`responsiveweb`<br>
Or you can run with your custom `port` with using `port`flag:<br>
`responsiveweb port=PORTNUMBER`
#### Examples
To test web server,you should use HTTP requests.Simply you can use cURL to test easily.<br>

```bash
$ curl --data "testK=testV" localhost:8080/post
{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "testK": "testV"
  },
  "headers": {
    "Accept": "*/*",
    "Content-Length": "11",
    "Content-Type": "application/x-www-form-urlencoded",
    "User-Agent": "curl/7.54.0"
  },
  "json": "",
  "origin": "[::1]:55628",
  "url": "localhost:8080/post"
}
```
```bash
$ echo "This sentence is a test for posting file.">>testFile.txt
$ curl -F "testFile.txt=@./testFile.txt" localhost:8080/post
{
  "args": {},
  "data": "",
  "files": {
    "testFile.txt": "This sentence is a test for posting file.\n"
  },
  "form": {},
  "headers": {
    "Accept": "*/*",
    "Content-Length": "240",
    "Content-Type": "multipart/form-data; boundary=------------------------c0cc45e9a422852d",
    "Expect": "100-continue",
    "User-Agent": "curl/7.54.0"
  },
  "json": "",
  "origin": "[::1]:55801",
  "url": "localhost:8080/post"
}
```
```bash
$ curl --data "you cant post this data to /get" localhost:8080/get
Method Not Allowed
```
```
$ curl localhost:8080/get
{
  "args": {},
  "headers": {
    "Accept": "*/*",
    "User-Agent": "curl/7.54.0"
  },
  "origin": "[::1]:55814",
  "url": "localhost:8080/get"
}

```
```
$ curl -X PUT -d testK=testV localhost:8080/put
{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "testK": "testV"
  },
  "headers": {
    "Accept": "*/*",
    "Connection": "",
    "Content-Length": "11",
    "Content-Type": "application/x-www-form-urlencoded",
    "User-Agent": "curl/7.54.0"
  },
  "json": "",
  "origin": "[::1]:51045",
  "url": "localhost:8080/put"
}
```
