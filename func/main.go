// Can't be run, needs to deloy to GCF
// Package p contains an HTTP Cloud Function.
package p

import (
        "net/http"
        "log"
        "context"
        "cloud.google.com/go/logging"
)

// To deploy:
// gcloud builds submit --tag gcr.io/go-logging/run
// gcloud run deploy --image gcr.io/go-logging/run --platform managed
func HelloWorld(w http.ResponseWriter, r *http.Request) {
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
