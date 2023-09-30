package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

const staticDir = "./static/"

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, staticDir+"index.html")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"message": "This is a sample API endpoint"}`)
}

func main() {
    router := mux.NewRouter()

    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

    router.HandleFunc("/", homeHandler)
    router.HandleFunc("/api", apiHandler)

    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    fmt.Println("Server is listening on :8080")
    if err := server.ListenAndServe(); err != nil {
        fmt.Println(err)
    }
}
