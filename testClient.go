package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	// "math/rand"
	"net/http"
	"time"
)

type Book struct {
	Name       string    `json:"name"`
	Author     string    `json:"author"`
	Pages      int       `json:"pages"`
	Year       int       `json:"year"`
	CreateTime time.Time `json:"createtime"`
}

const BookURL = "http://127.0.0.1:8080/api/0.1/"

// Pring a Book
func (b Book) String() string {
	s := ""
	s += fmt.Sprintln("Name:", b.Name)
	s += fmt.Sprintln("Author:", b.Author)
	s += fmt.Sprintln("Pages:", b.Pages)
	s += fmt.Sprintln("Year:", b.Year)
	s += fmt.Sprintln("CreateTime:", b.CreateTime)
	return s
}

func queryAll() {
	// Send request
	// resp, err := http.Get(BookURL + "queryAll")
	resp, err := http.Get(BookURL + "books")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print status
	fmt.Println(resp.Status, resp.StatusCode)

	// Get body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode body
	var books []Book
	if resp.StatusCode == http.StatusOK {
		// Decode as JSON
		if err := json.Unmarshal(body, &books); err != nil {
			fmt.Println(err, "in decoding JSON")
			return
		}
		for i, v := range books {
			fmt.Println(i, "-------------------------------")
			fmt.Println(v)
		}
		fmt.Println("Total", len(books), "books")
	} else {
		// Decode as text
		fmt.Printf("%s", body)
	}
}

// Return 0: success
// Return 1: failed
func storeTen() int {
	// Send request
	// resp, err := http.Get(BookURL + "storeTen")
	resp, err := http.Post(BookURL+"books", "", nil)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status, resp.StatusCode)
	if resp.StatusCode == http.StatusCreated {
		return 0
	} else {
		return 1
	}
}

// Return 0: success
// Return 1: failed
func deleteAll() int {
	// Send request
	resp, err := http.Get(BookURL + "deleteAll")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status, resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		return 0
	} else {
		return 1
	}
}

func main() {
	if storeTen() != 0 {
		fmt.Println("Store failed")
		return
	} else {
		fmt.Println("Store 10 books")
	}
	// queryAll()
	// if deleteAll() != 0 {
	// 	fmt.Println("Delete failed")
	// 	return
	// } else {
	// 	fmt.Println("Delete all")
	// }
}
