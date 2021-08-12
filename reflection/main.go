package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

// func (v Value) Interface() interface{}

func main() {
	fmt.Println("FIRST LAW")
	var y float64 = 4
	fmt.Println("type:", reflect.TypeOf(y))
	fmt.Println("value:", reflect.ValueOf(y).String())

	var z float64 = 3.4
	v := reflect.ValueOf(z)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var x uint8 = 'x'
	value := reflect.ValueOf(x)
	fmt.Println("type:", value.Type())
	fmt.Println("kind is uint8: ", value.Kind() == reflect.Uint8)
	// x = uint8(v.Uint())

	fmt.Println("SECOND LAW")

	type MyInt int
	var a MyInt = 7
	aval := reflect.ValueOf(a)
	fmt.Println(aval, aval.Kind(), aval.Type(), reflect.TypeOf(a))

	b := v.Interface().(float64)
	fmt.Println(b)
	fmt.Println(v.Interface())
	fmt.Printf("value is %7.1e\n", v.Interface())

	fmt.Println("THIRD LAW")
	var c float64 = 3.4
	v2 := reflect.ValueOf(c)
	// v2.SetFloat(7.1) // Error: will panic.
	fmt.Println("settability of v2:", v2.CanSet())

	p := reflect.ValueOf(&c)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v3 := p.Elem()
	fmt.Println("settability of v2:", v3.CanSet())

	v3.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(c)

	fmt.Println("STRUCT EXAMPLE")

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
