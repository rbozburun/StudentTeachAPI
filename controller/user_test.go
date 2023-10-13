package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rbozburun/StudentTeachAPI/config"
	"github.com/rbozburun/StudentTeachAPI/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func set_CreateUser_router(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	r.POST("/user", CreateUser)

	req, err := http.NewRequest(http.MethodPost, "/user", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil

}

func TestCreateUser(t *testing.T) {
	// Generate new assertion
	a := assert.New(t)

	// Setup the DB
	config.Connect()
	db := config.GetDB()
	config.DB = db

	input_model := models.UserTest{
		Username: "jdoe123",
		Name:     "John",
		Surname:  "Doe",
		Email:    "j@gmail.com",
		RoleID:   1,
	}

	reqBody, err := json.Marshal(input_model)
	if err != nil {
		a.Error(err)
	}

	req, w, err := set_CreateUser_router(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	// Check request method
	a.Equal(http.MethodPost, req.Method, "HTTP Request error: Unexpected Request Method!")

	// Read the response body
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	// Get the response model and compare it
	actual := models.UserTest{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := input_model
	a.Equal(expected, actual)
	db.Exec("DELETE FROM users WHERE id = (SELECT MAX(id)FROM users)")

}
