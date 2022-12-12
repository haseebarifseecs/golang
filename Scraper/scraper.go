package main

import (
	"fmt"
	"net/http"
)

type data struct {
	response   string
	statusCode string
	title      string
}

func (d *data) getRequest() {
	resp, _ := http.Get("https://www.goodreads.com/quotes")

	fmt.Println(resp)
}
func main() {
	var d data = data{}
	d.getRequest()
}
