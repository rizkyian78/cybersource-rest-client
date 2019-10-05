package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/tooolbox/jwalk"
	// "github.com/tooolbox/cybersource-rest-client-go/client"
)

// func main() {

// 	c := client.Default
// 	created, err := c.Payments.CreatePayment()

// 	log.Printf("Created: %#v", created)
// 	log.Printf("Err: %v", err)
// }

func main() {
	log.Print("Walking...")

	f, err := os.Open("./generator/cybersource-rest-spec.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	i, err := jwalk.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	var walk func(interface{}) error
	var walkObj func(interface{}) (interface{}, error)

	walk = func(i interface{}) error {
		switch v := i.(type) {
		case jwalk.ObjectWalker:
			return v.Walk(func(key string, value interface{}) (interface{}, error) {
				return walkObj(value)
			})
		case jwalk.ObjectsWalker:
			return v.Walk(func(obj jwalk.ObjectWalker) error {
				return obj.Walk(func(key string, value interface{}) (interface{}, error) {
					return walkObj(value)
				})
			})
		case jwalk.Value:
			return nil
		default:
			return nil
		}
	}

	walkObj = func(i interface{}) (interface{}, error) {
		switch v := i.(type) {
		case jwalk.ObjectWalker:
			return v, v.Walk(func(key string, value interface{}) (interface{}, error) {
				return walkObj(value)
			})
		case jwalk.ObjectsWalker:
			return v, v.Walk(func(obj jwalk.ObjectWalker) error {
				return obj.Walk(func(key string, value interface{}) (interface{}, error) {
					return walkObj(value)
				})
			})
		case jwalk.Value:
			return v, nil
		default:
			return v, nil
		}
	}

	if err := walk(i); err != nil {
		log.Fatal(err)
	}

	log.Print("Done.")
}
