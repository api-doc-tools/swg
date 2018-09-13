package main

import "fmt"

func main() {
	err := SaveSwaggerDoc()
	if err != nil {
		fmt.Println(err)
		return
	}
}
