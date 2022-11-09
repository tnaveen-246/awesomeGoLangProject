package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Go Lang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:33 Monday"))

	createDate := time.Date(2020, time.November, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
