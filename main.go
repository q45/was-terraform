package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CustomTags struct {
	Tags []string `json:"custom_tags"`
}

func main() {

	tags := CustomTags{
		Tags: []string{"ai", "dev", "cloudorchestration"},
	}

	json, err := json.Marshal(tags)
	if err != nil {
		log.Fatal("%v", err)
	}

	fmt.Println(string(json))

	err = os.WriteFile("data.tfvars", []byte(json), 0666)
	if err != nil {
		log.Fatal("%v", err)
	}

}
