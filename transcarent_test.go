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
    Name string
    Username string
    Email string
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
            expectedErr: nil,
        },
    }

    for _, tc := range testTable {
        t.Run(tc.name, func(t *testing.T) {
            tc.mockFunc()
            userPage(rr, r)

            assert.Equal(t, http.StatusOK, rr.Code)
            fmt.Printf("\n\n%v\n\n", rr.Body)
        })
    }
}
