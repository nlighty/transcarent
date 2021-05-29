package main

import (
    "errors"
    "reflect"
    "testing"
    "encoding/json"
    "net/http"
)

type CallAPI struct{}

func TestHttpCall(t *testing.T) {
    var w http.ResponseWriter
    var u user
    testTable := []struct {
        name string
        expectedResponse *user
        expectedErr error
    }{
        {
            name: "successful-request",
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
            queryString := "https://jsonplaceholder.typicode.com/users/1"
            sendRequestFunc = func(w http.ResponseWriter, q string) ([]byte, error) {
                return byte[]("Example Return"), nil
            }
            resp, err := sendRequestFunc(w, queryString)
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
