package handlers
//ALL TESTS WILL BE REVISED.
import(
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/tahasevim/responsiveweb/templates"
	"bytes"
	"net/url"
	"strings"
	"flag"
	"io/ioutil"
	"log"
)
var server string
func init(){
	flag.StringVar(&server,"server","localhost:8080","server flag")
}
func TestIpHandler(t *testing.T){
	testReq, err := http.NewRequest("GET","http://httpbin.org/ip",nil)
	if err != nil {
		t.Fatal(err)
	}

	resp,_ := http.DefaultClient.Do(testReq)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(getHandler)
	handler.ServeHTTP(resprec,testReq)
	log.Println(string(expectedResult))
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	log.Println(string(resprec.Body.Bytes()))
	log.Println(string(result))
	/*
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body Check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}*/
}
func TestHeadersHandler(t *testing.T){
	testReq, err := http.NewRequest("GET","http://httpbin.org/headers",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(testReq)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(headersHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")

	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}
func TestGetHandler(t *testing.T){
	testReq, err := http.NewRequest("GET","http://httpbin.org/get",nil)
	if err != nil {
		t.Fatal(err)
	}

	testReq.URL.Query().Add("testKey","testValue")	
	testReq.URL.Query().Add("testKey","testValue")
	resp,err := http.DefaultClient.Do(testReq)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(getHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}

func TestIndexHandler(t *testing.T){	
	req, err := http.NewRequest("GET",server+"/",nil)
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
	testReq, err := http.NewRequest("GET","http://httpbin.org/user-agent",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(testReq)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(useragentHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin","user-agent")
	expectedResult = deleteJSONval(expectedResult,"url","origin","user-agent")

	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}

func TestPostHandler(t *testing.T){
	body := strings.NewReader(`testK=testV`)
	testReq, err := http.NewRequest("POST","http://httpbin.org/post",body)
	if err != nil {
		t.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(testReq)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	log.Println(string(result))
	log.Println(string(expectedResult))	
	/*//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}*/

}

func TestDeleteHandler(t *testing.T){
	form := url.Values{}
	form.Add("test1K","test1V")
	form.Add("test2K","test2v")
	req, err := http.NewRequest("DELETE",server+"/delete",strings.NewReader(form.Encode()))
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
	log.Println(string(result))
	log.Println(string(expectedResult))	
	/*if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}*/
}