package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("By default this application runs on port 8080\nIncase if you want to run it in a different port tou can do that while running the docker image")
	fmt.Println("Server started...")
	fmt.Println("***************************************************")
	fmt.Println("Http Requests for reference")
	fmt.Println("Loadcache:", "Method:GET ", "http:IP:Port/loadcache")
	fmt.Println("Getcache:", "Method:POST/GET ", "http:IP:Port/getcache?page=1")
	fmt.Println("Putcache:", "Method:POST ", "http:IP:Port/putcache?data=test1,test2,tets3")
	fmt.Println("Clearcache:", "Method:GET ", "http:IP:Port/clearcache")
	fmt.Println("***************************************************")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to test website!")
	})

	http.HandleFunc("/putcache", PutCache)     // Adding elements into cache
	http.HandleFunc("/loadcache", LoadCache)   // Adding bulk elements into cache
	http.HandleFunc("/getcache", GetCache)     // Adding bulk elements into cache
	http.HandleFunc("/clearcache", ClearCache) //clearing all the cache

	http.ListenAndServe(":8080", nil) //This application by default will be running in port 8080
}
