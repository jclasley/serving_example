package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// -----------
// Try running each function once
// Feel free to play around with slashes, directory configs, whatever
// Doesn't seem to change anything
// -----------


func main() {

	// noStrictSlash()
	// First access `http://localhost:8080/home` and note that it says `JS has not loaded` -- this is the default HTML
	// Next add a trailing slash, so it's `http://localhost:8080/home/` -- JS has worked???

	// strictSlash()
	// Changes nothing

	// vanillaServe()
	// this just doesn't find the JS file at all -- check `Sources` tab under the chrome inspector

	ginFromRoot()
	// note the errors

	//ginFromHome()
}

func strictSlash() {
	mux := mux.NewRouter().StrictSlash(true)
	mux.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir("./client/"))))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func noStrictSlash() {
	mux := mux.NewRouter().StrictSlash(true)
	mux.PathPrefix("/home/").Handler(http.StripPrefix("/home/", http.FileServer(http.Dir("./client"))))
	// note above we have changed the trailing slashes around, nothing changes
	// this makes it so `http://localhost:8080/home` is now a 404 so neither the HTML nor the JS is served
	// `http://localhost:8080/home/` still works
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func vanillaServe() {
	http.Handle("/home", http.StripPrefix("/home", http.FileServer(http.Dir("./client/")))) // note no trailing slash
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ginFromRoot() {
	r := gin.Default()
	r.StaticFS("/", http.Dir("./client"))
	r.GET("/home", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "receieved",
		})
	})
	r.Run(":8080")
}

func ginFromHome() {
	r := gin.Default()
	r.StaticFS("/home", http.Dir("./client"))
	r.GET("/", func (c *gin.Context) {
		c.Redirect(http.StatusFound, "/home")
	})
	r.Run(":8080")
}