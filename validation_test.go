package validator

import (
	"reflect"
	"testing"
)

type function func(p *Property, value string) error

type minTestStruct struct {
	v1       int     `validator:"min:200"`
	v2       string  `validator:"min:4"`
	v3       float64 `validator:"min:10.5"`
	expected int
}

type maxTestStruct struct {
	v1       int     `validator:"max:10"`
	v2       string  `validator:"max:4"`
	v3       float64 `validator:"max:10.5"`
	expected int
}

type betweenTestStruct struct {
	v1       int     `validator:"between:15,20"`
	v2       int     `validator:"between:15,20"`
	v3       float64 `validator:"between:100,200"`
	v4       float64 `validator:"between:100,200"`
	v5       string  `validator:"between:3,6"`
	v6       string  `validator:"between:3,6"`
	expected int
}

func TestMinF(t *testing.T) {
	initialize("en")
	values := []minTestStruct{
		{199, "Bob", 10.49, 3},
		{300, "Jhonny", 1042.95, 0},
		{-2, "Brit", 10.5, 1},
	}

	for _, v := range values {
		testRunner(v, min)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		errMsg = nil
	}
}

func TestMaxF(t *testing.T) {
	initialize("en")
	values := []maxTestStruct{
		{11, "Jonny", 10.511, 3},
		{10, "Bobb", 10.5, 0},
		{-200, "abcdefg", 10.499, 1},
	}

	for _, v := range values {
		testRunner(v, max)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		errMsg = nil
	}
}

func TestBetweenF(t *testing.T) {
	initialize("en")
	values := []betweenTestStruct{
		{17, 2000, 175.2, 99.99, "Bob", "Jonny", 3},
		{16, 19, 199.99, 100.1, "Mats", "Jonny", 0},
		{15, 20, 200.0, 100.0, "Max", "Johnnys", 6},
	}

	for _, v := range values {
		testRunner(v, between)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		errMsg = nil
	}
}

func testRunner(obj interface{}, fn function) {
	data := reflect.ValueOf(obj)
	for i := 0; i < data.NumField(); i++ {
		p := NewProperty(data.Field(i), data.Type().Field(i))

		for _, v := range p.Validations {
			fn(p, v.Value)
		}
	}
}
