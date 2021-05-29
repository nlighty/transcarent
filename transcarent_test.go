package main

import (
    "errors"
    "reflect"
    "testing"
    "encoding/json"
    "net/http"
    "net/http/httptest"
)

func TestHttpCall(t *testing.T) {
    var w http.ResponseWriter
    var u user
    testTable := []struct {
        name string
        server *httptest.Server
        expectedResponse *user
        expectedErr error
    }{
        {
            name: "successful-request",
            server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"name": "Leanne Graham", "username": "Bret", "email": "Sincere@april.biz"}`))
            })),
            expectedResponse: &user{
                Name: "Leanne Graham",
                Username: "Bret",
                Email: "Sincere@april.biz",
            },
            expectedErr: nil,
        },
    }
    for _, tc := range testTable {
        t.Run(tc.name, func(t *testing.T) {
            defer tc.server.Close()
            resp, err := sendRequest(w, tc.server.URL)
            json.Unmarshal(resp, &u)
            if !reflect.DeepEqual(&u, tc.expectedResponse) {
                t.Errorf("expected (%v), got (%v)", tc.expectedResponse, u)
            }
            if !errors.Is(err, tc.expectedErr) {
                t.Errorf("expected (%v), got (%v)", tc.expectedErr, err)
            }
        })
    }
}

