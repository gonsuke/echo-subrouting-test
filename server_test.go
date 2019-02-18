package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParentRoute(t *testing.T) {
	client := new(http.Client)
	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/hello", nil)

	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, Hello, string(respBody))
}

func TestSubRoute1(t *testing.T) {
	client := new(http.Client)
	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/hello2", nil)

	resp, _ := client.Do(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestSubRoute2(t *testing.T) {
	client := new(http.Client)
	router := NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/subroute2/hello3", nil)

	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, Hello, string(respBody))
}
