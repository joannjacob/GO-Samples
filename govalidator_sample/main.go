package main

import (
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
)

type User struct {
	Name string      `valid:"type(string)"`
	Age  int         `valid:"type(int)"`
	Meta interface{} `valid:"type(string)"`
}

type Person struct {
	FirstName string
	LastName  string
}
type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"duck,ascii"`
	Message2 string `valid:"animal(dog)"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

func main() {
	// govalidator.SetFieldsRequiredByDefault(true)

	// print(govalidator.ValidateStruct(User{"j", "joann"}))

	println(govalidator.IsURL(`http://user@pass:domain.com/path/page`))
	println(govalidator.IsType("Bob", "string"))
	println(govalidator.IsType(1, "int"))
	i := 1
	println(govalidator.IsType(&i, "*int"))

	result, err := govalidator.ValidateStruct(User{"Bob", 20, "meta"})
	if err != nil {
		println("error: " + err.Error())
	}
	println("Validate struct", result)

	str := govalidator.ToString(&Person{"John", "Juan"})
	println("To string", str)

	data := []interface{}{1, 2, 3, 4, 5}

	fmt.Println("ITERATOR")
	var fn1 govalidator.Iterator = func(value interface{}, index int) {
		println(value.(int))
	}
	govalidator.Each(data, fn1)

	fmt.Println(" RESULT ITERATOR")
	var fn2 govalidator.ResultIterator = func(value interface{}, index int) interface{} {
		return value.(int) * 3
	}
	_ = govalidator.Map(data, fn2)

	fmt.Println(" CONDITION ITERATOR")
	data2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var fn3 govalidator.ConditionIterator = func(value interface{}, index int) bool {
		return value.(int)%2 == 0
	}
	_ = govalidator.Filter(data2, fn3) // result = []interface{}{2, 4, 6, 8, 10}
	_ = govalidator.Count(data2, fn3)  // result = 5

	fmt.Println("STRUCT VALIDATION TAGS")
	post := &Post{
		Title:    "My Example Post",
		Message:  "duck",
		Message2: "dog",
		AuthorIP: "123.234.54.3",
	}

	// Add your own struct validation tags
	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})

	// Add your own struct validation tags with parameter
	govalidator.ParamTagMap["animal"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		species := params[0]
		return str == species
	})
	govalidator.ParamTagRegexMap["animal"] = regexp.MustCompile("^animal\\((\\w+)\\)$")

	result, err1 := govalidator.ValidateStruct(post)
	if err1 != nil {
		println("error: " + err1.Error())
	}
	println(result)

	fmt.Println("VALIDATE MAPS")

	var mapTemplate = map[string]interface{}{
		"name":       "required,alpha",
		"family":     "required,alpha",
		"email":      "required,email",
		"cell-phone": "numeric",
		"address": map[string]interface{}{
			"line1":       "required,alphanum",
			"line2":       "alphanum",
			"postal-code": "numeric",
		},
	}

	var inputMap = map[string]interface{}{
		"name":   "Bob",
		"family": "Smith",
		"email":  "foo@bar.baz",
		"address": map[string]interface{}{
			"line1":       "",
			"line2":       "",
			"postal-code": "",
		},
	}

	result, err2 := govalidator.ValidateMap(mapTemplate, inputMap)
	if err2 != nil {
		println("error: " + err2.Error())
	}
	println(result)

	// Remove all characters from string ignoring characters between "a" and "z"
	println("WHITELIST")
	println(govalidator.WhiteList("a3a43a5a4a3a2a23a4a5a4a3a4", "a-z") == "aaaaaaaaaaaa")

}
