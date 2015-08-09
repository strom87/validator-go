package validator

import (
	"log"
	"reflect"
)

var (
	errMsg map[string][]string
)

func Validate(obj interface{}) (map[string][]string, error) {
	err := readObject(reflect.ValueOf(obj))

	log.Println(errMsg)

	return errMsg, err
}

func readObject(obj reflect.Value) error {
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
			err := validateProperty(NewProperty(field, obj.Type().Field(i)))
			if err != nil {
				return err
			}
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
		}
	}

	return nil
}
