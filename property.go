package validator

import (
	"reflect"
	"strings"
)

type Property struct {
	Name        string
	Value       reflect.Value
	Kind        reflect.Kind
	Validations []ValidationType
}

func NewProperty(value reflect.Value, field reflect.StructField) *Property {
	p := Property{Name: field.Name, Value: value, Kind: value.Kind()}
	p.addValidations(field)
	return &p
}

func (p *Property) addValidations(field reflect.StructField) {
	p.Validations = []ValidationType{}
	tag := field.Tag.Get("validator")
	validations := strings.Split(string(tag), "|")

	for _, v := range validations {
		val := strings.Split(v, ":")
		p.Validations = append(p.Validations, ValidationType{val[0], val[1]})
	}
}
