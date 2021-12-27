package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/tooolbox/jwalk"
)

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

	checkStringValue := func(i interface{}, expect string) bool {
		if val, ok := i.(jwalk.Value); ok {
			return val.String() == expect
		}
		return false
	}

	fixInvalid := func(name string, obj jwalk.ObjectWalker) (string, interface{}, error) {

		hasStringType, hasMaximum := false, false
		hasBoolType, hasMaxLength := false, false

		err := obj.Walk(func(key string, value interface{}) (string, interface{}, error) {

			if key == "type" && checkStringValue(value, "string") {
				hasStringType = true
			}
			if key == "maximum" {
				hasMaximum = true
			}

			if key == "type" && checkStringValue(value, "boolean") {
				hasBoolType = true
			}
			if key == "maxLength" {
				hasMaxLength = true
			}

			return key, value, nil
		})
		if err != nil {
			return name, obj, err
		}

		if hasStringType && hasMaximum {
			log.Print("Found a string type with a maximum field")
			err = obj.Walk(func(key string, value interface{}) (string, interface{}, error) {
				if key == "maximum" {
					return "maxLength", value, nil
				}
				return key, value, nil
			})
		}

		if hasBoolType && hasMaxLength {
			log.Print("Found a boolean type with a maxLength field")
			err = obj.Walk(func(key string, value interface{}) (string, interface{}, error) {
				if key == "maxLength" {
					return "", nil, jwalk.RemoveField
				}
				return key, value, nil
			})
		}

		return name, obj, err
	}

	var walk func(string, interface{}) (string, interface{}, error)

	walk = func(k string, i interface{}) (string, interface{}, error) {
		switch v := i.(type) {
		case jwalk.ObjectWalker:
			kRepl, vRepl, err := fixInvalid(k, v)
			if err != nil {
				return k, i, err
			}
			return kRepl, vRepl, vRepl.(jwalk.ObjectWalker).Walk(func(key string, value interface{}) (string, interface{}, error) {
				return walk(key, value)
			})
			// return k, v, v.Walk(func(key string, value interface{}) (string, interface{}, error) {
			// 	return walk(key, value)
			// })
		case jwalk.ObjectsWalker:
			return k, v, v.Walk(func(obj jwalk.ObjectWalker) error {

				return obj.Walk(func(key string, value interface{}) (string, interface{}, error) {
					return walk(key, value)
				})
			})
		case jwalk.Value:
			return k, v, nil
		default:
			return k, v, nil
		}
	}

	_, result, err := walk("", i)
	if err != nil {
		log.Fatal(err)
	}

	n, err := os.Create("./generator/cybersource-rest-spec-fixed.json")
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(n)
	enc.SetIndent("", "    ")
	if err := enc.Encode(result); err != nil {
		log.Fatal(err)
	}

	log.Print("Done.")
}
