/*------------Package consists of all the operations related to cache--------------*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

//A global variable to store the cache
var Pcache []interface{}

/*
Function to add the item into the cache
This function accepts items of any data type and puts the same into cache
-This function returns error if adding fails
*/
func AddItem(item interface{}) error {
	Pcache = append(Pcache, item)
	//	d1 := []byte(string(item.(string)))
	f, err := os.OpenFile("cache.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	if _, err = f.WriteString("\n" + item.(string)); err != nil {
		panic(err)
	}
	defer f.Close()
	//fmt.Println("added", Pcache)
	return nil
}

/*
Function to get the latest cache
-If no Page number specified then starting 5 items would be sent
-This Function returns cache items and the number pages available for pagination
*/
func GetCache_n(pagenumber int) ([]interface{}, int, error) {
	if len(Pcache) == 0 {
		return nil, 0, errors.New("No items present in cache in now")
	}
	total := float64(len(Pcache)) / float64(5)
	total_pages := int(math.Ceil(total))
	//fmt.Println("total pages", total_pages)
	if pagenumber*5 > len(Pcache) && pagenumber <= total_pages {
		startpage := (pagenumber - 1) * 5
		return Pcache[startpage:], total_pages, nil
	} else if pagenumber > total_pages {
		return nil, total_pages, errors.New("Page not found")
	} else {
		startpage := (pagenumber - 1) * 5
		endpage := startpage + 5
		return Pcache[startpage:endpage], total_pages, nil
	}
	return Pcache, total_pages, nil
}

/*
Function to load the cache after clear
-This function returns all the cache items
*/
func ReloadCache_n() ([]interface{}, error) {
	file, err := os.Open("cache.txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		Pcache = append(Pcache, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return Pcache, nil
}

/*
Function to load the cache for the first instance
-This function returns all the cache items
*/
func LoadCache_n() ([]interface{}, error) {
	file, err := os.Open("cache.txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		Pcache = append(Pcache, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return Pcache, nil
}

/*
Function to clear all the cache
-This function returns error if clear fails
*/
func ClearCache_n() error {
	err := os.Truncate("cache.txt", 0)
	if err != nil {
		fmt.Printf("error truncating cache: %v", err)
		return err
	}
	Pcache = Pcache[:0]
	return nil
}
