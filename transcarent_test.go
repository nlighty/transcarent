package main

import (
    "errors"
    "reflect"
    "testing"
    "encoding/json"
    "net/http"
)

func TestHttpCall(t *testing.T) {
    var w http.ResponseWriter
    var u user

    testTable := []struct {
        name string
        mockFunc func()
        expectedResponse *user
        expectedErr error
    }{
        {
            name: "successful-request",
            mockFunc: func() {
                sendRequestFunc = func(w http.ResponseWriter, q string) ([]byte, error) {
                    /*var User = &user {
                        Name: "Leanne Graham",
                        Username: "Bret",
                        Email: "Sincere@april.biz",
                    }*/
                    b := make([]byte, 5, 5)
                    return b, nil
                }
            },
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
            tc.mockFunc()
            queryString := "https://jsonplaceholder.typicode.com/users/1"
            //userPage(w, &r)
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
