package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//get the value of the ADDR environment variable
	//and use that as the address this server will listen on
	addr := os.Getenv("ADDR")
	//if not set, default to ":443", which means listen for
	//all requests to all hosts on port 443
	if len(addr) == 0 {
		addr = "localhost:4242"
	}

	//fmt.Println("Hello World!")
	mux := http.NewServeMux()

	mux.HandleFunc("/compliment", ComplimentHandler)

	fmt.Printf("server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

}
