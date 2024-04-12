package fileoperations

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chethanbhat/go-shopping-cart/errorhandling"
)

// Read from JSON File
func ReadJSONData[T any](filename string, data *T) error {
	jsonobj, err := os.ReadFile(filename)
	errorhandling.Manage(err)

	err = json.Unmarshal(jsonobj, data)
	return err
}

// Write to JSON File
func WriteJSONData[T any](filename string, data *T) error {
	jsonData, err := json.Marshal(data)
	errorhandling.Manage(err)

	file, err := os.Create(filename)
	errorhandling.Manage(err)

	defer file.Close()

	_, err = file.Write(jsonData)
	errorhandling.Manage(err)
	return err
}

// Read User Input
func ReadUserInput(prompt string, value *string) {
	fmt.Println(prompt)
	fmt.Scan(value)
}
