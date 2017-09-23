
# responsiveweb [![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://travis-ci.org/tahasevim/responsiveweb)
Basic HTTP server written in Go.<br>
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
- [ ] `/gzip` 
- [ ] `/deflate` 
- [ ] `/brotli`
- [ ] `/status/:code` 
- [ ] `/response-headers?key=val` 
- [ ] `/redirect/:n` 
- [ ] `/redirect-to?url=foo`
- [ ] `/redirect-to?url=foo&status_code=307` 
- [ ] `/relative-redirect/:n`
- [ ] `/absolute-redirect/:n` 
- [ ] `/cookies` 
- [ ] `/cookies/set?name=value`
- [ ] `/cookies/delete?name`
- [ ] `/basic-auth/:user/:passwd`
- [ ] `/hidden-basic-auth/:user/:passwd`
- [ ] `/digest-auth/:qop/:user/:passwd/:`
- [ ] `/stream/:n`
- [ ] `/delay/:n`
- [ ] `/drip?numbytes=n&duration=s&delay=s&code=code`
- [ ] `/range/1024?duration=s&chunk_size=code`
- [ ] `/html`
- [ ] `/robots.txt`
- [ ] `/deny`
- [ ] `/cache` 
- [ ] `/etag/:etag` 
- [ ] `/cache/:n`
- [ ] `/bytes/:n`
- [ ] `/stream-bytes/:n` 
- [ ] `/links/:n` 
- [ ] `/image`
- [ ] `/image/png`
- [ ] `/image/jpeg`
- [ ] `/image/webp`
- [ ] `/image/svg`
- [ ] `/forms/post`
- [ ] `/xml`

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
