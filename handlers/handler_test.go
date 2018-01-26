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
	"io/ioutil"
)

//server is the test flag.
var server string

//init function is the constructor of package for flag.
func init(){
	flag.StringVar(&server,"server","localhost:8080","server flag")	
}
	
func TestIpHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/ip",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(ipHandler)
	handler.ServeHTTP(resprec,req)
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}
}
func TestHeadersHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/headers",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

func TestGetHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/get",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}

func TestIndexHandler(t *testing.T){
	flag.Parse()
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
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/user-agent",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",expectedResult, result)
	}
}

func TestPostHandler(t *testing.T){
	flag.Parse()
	body := strings.NewReader(`testK=testV`)
	req, err := http.NewRequest("POST",server+"/post",body)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}

}

func TestDeleteHandler(t *testing.T){
	flag.Parse()
	form := url.Values{}
	form.Add("test1K","test1V")
	form.Add("test2K","test2v")
	req, err := http.NewRequest("DELETE",server+"/delete",strings.NewReader(form.Encode()))
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}

}
func TestPutHandler(t *testing.T){
	flag.Parse()
	body := strings.NewReader(`testK=testV`)
	req, err := http.NewRequest("PUT",server+"/put",body)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}

}
func TestAnythingHandler(t *testing.T){
	flag.Parse()
	body := strings.NewReader(`testK=testV`)
	req, err := http.NewRequest("POST",server+"/anything",body)
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
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/encoding/utf8",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func  TestGzipHandler( t* testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/gzip",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

/*func  TestDeflateHandler( t* testing.T){
	req, err := http.NewRequest("GET",server+"/deflate",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
/*func  TestBrotliHandler( t* testing.T){
	req, err := http.NewRequest("GET",server+"/brotli",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}*/
func TestStatusHandler( t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/status/418",nil)
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
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}
}
func TestResponseHeadersHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/response-headers",nil)
	if err != nil {
		t.Fatal(err)
	}
	req.URL.Query().Add("testKey","testValue")	
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/response-headers",nil)
	if err != nil {
		t.Fatal(err)
	}
	testReq.URL.Query().Add("testKey","testValue")	
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(responseHeaderHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.StatusCode,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
	
func TestRedirectMultiHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/redirect/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/redirect/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(redirectMultiHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

/*func TestRedirectToHandler(t *testing.T){
	req, err := http.NewRequest("GET",server+"/redirect-to",nil)
	if err != nil {
		t.Fatal(err)
	}
	req.URL.Query().Add("url","https://httpbin.org/")
	req.URL.Query().Add("status_code","302")
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/redirect-to",nil)
	if err != nil {
		t.Fatal(err)
	}
	testReq.URL.Query().Add("url","https://httpbin.org/")
	testReq.URL.Query().Add("status_code","302")
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(redirectToHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}*/
func TestCookiesHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/cookies",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/cookies",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(cookieHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestCookiesSetDelHandler(t *testing.T){
	flag.Parse()
	//Set
	req, err := http.NewRequest("GET",server+"/cookies/set",nil)
	if err != nil {
		t.Fatal(err)
	}
	req.URL.Query().Add("testKey","testValue")	
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/cookies/set",nil)
	if err != nil {
		t.Fatal(err)
	}
	testReq.URL.Query().Add("testKey","testValue")	
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(cookieSetDelhandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
	//Del
	reqDel, err := http.NewRequest("GET",server+"/cookies/delete",nil)
	if err != nil {
		t.Fatal(err)
	}
	req.URL.Query().Add("testKey","testValue")	
	respDel,_ := http.DefaultClient.Do(reqDel)
	expectedResultDel ,_:= ioutil.ReadAll(resp.Body)

	testReqDel, err := http.NewRequest("GET","/cookies/delete",nil)
	if err != nil {
		t.Fatal(err)
	}
	testReq.URL.Query().Add("testKey","testValue")	
	resprecDel := httptest.NewRecorder()
	handlerDel := http.HandlerFunc(cookieSetDelhandler)
	handlerDel.ServeHTTP(resprec,testReqDel)//host url will be ignored while comparing
	
	resultDel := deleteJSONval(resprecDel.Body.Bytes(),"url","origin")
	expectedResultDel = deleteJSONval(expectedResultDel,"url","origin")
	
	//Status Check	
	if resprecDel.Code != respDel.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",respDel.Status,resprecDel.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(resultDel) != string(expectedResultDel) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResultDel), string(resultDel))
	}
}

func TestBasicAuthHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/basic-auth/testID/testPW",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/basic-auth/testID/testPW",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(basicAuthHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestHiddenBasicAuthHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/hidden-basic-auth/testID/testPW",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/hidden-basic-auth/testID/testPW",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(hiddenBasicAuthHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestStreamHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/stream/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/stream/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(streamHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestDelayHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/delay/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/delay/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(delayHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestHtmlHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/html",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/html",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(htmlHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

func TestRobotsTextHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/robots.txt",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/robots.txt",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(robotsTextHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestDenyHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/deny",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/deny",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(denyHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestCacheHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/cache",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/cache",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(cacheHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")

	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestCacheControlHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/cache/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/cache/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(cacheControlHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestBytesHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/bytes/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/bytes/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(bytesHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestLinkHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/links/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/links/10",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(linkHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

func TestImageHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/image",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/image",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(imageHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestImagePNGHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/image/png",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/image/png",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(pngHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestImageJPEGHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/image/jpeg",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/image/jpeg",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(jpegHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestImageWEBPHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/image/webp",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/image/webp",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(webpHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}

func TestImageSVGHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/image/svg",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/image/svg",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(svgHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestFormsPostHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/forms/post",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/forms/post",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(formsHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}
func TestXMLHandler(t *testing.T){
	flag.Parse()
	req, err := http.NewRequest("GET",server+"/xml",nil)
	if err != nil {
		t.Fatal(err)
	}
	resp,_ := http.DefaultClient.Do(req)
	expectedResult ,_:= ioutil.ReadAll(resp.Body)

	testReq, err := http.NewRequest("GET","/xml",nil)
	if err != nil {
		t.Fatal(err)
	}
	resprec := httptest.NewRecorder()
	handler := http.HandlerFunc(xmlHandler)
	handler.ServeHTTP(resprec,testReq)//host url will be ignored while comparing
	
	result := deleteJSONval(resprec.Body.Bytes(),"url","origin")
	expectedResult = deleteJSONval(expectedResult,"url","origin")
	//Status Check	
	if resprec.Code != resp.StatusCode{
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",resp.Status,resprec.Code)
	}

	//Body check.Since body has a header key we don't have to check headers separately
	if string(result) != string(expectedResult) {
		t.Errorf("Unexpected result occurred.\nExpected Result:%v\n Result:%v",string(expectedResult), string(result))
	}
}