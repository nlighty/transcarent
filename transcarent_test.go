package main

import (
    "testing"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

type expected struct {
    Name string `json:"name"`
    Username string `json:"username"`
    Email string `json:"email"`
}

type officialResponse struct {
    Userinfo user `json:"userinfo"`
    Posts []posts `json:"posts"`
}

func TestHttpCall(t *testing.T) {
    r, _ := http.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()

    vars := map[string]string {
        "id": "1",
    }
    r = mux.SetURLVars(r, vars)

    testTable := []struct {
        name string
        mockFunc func()
        expectedResponse officialResponse
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
            expectedResponse: officialResponse {
                Userinfo: user {
                    Name: "Leanne Graham",
                    Username: "Bret",
                    Email: "Sincere@april.biz",
                },
            },
            expectedErr: nil,
        },
    }

    for _, tc := range testTable {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            userPage(rr, r)

            assert.Equal(t, http.StatusOK, rr.Code)
            b, _ := json.Marshal(tc.expectedResponse)
            bodyBytes, _ := ioutil.ReadAll(rr.Body)
            assert.Equal(t, b, bodyBytes)
        })
    }
}
