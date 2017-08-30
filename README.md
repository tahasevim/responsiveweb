# responsiveweb
Basic HTTP server written in Go.<br>
Clients can make any request to this HTTP server in order to test their requests.
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
