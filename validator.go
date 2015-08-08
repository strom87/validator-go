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
		}
	}

	return nil
}

func min(p *Property, value string) error {
	switch p.Kind {
	case reflect.String:
		num, err := strToInt(value)
		if err != nil {
			return err
		}
		if int64(p.Value.Len()) < num {
			addErrMsg(p.Name, "min string wrong")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num, err := strToInt(value)
		if err != nil {
			return err
		}
		if int64(p.Value.Int()) < num {
			addErrMsg(p.Name, "min int wrong")
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		num, err := strToUint(value)
		if err != nil {
			return err
		}
		if p.Value.Uint() < num {
			addErrMsg(p.Name, "min uint wrong")
		}
	case reflect.Float32, reflect.Float64:
		num, err := strToFloat(value)
		if err != nil {
			return err
		}
		if float64(p.Value.Float()) < num {
			addErrMsg(p.Name, "min float wrong")
		}
	}

	return nil
}
