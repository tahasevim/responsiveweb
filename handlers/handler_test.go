package handlers

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/tahasevim/responsiveweb/templates"
	"bytes"
	"net/url"
	"strings"
	"flag"
)
var server = flag.String("server","localhost:8080","server name")
func TestIpHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",*server+"/ip",nil)
	if err != nil {
		t.Fatal(err)
	}

	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(ipHandler)
	handler.ServeHTTP(resprec,req)
	if stat := resprec.Code;stat != http.StatusOK{
		t.Errorf("Something has gone wrong! Error Code:%v",stat)
	}
	js := getAllJSONdata(req,"origin")
	expectedResult := makeJSONresponse(js)
	result := resprec.Body.Bytes()
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}
func TestHeadersHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",*server+"/headers",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(headersHandler)
	handler.ServeHTTP(resprec,req)
	if stat := resprec.Code;stat != http.StatusOK{
		t.Errorf("Something has gone wrong! Error Code:%v",stat)
	}
	js := getAllJSONdata(req, "headers")
	expectedResult := makeJSONresponse(js)
	result := resprec.Body.Bytes()
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}
func TestGetHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",*server+"/get",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(getHandler)
	req.URL.Query().Add("testKey","testValue")
	handler.ServeHTTP(resprec,req)
	if stat := resprec.Code;stat != http.StatusOK{
		t.Errorf("Something has gone wrong! Error Code:%v",stat)
	}
	js := getAllJSONdata(req,"args","headers","origin","url")
	expectedResult := makeJSONresponse(js)
	result := resprec.Body.Bytes()
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}

func TestIndexHandler(t *testing.T){	
	flag.Parse()
	req, err := http.NewRequest("GET",*server+"/",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(resprec,req)
	if stat := resprec.Code;stat != http.StatusOK{
		t.Errorf("Something has gone wrong! Error Code:%v",stat)
	}

	indexTemplate := templates.IndexTemplate
	var templ bytes.Buffer
	indexTemplate.Execute(&templ,nil)
	expectedResult := templ.String()
	result := resprec.Body.String()
	if string(expectedResult) != result {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)		
	}
}

func TestUseragentHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",*server+"/user-agent",nil)	
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(useragentHandler)
	handler.ServeHTTP(resprec,req)
	if stat := resprec.Code;stat != http.StatusOK{
		t.Errorf("Something has gone wrong! Error Code:%v",stat)		
	}
	js := getAllJSONdata(req,"user-agent")
	expectedResult := makeJSONresponse(js)
	result := resprec.Body.Bytes()
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}

func TestPostHandler(t *testing.T){
	form := url.Values{}
	form.Add("test1K","test1V")
	form.Add("test2K","test2v")
	flag.Parse()
	req, err := http.NewRequest("POST",*server+"/post",strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(resprec,req)
	if stat := resprec.Code;stat != http.StatusOK{
		t.Errorf("Something has gone wrong! Error Code:%v",stat)				
	}
	js := getAllJSONdata(req ,"args","data","files","form","headers","json","origin","url")
	expectedResult := makeJSONresponse(js)
	result := resprec.Body.Bytes()
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}

func TestDeleteHandler(t *testing.T){
	form := url.Values{}
	form.Add("test1K","test1V")
	form.Add("test2K","test2v")
	flag.Parse()
	req, err := http.NewRequest("DELETE",*server+"/delete",strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteHandler)
	handler.ServeHTTP(resprec,req)
	if stat := resprec.Code;stat != http.StatusOK{
		t.Errorf("Something has gone wrong! Error Code:%v",stat)				
	}
	js := getAllJSONdata(req ,"args","data","files","form","headers","json","origin","url")
	expectedResult := makeJSONresponse(js)
	result := resprec.Body.Bytes()
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}