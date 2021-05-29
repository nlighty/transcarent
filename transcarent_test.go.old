package main

import (
    "fmt"
    "log"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "testing"
)

func TestQuery(t *testing.T) {
    var Response response
    cases := []response {
        {
            User: user {
                Name: "Leanne Graham",
                Username: "Bret",
                Email: "Sincere@april.biz",
            },
            Posts: []posts {
                {
                    UserId: 1,
                    Id: 1,
                    Title: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
                    Body: "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
                },
                {
                    UserId: 1,
                    Id: 2,
                    Title: "qui est esse",
                    Body: "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
                },
                {
                    UserId: 1,
                    Id: 3,
                    Title: "ea molestias quasi exercitationem repellat qui ipsa sit aut",
                    Body: "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut",
                },
                {
                    UserId: 1,
                    Id: 4,
                    Title: "eum et est occaecati",
                    Body: "ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit",
                },
                {
                    UserId: 1,
                    Id: 5,
                    Title: "nesciunt quas odio",
                    Body: "repudiandae veniam quaerat sunt sed\nalias aut fugiat sit autem sed est\nvoluptatem omnis possimus esse voluptatibus quis\nest aut tenetur dolor neque",
                },
                {
                    UserId: 1,
                    Id: 6,
                    Title: "dolorem eum magni eos aperiam quia",
                    Body: "ut aspernatur corporis harum nihil quis provident sequi\nmollitia nobis aliquid molestiae\nperspiciatis et ea nemo ab reprehenderit accusantium quas\nvoluptate dolores velit et doloremque molestiae",
                },
                {
                    UserId: 1,
                    Id: 7,
                    Title: "magnam facilis autem",
                    Body: "dolore placeat quibusdam ea quo vitae\nmagni quis enim qui quis quo nemo aut saepe\nquidem repellat excepturi ut quia\nsunt ut sequi eos ea sed quas",
                },
                {
                    UserId: 1,
                    Id: 8,
                    Title: "dolorem dolore est ipsam",
                    Body: "dignissimos aperiam dolorem qui eum\nfacilis quibusdam animi sint suscipit qui sint possimus cum\nquaerat magni maiores excepturi\nipsam ut commodi dolor voluptatum modi aut vitae",
                },
                {
                    UserId: 1,
                    Id: 9,
                    Title: "nesciunt iure omnis dolorem tempora et accusantium",
                    Body: "consectetur animi nesciunt iure dolore\nenim quia ad\nveniam autem ut quam aut nobis\net est aut quod aut provident voluptas autem voluptas",
                },
                {
                    UserId: 1,
                    Id: 10,
                    Title: "optio molestias id quia eum",
                    Body: "quo et expedita modi cum officia vel magni\ndoloribus qui repudiandae\nvero nisi sit\nquos veniam quod sed accusamus veritatis error",
                },
            },
        },
    }

    for idx, tc := range cases {
        queryString := fmt.Sprintf("http://127.0.0.1:10000/users/%d", idx + 1)
        fmt.Println(queryString)
        resp, err := http.Get(queryString)
        if err != nil {
            log.Fatal(err)
        }
        defer resp.Body.Close()

        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }

        json.Unmarshal(bodyBytes, &Response)
        fmt.Printf("USER: %v\n", tc.User)
        fmt.Printf("RESPONSE: %v\n", Response.User)

        if Response.User != tc.User {
            t.Error("The user info did not match the expected result")
        }
        for idx, _ := range Response.Posts {
            fmt.Printf("POST: %v\n", tc.Posts[idx])
            fmt.Printf("RESPONSE: %v\n\n", Response.Posts[idx])
            if tc.Posts[idx] != Response.Posts[idx] {
                t.Error("The posts did not match the expected result")
            }
        }
    }
}

func TestErrors(t *testing.T) {
    var Errors jsonErrors
    cases := jsonErrors {
            Message: "ID must be between 1 and 10.",
    }

    resp, err := http.Get("http://127.0.0.1:10000/users/11")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    err2 := json.Unmarshal(bodyBytes, &Errors)
    if err2 != nil {
        log.Fatal(err2)
    }

    if Errors.Message != cases.Message {
        fmt.Printf("Errors.Message: %v\n", Errors.Message)
        fmt.Printf("cases.Message: %v\n", cases.Message)
       t.Error("The error info did not match the expected result")
    }
}
