package main

import "log"
import "fmt"
import "net/http"
import "requestModel"
import "github.com/gorilla/mux"

func getStocks(responseWrite http.ResponseWriter, request *http.Request) {
    modelObject = new(RequestModel)
    modelObject.Init()
    ParseFromRequest(modelObject, request)
}

func setStocks(w http.ResponseWriter, r *http.Request) {

    fmt.Println("trrgrbrbbh")
}

func bulkStocks(w http.ResponseWriter, r *http.Request) {

    fmt.Println("trrgrbrbbh")
}

func main() {
    
    routes := mux.NewRouter()
    
    routes.HandleFunc("/app/v1/getstocks", getStocks).Methods("GET")
    routes.HandleFunc("/app/v1/setstocks", setStocks).Methods("PUT")
    routes.HandleFunc("/app/v1/setbulkstocks", bulkStocks).Methods("POST")
    
    err := http.ListenAndServe(":8088", routes)
    
    if err != nil{
        log.Fatal(err)
    }
}
