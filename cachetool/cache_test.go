package main

import (
	"fmt"
	"testing"
)

/*
Note: Instead of a database here i have used the file operations
      to make this application lighter and fast.
*/
func TestCacheTool(t *testing.T) {
	//Step1: Loading the cache from the file instead of a database
	items, err := LoadCache_n()
	if err != nil {
		fmt.Println("cache loading failed", err) // send data to client side
	} else {
		fmt.Println("items are", items)
	}
	//Step2: Adding an item into the cache array and the file as well
	err = AddItem("name:virat age:30 mobile:5654654667,name:anushka age:31 mobile:43543564364")
	if err != nil {
		fmt.Println("failed adding data into cache")
	} else {
		fmt.Println("Data added successfully")
	}
	//Get the latest cache in json format
	items, total, err := GetCache_n(1)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("items:", items, " total number of pages:", total)
	}
	//Clear all the cache  items and empty the file
	//	err = ClearCache_n()
	//	if err != nil {
	//		fmt.Println("error while clearing cache", err)
	//	} else {
	//		fmt.Println("Data cleared from cache")
	//	}
	//Reload recomended after clearing the cache
	//fmt.Println(models.ReloadCache())
}
