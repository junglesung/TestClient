package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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
const BookMaxPages = 1000

var BookName = []string{"AAA", "BBB", "CCC", "DDD", "EEE", "FFF", "GGG", "HHH", "III", "JJJ"}
var BookAuthor = []string{"AuthorA", "AuthorB", "AuthorC", "AuthorD", "AuthorE", "AuthorF", "AuthorG", "AuthorH", "AuthorI", "AuthorJ"}

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
	var books map[string]Book = make(map[string]Book)
	if resp.StatusCode == http.StatusOK {
		// Decode as JSON
		if err := json.Unmarshal(body, &books); err != nil {
			fmt.Println(err, "in decoding JSON")
			return
		}
		for i, v := range books {
			fmt.Println("-------------------------------")
			fmt.Println("Key:", i)
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
func storeBook() int {
	// Make body
	book := Book{
		Name:       BookName[rand.Intn(len(BookName))],
		Author:     BookAuthor[rand.Intn(len(BookAuthor))],
		Pages:      rand.Intn(BookMaxPages),
		Year:       rand.Intn(time.Now().Year()),
		CreateTime: time.Now(),
	}
	b, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err, "in encoding a book as JSON")
		return 1
	}

	// Send request
	resp, err := http.Post(BookURL+"books", "application/json", bytes.NewReader(b))
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
func deleteBook(key string) int {
	// Send request
	pReq, err := http.NewRequest("DELETE", BookURL+"books/"+key, nil)
	if err != nil {
		fmt.Println(err, "in making request")
		return 1
	}
	resp, err := http.DefaultClient.Do(pReq)
	if err != nil {
		fmt.Println(err, "in sending request")
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

// Return 0: success
// Return 1: failed
func deleteAll() int {
	// Send request
	pReq, err := http.NewRequest("DELETE", BookURL+"books", nil)
	if err != nil {
		fmt.Println(err, "in making request")
		return 1
	}
	resp, err := http.DefaultClient.Do(pReq)
	if err != nil {
		fmt.Println(err, "in sending request")
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
	// Random seed
	rand.Seed(time.Now().Unix())

	// Test suite
	if storeBook() != 0 {
		fmt.Println("Store failed")
		return
	} else {
		fmt.Println("Store a book")
	}
	queryAll()
	if deleteAll() != 0 {
		fmt.Println("Delete failed")
		return
	} else {
		fmt.Println("Delete all")
	}
}
