package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("testdata/testfile.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
