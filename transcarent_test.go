package main

import (
    "testing"
    "strings"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

type expected struct {
    Name string
    Username string
    Email string
}

func TestUserPage(t *testing.T) {
    r, _ := http.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()

    testTable := []struct {
        name string
        mockFunc func()
        expectedResponse []byte
        expectedErr error
    }{
        {
            name: "successful-request",
            mockFunc: func() {
                sendRequestFunc = func(w http.ResponseWriter, q string) ([]byte, error) {
                    var Expected = &expected {
                        Name: "Leanne Graham",
                        Username: "Bret",
                        Email: "Sincere@april.biz",
                    }
                    jsonObj, _ := json.Marshal(Expected)
                    return jsonObj, nil
                }
            },
            expectedResponse: []byte(`{"userinfo":{"name":"Leanne Graham","username":"Bret","email":"Sincere@april.biz"},"posts":null}`),
            expectedErr: nil,
        },
        {
            name: "id-too-low",
            mockFunc: func() {
                sendRequestFunc = func(w http.ResponseWriter, q string) ([]byte, error) {
                    var Expected = &expected {
                        Name: "Should Fail",
                        Username: "With Error",
                        Email: "idbetween1and10@gmail.com",
                    }
                    jsonObj, _ := json.Marshal(Expected)
                    return jsonObj, nil
                }
            },
            expectedResponse: []byte(`{"message":"ID must be between 1 and 10."}500 - ID of user must be between 1 and 10.`),
            expectedErr: nil,
        },
        {
            name: "id-too-high",
            mockFunc: func() {
                sendRequestFunc = func(w http.ResponseWriter, q string) ([]byte, error) {
                    var Expected = &expected {
                        Name: "Should Fail",
                        Username: "With Error",
                        Email: "idbetween1and10@gmail.com",
                    }
                    jsonObj, _ := json.Marshal(Expected)
                    return jsonObj, nil
                }
            },
            expectedResponse: []byte(`{"message":"ID must be between 1 and 10."}500 - ID of user must be between 1 and 10.`),
            expectedErr: nil,
        },
    }

    for idx, tc := range testTable {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            switch {
            case idx == 0:
                vars := map[string]string {
                    "id": "1",
                }
                r = mux.SetURLVars(r, vars)
                break
            case idx == 1:
                vars := map[string]string {
                    "id": "0",
                }
                r = mux.SetURLVars(r, vars)
                break
            case idx == 2:
                vars := map[string]string {
                    "id": "11",
                }
                r = mux.SetURLVars(r, vars)
                break
            }

            userPage(rr, r)

            assert.Equal(t, http.StatusOK, rr.Code)
            bodyBytes, _ := ioutil.ReadAll(rr.Body)
            newString := strings.Replace(string(bodyBytes), "\n", "", -1)
            assert.Equal(t, newString, string(tc.expectedResponse))
        })
    }
}
