package main_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	RH "./handlers"
)

type TestSignInResponse struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	Result string `json:"result"`
}

func TestSignMeIn(t *testing.T) {
	bodyReader := strings.NewReader("{\"username\":\"RPUSERTEST\",\"password\":\"RPUSERtest69!\"}")
	req, err := http.NewRequest("GET", "https://localhost:8080/uniname", bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RH.SignMeInHandlerFunc)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func runLater(resp *http.Response, testSignin TestSignInResponse, re *regexp.Regexp) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString, "this")

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println("{\"name\":\"" + testSignin.Name + "\",\"token\":\"" + re.FindString(testSignin.Token) + "\",\"result\":\"" + testSignin.Result + "\"}")
}

func TestUniqeUsername(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// here we write our expected response, in this case, we return a
		// JSON string which is typical when dealing with REST APIs
		io.WriteString(w, "{ \"name\": \"RPUSERTST\",\"valid\": \"a valid\"}")
	}
	bodyReader := strings.NewReader("{\"username\":\"RPUSERTST\"}")
	req := httptest.NewRequest("GET", "https://localhost:8080/uniname", bodyReader)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	handler_invalid := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"name\": \"RPUSERTEST\",\"valid\": \"an valid\"}")
	}
	bodyReader_invalid := strings.NewReader("{\"username\":\"RPUSERTST\"}")
	req_invalid := httptest.NewRequest("GET", "https://localhost:8080/uniname", bodyReader_invalid)
	w_invalid := httptest.NewRecorder()
	handler_invalid(w_invalid, req_invalid)
	resp_invalid := w_invalid.Result()
	body_invalid, _ := ioutil.ReadAll(resp_invalid.Body)
	fmt.Println(resp_invalid.StatusCode)
	fmt.Println(resp_invalid.Header.Get("Content-Type"))
	fmt.Println(string(body_invalid))
}
