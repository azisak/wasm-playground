package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Serving at port 9090")
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir(os.Getenv("STATIC_PATH"))))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
