package validator

import (
	"reflect"
	"strings"
)

type Property struct {
	Name        string
	NameJson    string
	Value       reflect.Value
	Kind        reflect.Kind
	Validations []ValidationType
}

func NewProperty(value reflect.Value, field reflect.StructField) *Property {
	p := Property{Name: field.Name, Value: value, Kind: value.Kind()}
	p.setJsonName(field)
	p.addValidations(field)
	return &p
}

func (p *Property) addValidations(field reflect.StructField) {
	p.Validations = []ValidationType{}
	if tag := field.Tag.Get("validator"); tag != "" {
		validations := strings.Split(string(tag), "|")
		for _, v := range validations {
			a, b := strSplit(v, ":")
			p.Validations = append(p.Validations, ValidationType{a, b})
		}
	}
}

func (p *Property) setJsonName(field reflect.StructField) {
	if tag := field.Tag.Get("json"); tag != "" {
		p.NameJson = tag
	}
}
