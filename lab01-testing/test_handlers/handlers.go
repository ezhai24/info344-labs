package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func ComplimentHandler(w http.ResponseWriter, r *http.Request) {
	compliments := []string{"wonderful", "strong",
		"beautiful", "smart", "loyal", "friendly",
		"good", "asynchronous", "savvy", "handsome",
		"meme-y", "athletic", "intelligent", "good at Golang",
		"creative", "groundbreaking", "artistic", "suave",
		"capable", "bodacious"}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	name := r.FormValue("name")

	// randomly select a compliment from the Compliments
	rand.Seed(time.Now().Unix())
	compliment := compliments[rand.Intn(len(compliments))]
	fmt.Fprintf(w, "%s, you are %s", name, compliment)

}
