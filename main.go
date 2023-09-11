package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	http.Handle("/", http.FileServer(http.Dir("./public_html")))
	http.ListenAndServeTLS(":6969", os.Getenv("CERT_PATH"), os.Getenv("PRIVKEY_PATH"), nil)
}
