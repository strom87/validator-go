package validator

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Int1    int     `validator:"min:3|max:2"`
	Int2    int     `validator:"min:3|max:2"`
	Float1  float64 `validator:"min:5.5|max:4.2"`
	Float2  float64 `validator:"min:5.5|max:4.2"`
	String1 string  `validator:"min:4|max:3"`
	String2 string  `validator:"min:4|max:3"`
}

func TestMin(t *testing.T) {
	errMsg = nil
	ts := TestStruct{2, 3, 5.4, 5.6, "123", "1234"}
	obj := reflect.ValueOf(ts)

	for i := 0; i < obj.NumField(); i++ {
		p := NewProperty(obj.Field(i), obj.Type().Field(i))

		for _, v := range p.Validations {
			min(p, v.Value)
		}
	}

	if errMsg["Int1"] == nil || errMsg["Float1"] == nil || errMsg["String1"] == nil {
		t.Error("Int1 is not greater than 3 or Float1 is greater then 5.4 or String1 not greater than 4")
	}
}

func TestMax(t *testing.T) {
	errMsg = nil
	ts := TestStruct{2, 3, 5.4, 5.6, "123", "1234"}
	obj := reflect.ValueOf(ts)

	for i := 0; i < obj.NumField(); i++ {
		p := NewProperty(obj.Field(i), obj.Type().Field(i))

		for _, v := range p.Validations {
			max(p, v.Value)
		}
	}

	if errMsg["Int2"] == nil || errMsg["Float2"] == nil || errMsg["String2"] == nil {
		t.Error("Int2 is not less then 2 or Float2 is not less then 5.5 or String2 is not less then 3")
	}
}
