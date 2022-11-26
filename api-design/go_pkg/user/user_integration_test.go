//go:build integration

package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
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
	// req.Header.Add("Content-Type", "application/json")
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

func TestGetAllUser(t *testing.T) {
	seedUser(t)

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

func seedUser(t *testing.T) User {
	var c User
	body := bytes.NewBufferString(`{
		"name": "sing",
		"age": 25
	}`)

	if err := request(http.MethodPost, uri("users"), body).Decode(&c); err != nil {
		t.Fatal("can't create user", err)
	}
	return c
}

func TestCreateUserHandler(t *testing.T) {
	var c User
	body := bytes.NewBufferString(`{
		"name": "sing",
		"age": 25
	}`)

	res := request(http.MethodPost, uri("users"), body)
	err := res.Decode(&c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, c.ID)

}

func TestUserByID(t *testing.T) {
	u := seedUser(t)

	var user User
	res := request(http.MethodGet, uri("users", strconv.Itoa(u.ID)), nil)
	err := res.Decode(&user)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, u.ID, user.ID)
}

func TestUpdateUserByID(t *testing.T) {
	u := seedUser(t)
	update := bytes.NewBufferString(`{
		"name": "sing22",
		"age": 25
	}`)

	res := request(http.MethodPut, uri("users", strconv.Itoa(u.ID)), update)

	assert.EqualValues(t, http.StatusNoContent, res.StatusCode)
}

func TestDeleteUserByID(t *testing.T) {
	u := seedUser(t)

	res := request(http.MethodDelete, uri("users", strconv.Itoa(u.ID)), nil)

	assert.EqualValues(t, http.StatusNoContent, res.StatusCode)
}
