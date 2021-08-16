package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	imgPath := "./image.jpg"

	dat, err := ioutil.ReadFile(imgPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	blob, err := json.Marshal(dat)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s\n", blob)
}