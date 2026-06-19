package main

import (
	"fmt"
	"net/http"
	"time"
	"flag"

	femApp "github.com/yogendraatluri990/fem_project_go/internal/app"
)

func main() {
	var port int
	app, err := femApp.NewApplication()
    
	/*
	 @NOTE: The port flag is defined and will passed to flag package to parse the command line arguments. If the flag is not provided, it will default to 8080.
	*/
	flag.IntVar(&port, "onPort", 8080, "Port on which the server will listen")
	flag.Parse()


	if err != nil {
		fmt.Printf("Error initializing application: %v\n", err)
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	serverErr := server.ListenAndServe()

	app.Logger.Printf("Server is up and running on port %d\n",)

	if serverErr != nil {
		app.Logger.Fatal(serverErr)
	}
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health check endpoint hit and this is %s\n", time.Now().Format(time.RFC1123))
}