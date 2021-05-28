/* This project is used to create a web server that will query other external


   APIs. The responses from the respective APIs will be stored into a single
   JSON object that will be served back to the requesting user.
 */

package main

import (
    "errors"
    "fmt"
    "log"
    "strconv"
    "io/ioutil"
    "net/http"
    "encoding/json"

    "golang.org/x/sync/errgroup"

    "github.com/rainycape/memcache"
    "github.com/gorilla/mux"
)

// Memcache client used to store data from previous requests to reduce latency
var mc *memcache.Client

// Global error to set if something goes wrong
var glbErr error

// Struct to hold the user data response from the external API
type user struct {
    Name string `json:"name"`
    Username string `json:"username"`
    Email string `json:"email"`
}

// Struct to hold the posts data response from the external API
type posts struct {
    UserId int `json:"userId"`
    Id int `json:"id"`
    Title string `json:"title"`
    Body string `json:"body"`
}

// Struct to hold the data that will be returned from our API
type response struct {
    User user `json:"userinfo"`
    Posts []posts `json:"posts"`
}

type jsonErrors struct {
    Message string `json:"message"`
}

// This function just serves a basic home page
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page.")
}

// This function is used to query the appropriate external API 
// for data retrieval regarding the user and their posts
func sendRequest(w http.ResponseWriter, q string) (*http.Response, error) {
    var Errors jsonErrors

    resp, err := http.Get(q)
    if err != nil {
        glbErr = err
        log.Println(err)
        Errors.Message = "Error retrieving data from external API."
        json.NewEncoder(w).Encode(Errors)
        return resp, err
    }
    return resp, nil
}

/* userPage function contains the majority of the logical code
   for retrieving data from the appropriate external APIs.
   It takes in an http Response Writer and an http Request.
 */
func userPage(w http.ResponseWriter, r *http.Request) {
    var Errors jsonErrors
    var g errgroup.Group

    w.Header().Set("Content-Type", "application/json") // Set the header to return Json
    var Response response // Structure to contain the response data that will be displayed
    vars := mux.Vars(r) // Reads in the variables from the http request
    id := vars["id"] // Sets a local variable to the value of the id entered by the user
    idVal, _ := strconv.Atoi(id)
    if idVal < 1 || idVal > 10 {
        log.Println("ID must be between 1 and 10.")
        Errors.Message = "ID must be between 1 and 10."
        json.NewEncoder(w).Encode(Errors)
        glbErr = errors.New("ID must be between 1 and 10.")
    }

    listCached, cacheErr := mc.Get(fmt.Sprintf(id)) // Checks to see if the requested data
                                                    // is already cached
    if cacheErr == nil {
        w.Header().Set("cached", string(id)) // Set a key for the data to be cached
        w.Write(listCached.Value) // Write the data to be cached
    } else if cacheErr != memcache.ErrCacheMiss {
        log.Printf("memcached error: %v\n", cacheErr)
    }

    // Set the queryString value to the URL to send the request
    queryString := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)
    g.Go(func() error {
        resp, sendErr := sendRequest(w, queryString) // Send the GET request
        defer resp.Body.Close() // Ensure we are cleaning up after ourselves
        if sendErr != nil {
            return sendErr
        }
        bodyBytes, err := ioutil.ReadAll(resp.Body) // Read the response into a buffer
        if err != nil {
            log.Println(err)
            Errors.Message = "Error reading response from external API."
            json.NewEncoder(w).Encode(Errors)
            glbErr = errors.New("Error reading response from external API.")
            return err
        }
        json.Unmarshal(bodyBytes, &Response.User)
        return nil
    })

    // Set the queryString value for the next URL to send the request
    queryString2 := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%s", id)
    g.Go(func() error {
        resp, sendErr := sendRequest(w, queryString2) // Send the GET request
        defer resp.Body.Close() // Ensure we clean up
        if sendErr != nil {
            return sendErr
        }
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Println(err)
            Errors.Message = "Error reading response from external API."
            json.NewEncoder(w).Encode(Errors)
            glbErr = errors.New("Error reading response from external API.")
            return err
        }
        json.Unmarshal(bodyBytes, &Response.Posts)
        return nil
    })

    waitErr := g.Wait()
    if waitErr != nil {
        Errors.Message = "Received Error"
        json.NewEncoder(w).Encode(Errors)
        glbErr = errors.New("Received Error")
    }

    if glbErr == nil {
        // Encode the json data into the webage
        json.NewEncoder(w).Encode(Response)
    }
    glbErr = nil
}

// Function to route the web requests to the proper pages and start the local server
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/users/{id}", userPage)

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// main function to initiate the web server
func main() {
    var err error
    mc, err = memcache.New("127.0.0.1:11211") // Initiate the memcache
    if err != nil {
        log.Fatal(err)
    }
    handleRequests() // Start the server and prepare for incoming requests
}
