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
	//"log"
)
var server string
func init(){
	flag.StringVar(&server,"server","localhost:8080","server flag")
}
func TestIpHandler(t *testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/ip",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(ipHandler)
	handler.ServeHTTP(resprec,req)
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}
}
func TestHeadersHandler(t *testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/headers",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/headers",nil)
	if err != nil {
		t.Fatal(err)
	}
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

func TestGetHandler(t *testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/get",nil)
	if err != nil {
		t.Fatal(err)
	}
	req.URL.Query().Add("testKey","testValue")	
	resp,err := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/get",nil)
	if err != nil {
		t.Fatal(err)
	}
	testReq.URL.Query().Add("testKey","testValue")	
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
	if expectedResult != result {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)		
	}
}

func TestUseragentHandler(t *testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/user-agent",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)
	
	testReq, err := http.NewRequest("GET","/user-agent",nil)
	if err != nil {
		t.Fatal(err)
	}
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
	req, err := http.NewRequest("POST","http://httpbin.org/post",body)
	if err != nil {
		t.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)
	
	body2 := strings.NewReader(`testK=testV`)
	testReq, err := http.NewRequest("POST","/post",body2)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin","json")
	expectedResult = deleteJSONval(expectedResult,"url","origin","json")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}

}

func TestDeleteHandler(t *testing.T){
	form := url.Values{}
	form.Add("test1K","test1V")
	form.Add("test2K","test2v")
	req, err := http.NewRequest("DELETE","http://httpbin.org/delete",strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)
	
	form2 := url.Values{}
	form2.Add("test1K","test1V")
	form2.Add("test2K","test2v")
	testReq, err := http.NewRequest("DELETE","/delete",strings.NewReader(form2.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteHandler)
	handler.ServeHTTP(resprec,testReq)
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin","json")
	expectedResult = deleteJSONval(expectedResult,"url","origin","json")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}

}
func TestPutHandler(t *testing.T){
	body := strings.NewReader(`testK=testV`)
	req, err := http.NewRequest("PUT","http://httpbin.org/put",body)
	if err != nil {
		t.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)
	
	body2 := strings.NewReader(`testK=testV`)
	testReq, err := http.NewRequest("PUT","/put",body2)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(putHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin","json")
	expectedResult = deleteJSONval(expectedResult,"url","origin","json")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}

}
func TestAnythingHandler(t *testing.T){
	body := strings.NewReader(`testK=testV`)
	req, err := http.NewRequest("POST","http://httpbin.org/anything",body)
	if err != nil {
		t.Fatal(err)
	}
	resp,err := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)
	
	body2 := strings.NewReader(`testK=testV`)
	testReq, err := http.NewRequest("POST","/anything",body2)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(anythingHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin","json")
	expectedResult = deleteJSONval(expectedResult,"url","origin","json")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}

}

func TestEncodingUtfHandler(t *testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/encoding/utf8",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/encodindg/utf8",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(utf8Handler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func  TestGzipHandler( t* testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/gzip",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/gzip",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(gzipHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

/*func  TestDeflateHandler( t* testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/deflate",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/deflate",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(deflateHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}*/
/*func  TestBrotliHandler( t* testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/brotli",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/brotli",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(brotliHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}*/
func TestStatusHandler( t *testing.T){
	req, err := http.NewRequest("GET","http://httpbin.org/status/418",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	
	testReq, err := http.NewRequest("GET","/status/418",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(statusHandler)
	handler.ServeHTTP(resprec,testReq)
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resprec.Code, resp.Status)
	}
}