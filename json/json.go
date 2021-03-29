package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/pborman/uuid"
)

// Book struct for marshaling
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Author struct to be used in book struct
type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

// SensorReading type to read from JSON, unmarshalling
type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {

	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	// Marshalling, from struct to JSON string
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	// Indented marshalling
	byteArrayIndented, err := json.MarshalIndent(book, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArrayIndented))

	// Unmarshalling a sensor data to a SensorReading struct
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading SensorReading
	json.Unmarshal([]byte(jsonString), &reading)

	fmt.Printf("\n%+v\n", reading)
	fmt.Println("Sensor: " + reading.Name)

	/*
	  Unstructured Data: Sometimes, you might not have knowledge of the structure of the JSON string that you are reading
	  use map[string]interface{} as the type you unmarshal into
	*/

	var readingUnstructured map[string]interface{}
	json.Unmarshal([]byte(jsonString), &readingUnstructured)
	fmt.Printf("\n%+v\n", readingUnstructured)
	fmt.Println(readingUnstructured["capacity"])
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", ".", -1)
	fmt.Println(uuid)
}
