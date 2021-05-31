package main

import (
    "fmt"
    "testing"
    "encoding/json"
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
        expectedResponse *response
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
            expectedResponse: &response {
                User: user {
                    Name: "Leanne Graham",
                    Username: "Bret",
                    Email: "Sincere@april.biz",
                },
                Posts: []posts {},
            },
            expectedErr: nil,
        },
    }

    for _, tc := range testTable {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            userPage(rr, r)

            assert.Equal(t, http.StatusOK, rr.Code)
            fmt.Printf("%v\n\n", rr.Body)
            assert.Equal(t, tc.expectedResponse, rr.Body)
        })
    }
}
