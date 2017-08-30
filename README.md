
# responsiveweb [![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://travis-ci.org/tahasevim/responsiveweb)
Basic HTTP server written in Go.<br>
Clients can make any request to this HTTP server in order to test their requests.
## To-Do List
- [x] `/` endpoint handler
- [x] `/ip` endpoint handler
- [x] `/uuid` endpoint handler
- [x] `/user-agent` endpoint handler
- [x] `/header` endpoint handler
- [x] `/get` endpoint handler
- [x] `/post` endpoint handler
- [x] `/delete` endpoint handler
- [ ] `/put` endpoint handler
- [ ] `/patch` endpoint handler
- [ ] `/anything` endpoint handler
- [ ] `/encoding/utf8` endpoint handler
- [ ] `/gzip` endpoint handler
- [ ] `/deflate` endpoint handler
- [ ] `/brotli` endpoint handler
- [ ] `/status/:code` endpoint handler
- [ ] `/response-headers?key=val` endpoint handler
- [ ] `/redirect/:n` endpoint handler
- [ ] `/redirect-to?url=foo` endpoint handler
- [ ] `/redirect-to?url=foo&status_code=307` endpoint handler
- [ ] `/relative-redirect/:n` endpoint handler
- [ ] `/absolute-redirect/:n` endpoint handler
- [ ] `/cookies` endpoint handler
- [ ] `/cookies/set?name=value` endpoint handler
- [ ] `/cookies/delete?name` endpoint handler
- [ ] `/basic-auth/:user/:passwd` endpoint handler
- [ ] `/hidden-basic-auth/:user/:passwd` endpoint handler 
- [ ] `/digest-auth/:qop/:user/:passwd/:`endpoint handler
- [ ] `/digest-auth/:qop/:user/:passwd` endpoint handler
- [ ] `/stream/:n` endpoint handler
- [ ] `/delay/:n` endpoint handler 
- [ ] `/drip?numbytes=n&duration=s&delay=s&code=code` endpoint handler 
- [ ] `/range/1024?duration=s&chunk_size=code` endpoint handler 
- [ ] `/html` endpoint handler 
- [ ] `/robots.txt` endpoint handler 
- [ ] `/deny` endpoint handler 
- [ ] `/cache` endpoint handler 
- [ ] `/etag/:etag` endpoint handler 
- [ ] `/cache/:n` endpoint handler 
- [ ] `/bytes/:n` endpoint handler 
- [ ] `/stream-bytes/:n` endpoint handler 
- [ ] `/links/:n` endpoint handler 
- [ ] `/image` endpoint handler 
- [ ] `/image/png` endpoint handler 
- [ ] `/image/jpeg` endpoint handler 
- [ ] `/image/webp` endpoint handler 
- [ ] `/image/svg` endpoint handler 
- [ ] `/forms/post` endpoint handler 
- [ ] `/xml` endpoint handler 

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
    "testK": [
      "testV"
    ]
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
