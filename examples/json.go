package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Message struct {
	Name string
	Body string
	Time int64
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}
type Bar struct {
}
type Foo struct {
	Bar *Bar
}
type Command struct {
}
type IncomingMessage struct {
	Cmd *Command
	Msg *Message
}

func main() {

	// EXAMPLE 1
	m := Message{"Alice", "Hello", 1294706395881547000}

	b, err := json.Marshal(m)
	fmt.Printf("%+q %T", b, err)
	fmt.Println()

	var msg Message

	err1 := json.Unmarshal(b, &msg)
	fmt.Println(msg, err1)

	// EXAMPLE 2
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777

	r := i.(float64)
	fmt.Println("the circle's area", math.Pi*r*r)

	switch v := i.(type) {
	case int:
		fmt.Println("twice i is", v*2)
	case float64:
		fmt.Println("the reciprocal of i is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is", v[h:]+v[:h])
	default:
		fmt.Println("DEFAULT")
	}

	// EXAMPLE 3
	b1 := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var f interface{}
	err2 := json.Unmarshal(b1, &f)
	fmt.Println(f, err2)

	fmap := f.(map[string]interface{})

	for k, v := range fmap {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	//EXAMPLE 4

	var fm FamilyMember
	var foomem Foo
	var im IncomingMessage
	err3 := json.Unmarshal(b1, &fm)
	err4 := json.Unmarshal(b1, &foomem)
	err5 := json.Unmarshal(b1, &im)

	fmt.Println(fm, err3)
	fmt.Println(foomem, err4)
	fmt.Println(im, err5)
}
