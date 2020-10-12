// Bucket name: go-logging

// So vm treats my code as executable program
package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "context"
        "cloud.google.com/go/logging"
)

//                assembles http response    obj/data structure of incoming request
func indexHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w,r)
    return
  }
  fmt.Fprint(w, "Hello World!")
}

func main() {
  ctx := context.Background()
  projectID := "go-logging"

  // Create a logging client
  client, err := logging.NewClient(ctx, projectID)
  if err != nil {
    log.Fatalf("Failed to create logging client: %V", err)
  }

  defer client.Close()
  // Sets name of log to write to
  logger := client.Logger("gce-log")

  log.Print("hello world!")
  logger.Log(logging.Entry{Payload: "Hello world!"})
  logger.StandardLogger(logging.Info).Println("Hello World!")

  // Handle routes & ports
  http.HandleFunc("/", indexHandler)
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  logger.Log(logging.Entry{Payload: "Listening on port: 8080"})

  // Start server
  if err := http.ListenAndServe(":"+port, nil); err != nil {  
    log.Fatal(err)
  }
}
