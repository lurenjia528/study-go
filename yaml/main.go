package main

//
//import (
//	"fmt"
//	"github.com/projectcalico/go-yaml-wrapper"
//)
//
//func main() {
//	json := `{
//  "name": "yaml",
//  "version": "1.0.0",
//  "dependencies": {
//    "host": "ygt.com",
//    "ip": [
//      "1.2.3.4",
//      "2.3.4.5"
//    ]
//  }
//}`
//	toYAML, err := yaml.JSONToYAML([]byte(json))
//	if err!=nil{
//		panic(err)
//	}
//	fmt.Println(string(toYAML))
//	yml := `
//dependencies:
//  host: ygt.com
//  ip:
//  - 1.2.3.4
//  - 2.3.4.5
//name: yaml
//version: 1.0.0
//`
//	toJSON, err := yaml.YAMLToJSON([]byte(yml))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(toJSON))
//}


//
//import (
//	"fmt"
//	"log"
//
//	"gopkg.in/yaml.v2"
//)
//var data = `
//a: Easy!
//b:
//  c: 2
//  d: [3, 4]
//`
//type T struct {
//	A string
//	B struct {
//		RenamedC int   `yaml:"c"`
//		D        []int `yaml:",flow"`
//	}
//}
//func main() {
//	t := T{}
//
//	err := yaml.Unmarshal([]byte(data), &t)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- t:\n%v\n\n", t)
//
//	d, err := yaml.Marshal(&t)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- t dump:\n%s\n\n", string(d))
//
//	m := make(map[interface{}]interface{})
//
//	err = yaml.Unmarshal([]byte(data), &m)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- m:\n%v\n\n", m)
//
//	d, err = yaml.Marshal(&m)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- m dump:\n%s\n\n", string(d))
//}

import (
	"fmt"

	"github.com/ghodss/yaml"
)
type Person struct {
	Name string `json:"name"` // Affects YAML field names too.
	Age  int    `json:"age"`
}
func main() {
	// Marshal a Person struct to YAML.
	p := Person{"John", 30}
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
	/* Output:
	age: 30
	name: John
	*/

	// Unmarshal the YAML back into a Person struct.
	var p2 Person
	err = yaml.Unmarshal(y, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(p2)
	/* Output:
	{John 30}
	*/

	fmt.Println("------------")
	j := []byte(`{"name": "John", "age": 30}`)
	y, err = yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
	/* Output:
	name: John
	age: 30
	*/
	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
	/* Output:
	{"age":30,"name":"John"}
	*/
}