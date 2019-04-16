package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Struct to return the cache items in json format
type CacheStruct struct {
	ID          []string `json:"item,omitempty"`
	Pages       int      `json:"total_pages,omitempty"`
	Description string   `json:"description,omitempty"`
}

/*
Http Handler responsible for accepting the http request to add the item into cache
-User needs to send the request with following details
-EndPoint:POST http://localhost:8080/putcache
-Request Parameters: FormData
-Variable name:"data" value:" a single item or an array of item seperated with comma"
-Example data="abc,def,123,0909"
-Response: Array of items in json format
*/
func PutCache(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/putcache")
	data := r.FormValue("data")
	fmt.Println("form data is", data)
	items := strings.Split(data, ",")
	for _, item := range items {
		err := AddItem(item)
		if err != nil {
			fmt.Fprintf(w, "cache loading failed %s", err.Error()) // send data to client side
		}
	}
	fmt.Fprintf(w, "Data Added Successfully") // send data to client side
}

/*
Http Handler responsible for loading cache from the http request
-User needs to send the request with following details
-EndPoint:GET http://localhost:8080/loadcache
-Request Parameters: nil
-Variable name:nil
-Response: Array of all the items present in json format
*/
func LoadCache(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/loadcache")
	var itemsArray CacheStruct
	items, err := LoadCache_n()
	if err != nil {
		fmt.Fprintf(w, "cache loading failed %s", err.Error()) // send data to client side
	}
	for _, k := range items {
		itemsArray.ID = append(itemsArray.ID, k.(string))
	}
	js, err := json.Marshal(itemsArray)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Server", "A Go Web Server for Cache")
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/*
Http Handler responsible for Returning the cache from the http request
-User needs to send the request with following details
-EndPoint:GET http://localhost:8080/getcache
-Request Parameters: nil
-Variable name:nil
-Response: Array of all the items present in json format
*/
func GetCache(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/getcache")
	var itemsArray CacheStruct
	var pagenumber string
	pagenumber = r.FormValue("page")
	if pagenumber == "" {
		pagenumber = "1"
		fmt.Println("page number not sent")
	}
	fmt.Println("page number selected", pagenumber)
	page, err := strconv.Atoi(pagenumber)
	if err != nil {
		fmt.Fprintf(w, "wrong input for page parameter")
	}
	items, total, err := GetCache_n(page)
	if err != nil {
		itemsArray.Description = err.Error()
		itemsArray.Pages = total
	}
	if len(items) == 0 {
		fmt.Println("GET CACHE FAILED CHECK WEB CONSOLE FOR MORE INFO")
	} else {
		for _, k := range items {
			itemsArray.ID = append(itemsArray.ID, k.(string))
		}
		itemsArray.Pages = total
	}
	js, err := json.Marshal(itemsArray)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Server", "A Go Web Server for Cache")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/*
Http Handler responsible for clearing the cache from the http request
-User needs to send the request with following details
-EndPoint:GET http://localhost:8080/clearcache
-Request Parameters: nil
-Variable name:nil
-Response: Response message in raw string format
*/
func ClearCache(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/clearcache")
	err := ClearCache_n()
	if err != nil {
		w.Header().Set("Server", "A Go Web Server for Cache")
		w.WriteHeader(202)
		w.Write([]byte("Cache was not cleared"))
	}
	w.Header().Set("Server", "A Go Web Server for Cache")
	w.WriteHeader(200)
	w.Write([]byte("Cache cleared Successfully"))
}
