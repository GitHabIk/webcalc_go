package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GitHabIk/webcalc_go/internal/application"
)

func main() {
	http.HandleFunc("/api/v1/calculate", application.HandleCalculate)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
