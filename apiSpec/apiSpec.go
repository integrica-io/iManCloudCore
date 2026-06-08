package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
   	http.Handle("/", http.FileServer(http.Dir("./")))
    port := ":8080"
    slog.Info(fmt.Sprintf("Starting Server on port %s", port))

    // Start server on port specified above
    if err := http.ListenAndServe(port, nil); err != nil {
    	slog.Error("listenAndServe", "error", err)
    }
}