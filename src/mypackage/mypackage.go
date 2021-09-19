package mypackage

import "fmt"

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage Print a message
func PrintMessage(text string) {

	fmt.Println(text)
}
