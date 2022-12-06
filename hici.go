package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Data struct {
	Repo string
	User string
}

func ReadConfig() Data {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	// Now let's unmarshall the data into `payload`
	var payload Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}

func main() {

	payload := ReadConfig()

	if len(os.Args[0:]) >= 3 {
		payload.User = os.Args[1]
		payload.Repo = os.Args[2]
	}

	fmt.Printf("!...Hello Continuous Integration, it's %s from %s...!", payload.User, payload.Repo)
	return
}
