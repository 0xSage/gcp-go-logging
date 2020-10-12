// Can't be run, needs to deloy to GCF
// Package p contains an HTTP Cloud Function.
package main

import (
        "log"
        "context"
        "cloud.google.com/go/logging"
)

func main() {
  hello()
}
// To deploy:
// gcloud builds submit --tag gcr.io/go-logging/run
// gcloud run deploy --image gcr.io/go-logging/run --platform managed
func hello() {
  ctx := context.Background()
  projectID := "go-logging"

  // Create a logging client
  client, err := logging.NewClient(ctx, projectID)
  if err != nil {
    log.Fatalf("Failed to create logging client: %V", err)
  }

  defer client.Close()
  logger := client.Logger("gcf-log")

  log.Print("hello world!")
  logger.Log(logging.Entry{Payload: "Hello world!"})
  logger.StandardLogger(logging.Info).Println("Hello World!")
}
