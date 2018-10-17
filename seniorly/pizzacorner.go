package main

import (
	//	"encoding/csv"
	//	"encoding/json"
	"fmt"
	//	"io"
	"log"
	"net/http"
	"time"

	"os"
	"strings"
	//	"github.com/gocarina/gocsv"
)

/*struct to unmarshall the csv values*/
type CsvLine_csv struct {
	Key   string `csv:"key,omitempty"`
	Value string `csv:"value,omitempty"`
}

/* Universal map to maintain user order status based on phone number */
var User_order_status map[string]string = make(map[string]string)

func main() {
	/*creating a text file to write the user noyification details */
	file, err := os.Create("User_Notifications.txt")
	if err != nil {
		fmt.Println("could not create file")
	}
	http.HandleFunc("/buy_pizza", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		request_form := r.Form
		fmt.Println("Taking Pizza Order...", request_form.Get("name"))
		phone := request_form.Get("phone")
		pizza_title := request_form.Get("title")
		name := request_form.Get("name")
		if strings.EqualFold(pizza_title, "Veggie Lovers") || strings.EqualFold(pizza_title, "Meat Lovers") {
			time.Sleep(10 * time.Second)
			w.WriteHeader(http.StatusOK)
			delivery_time := time.Now().Add(30 * time.Second).Local()
			User_order_status[phone] = "Order under process will dispatch on or before: " + delivery_time.String()
			response := "Hi " + name + " Thanks for ordering Pizza with us\nOrder deatils: " + pizza_title + "\nPhone: " + phone + "\nEstimated dispatch by:" +
				delivery_time.String() + "\nYour Order is taken and processing, we will send an SMS notification once your order is ready"
			w.Write([]byte(response))
			go sendNotification(name, phone, pizza_title, file)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Selected Pizza is not available\n sorry for the Inconvinience..."))
		}
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		params, ok := r.URL.Query()["phone"]

		if !ok || len(params[0]) < 1 {
			log.Println("Url Param 'Phone' is missing")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Url Parameter 'Phone' is missing"))
			return
		}
		_, present := User_order_status[params[0]]
		if !present {
			fmt.Println("order not found")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No status found for Phone number: " + params[0]))
		} else {
			fmt.Println("Hi ", params[0], " is", User_order_status[params[0]])
			// Query()["phone"] will return an array of items,
			// we only want the single item.
			w.WriteHeader(http.StatusFound)
			w.Write([]byte(User_order_status[params[0]]))
		}
	})
	fmt.Println("Pizza server listening @ port 8080")
	http.ListenAndServe(":8080", nil)
}
func sendNotification(name, phone, pizza_title string, file_head *os.File) {
	time.Sleep(30 * time.Second)
	fmt.Printf("\nAn SMS has been sent to customer with order details\nPhone:'%s'\nname:'%s'", phone, name)
	User_order_status[phone] = "Your " + pizza_title + " is ready and can be picked up at our Store"
	fmt.Println("\nstatus: ", User_order_status[phone])
	stringbyte := "Time " + time.Now().Local().String() + "\n[contact: " + phone + "] [Message: Hi " + name + " Your order for " + pizza_title + " is ready to pickup from our Store]\n"
	file_head.WriteString(stringbyte)
}
