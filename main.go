package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

type Fixture struct {
	Foo string
	Bar string
}

func main() {
	buffer, err := ioutil.ReadFile("fixture.yml")
	fixtures := map[string]Fixture{}
	if err == nil {
		yaml.Unmarshal(buffer, &fixtures)

		for key, fixture := range fixtures {
			log.Printf("Fixture %s: %+v", key, fixture)
		}
	}
}
