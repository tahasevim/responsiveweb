
# responsiveweb
Basic HTTP server written in Go.<br>
Clients can make any request to this HTTP server in order to test their requests.
## To-Do List
- [] `/put` endpoint handler
- [] `/patch` endpoint handler
- [] `/anything` endpoint handler
- [] `/encoding/utf8` endpoint handler
- [] `/gzip` endpoint handler
- [] `/deflate` endpoint handler
- [] `/brotli` endpoint handler
- [] `/status/:code` endpoint handler
- [] `/response-headers?key=val` endpoint handler
- [] `/redirect/:n` endpoint handler
- [] `/redirect-to?url=foo` endpoint handler
- [] `/redirect-to?url=foo&status_code=307` endpoint handler
- [] `/relative-redirect/:n` endpoint handler
- [] `/absolute-redirect/:n` endpoint handler
- [] `/cookies` endpoint handler
- [] `/cookies/set?name=value` endpoint handler
- [] `/cookies/delete?name` endpoint handler
- [] `/basic-auth/:user/:passwd` endpoint handler
- [] `/hidden-basic-auth/:user/:passwd` endpoint handler 
- [] `/digest-auth/:qop/:user/:passwd/:`endpoint handler

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

```
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
```
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
```
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
