package validator

import (
	"reflect"
)

var (
	errMsg     map[string][]string
	properties []*Property
)

func Validate(obj interface{}) (bool, map[string][]string, error) {
	errMsg = nil
	readObject(reflect.ValueOf(obj))

	err := loopProperties()
	if err != nil {
		return false, nil, err
	}

	return isValid(), errMsg, nil
}

func ValidateJson(obj interface{}) (bool, map[string][]string, error) {
	_, _, err := Validate(obj)
	if err != nil {
		return false, nil, err
	}

	propertyNamesToJson()

	return isValid(), errMsg, nil
}

func readObject(obj reflect.Value) {
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			readObject(field)
		case reflect.Slice:
			for n := 0; n < field.Len(); n++ {
				readObject(field.Index(n))
			}
		default:
			addProperties(NewProperty(field, obj.Type().Field(i)))
		}
	}
}

func loopProperties() error {
	for _, p := range properties {
		err := validateProperty(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateProperty(p *Property) error {
	for _, v := range p.Validations {
		switch v.Type {
		case "min":
			err := min(p, v.Value)
			if err != nil {
				return err
			}
		case "max":
			err := max(p, v.Value)
			if err != nil {
				return err
			}
		case "between":
			err := between(p, v.Value)
			if err != nil {
				return err
			}
		case "equals":
			err := equals(p, v.Value)
			if err != nil {
				return err
			}
		case "regexp":
			err := regexp(p, v.Value)
			if err != nil {
				return err
			}
		case "required":
			required(p)
		}
	}

	return nil
}
