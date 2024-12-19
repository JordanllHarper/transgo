package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

type ErrorResponse struct {
	code    uint
	message string
}

// Crud operations on trains
func handleTrains(w http.ResponseWriter, r *http.Request, db *gorm.DB) error {
	method := r.Method
	encoder := json.NewEncoder(w)
	var response any
	response = struct {
		code    uint
		message string
	}{
		code:    500,
		message: "Internal Server Error ",
	}

	var err error

	switch method {
	case "GET":
		response, err = onTrainGet(db, r)

	case "POST":
		response, err = onTrainPost(db, r)

	default:
		response = struct {
			code    uint
			message string
		}{
			code:    405,
			message: "Method Not Allowed",
		}
	}

	encoder.Encode(response)

	return err
}

func main() {
	// db
	dsn := "root:@tcp(127.0.0.1:3306)/trainsgo?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)

	checkError(err)

	{
		http.HandleFunc("/", handleRoot)
		http.HandleFunc("/trains", func(w http.ResponseWriter, r *http.Request) {
			err := handleTrains(
				w,
				r,
				db,
			)
			if err != nil {
				checkError(err)
			}
		})
	}

	{
		address := "127.0.0.1"
		port := "3333"
		listening := fmt.Sprintf("%v:%v", address, port)

		fmt.Printf("Starting server. Listening on port: %v\n", listening)
		err = http.ListenAndServe(listening, nil)
	}

	fmt.Printf("%v\n", err)

}

// e = ln
