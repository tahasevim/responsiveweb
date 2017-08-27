package handlers

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/tahasevim/responsiveweb/templates"
	"bytes"
)

func TestIpHandler(t *testing.T){
	req, err := http.NewRequest("GET","/ip",nil)
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
	req, err := http.NewRequest("GET","/headers",nil)
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
	req, err := http.NewRequest("GET","/get",nil)
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
	req, err := http.NewRequest("GET","/",nil)
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