package main

import (
    "fmt"
    "log"
    "net/http"
    "os" 
)

func main() {
    port := "3001"
    if os.Getenv("HTTP_PLATFORM_PORT") != "" {
        port = os.Getenv("HTTP_PLATFORM_PORT")
    }
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, `
        <html>
            <body>
                <h1>Hello from Go!</h1>
                <br />
            </body>
        </html>`)        
    })
    
    log.Println("--->   UP @ " + port +"  <------")   
    http.ListenAndServe(":"+port, nil)
}

