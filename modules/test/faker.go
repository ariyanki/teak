package test

import (
	"net/http"
)

var (
	UserJSON             = `{"username":"teak","password":"teak"}`
	//DummyHandlerLogin DummyHandlerLogin
	DummyHandlerLogin = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body := `{
				"token": "a.b.c"
			}`
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		}
	}
	//DummyHandlerProcess DummyHandlerProcess
	DummyHandlerProcess = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body := `{"body":""}`
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		}
	}
	//DummyHandlerProcessNot200 DummyHandlerProcessNot200
	DummyHandlerProcessNot200 = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.WriteHeader(http.StatusForbidden)
		}
	}
	//DummyHandlerProcessReadBodyError DummyHandlerProcessReadBodyError
	DummyHandlerProcessReadBodyError = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.WriteHeader(http.StatusOK)
		}
	}
)
