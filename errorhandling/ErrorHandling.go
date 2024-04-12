package errorhandling

import "fmt"

func Manage(e error) {
	if e != nil {
		fmt.Println("Something went wrong")
		panic(e)
	}
}
