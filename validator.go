package validator

import (
	"reflect"
)

var (
	errMsg     map[string][]string
	properties []*Property
	lang       *Language
	useJson    bool
)

func initialize(language string) {
	errMsg = nil
	useJson = false
	properties = nil
	lang = NewLanguage(language)
}

func Validate(obj interface{}) (bool, map[string][]string, error) {
	return ValidateLang(obj, "en")
}

func ValidateJson(obj interface{}) (bool, map[string][]string, error) {
	return ValidateJsonLang(obj, "en")
}

func ValidateLang(obj interface{}, language string) (bool, map[string][]string, error) {
	initialize(language)
	readObject(reflect.ValueOf(obj))
	if err := loopProperties(); err != nil {
		return false, nil, err
	}

	return isValid(), errMsg, nil
}

func ValidateJsonLang(obj interface{}, language string) (bool, map[string][]string, error) {
	initialize(language)
	useJson = true

	readObject(reflect.ValueOf(obj))
	if err := loopProperties(); err != nil {
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
		default:
			addProperties(NewProperty(field, obj.Type().Field(i)))
		}
	}
}

func loopProperties() error {
	for _, p := range properties {
		if err := validateProperty(p); err != nil {
			return err
		}
	}

	return nil
}
