package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", "Basic YWRtaW46MTIzNA==")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

func TestGetAllUser(t *testing.T) {
	var c User
	body := bytes.NewBufferString(`{
		"name": "sing",
		"age": 25
	}`)

	if err := request(http.MethodPost, uri("users"), body).Decode(&c); err != nil {
		t.Fatal("can't create user", err)
	}

	var users []User
	res := request(http.MethodGet, uri("users"), nil)
	err := res.Decode(&users)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(users), 0)
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}
	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}
