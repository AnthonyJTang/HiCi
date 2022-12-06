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

	// fmt.Println(os.Args)
	githubActor := os.Args[1]

	payload := ReadConfig()
	configedActor := payload.User + "/" + payload.Repo
	if githubActor == configedActor {
		fmt.Printf("!...HiCI it's %s from repo %s...!", payload.User, payload.Repo)
	}

}
