package validator

import (
	"errors"
	"reflect"
	reg "regexp"
)

func validateProperty(p *Property) error {
	for _, v := range p.Validations {
		switch v.Type {
		case "min":
			if err := min(p, v.Value); err != nil {
				return err
			}
		case "max":
			if err := max(p, v.Value); err != nil {
				return err
			}
		case "between":
			if err := between(p, v.Value); err != nil {
				return err
			}
		case "equals":
			if err := equals(p, v.Value); err != nil {
				return err
			}
		case "regexp":
			if err := regexp(p, v.Value); err != nil {
				return err
			}
		case "required":
			required(p)
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
			addErrMsg(p.Name, strReplace(lang.MinString, "{min}", value))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num, err := strToInt(value)
		if err != nil {
			return err
		}
		if int64(p.Value.Int()) < num {
			addErrMsg(p.Name, strReplace(lang.MinNumber, "{min}", value))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		num, err := strToUint(value)
		if err != nil {
			return err
		}
		if p.Value.Uint() < num {
			addErrMsg(p.Name, strReplace(lang.MinNumber, "{min}", value))
		}
	case reflect.Float32, reflect.Float64:
		num, err := strToFloat(value)
		if err != nil {
			return err
		}
		if float64(p.Value.Float()) < num {
			addErrMsg(p.Name, strReplace(lang.MinNumber, "{min}", value))
		}
	}

	return nil
}

func max(p *Property, value string) error {
	switch p.Kind {
	case reflect.String:
		num, err := strToInt(value)
		if err != nil {
			return err
		}
		if int64(p.Value.Len()) > num {
			addErrMsg(p.Name, strReplace(lang.MaxString, "{max}", value))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num, err := strToInt(value)
		if err != nil {
			return err
		}
		if int64(p.Value.Int()) > num {
			addErrMsg(p.Name, strReplace(lang.MaxNumber, "{max}", value))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		num, err := strToUint(value)
		if err != nil {
			return err
		}
		if p.Value.Uint() > num {
			addErrMsg(p.Name, strReplace(lang.MaxNumber, "{max}", value))
		}
	case reflect.Float32, reflect.Float64:
		num, err := strToFloat(value)
		if err != nil {
			return err
		}
		if float64(p.Value.Float()) > num {
			addErrMsg(p.Name, strReplace(lang.MaxNumber, "{max}", value))
		}
	}

	return nil
}

func between(p *Property, value string) error {
	minStr, maxStr := splitAtt(value)

	strError := strReplace(lang.BetweenString, "{min}", minStr)
	strError = strReplace(strError, "{max}", maxStr)
	numError := strReplace(lang.BetweenNumber, "{min}", minStr)
	numError = strReplace(numError, "{max}", maxStr)

	switch p.Kind {
	case reflect.String:
		min, err := strToInt(minStr)
		if err != nil {
			return err
		}
		max, err := strToInt(maxStr)
		if err != nil {
			return err
		}
		len := int64(p.Value.Len())
		if len <= min || len >= max {
			addErrMsg(p.Name, strError)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		min, err := strToInt(minStr)
		if err != nil {
			return err
		}
		max, err := strToInt(maxStr)
		if err != nil {
			return err
		}
		len := int64(p.Value.Int())
		if len <= min || len >= max {
			addErrMsg(p.Name, numError)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		min, err := strToUint(minStr)
		if err != nil {
			return err
		}
		max, err := strToUint(maxStr)
		if err != nil {
			return err
		}

		if p.Value.Uint() <= min || p.Value.Uint() >= max {
			addErrMsg(p.Name, numError)
		}
	case reflect.Float32, reflect.Float64:
		min, err := strToFloat(minStr)
		if err != nil {
			return err
		}
		max, err := strToFloat(maxStr)
		if err != nil {
			return err
		}
		len := float64(p.Value.Float())
		if len <= min || len >= max {
			addErrMsg(p.Name, numError)
		}
	}

	return nil
}

func equals(p *Property, value string) error {
	property, err := getProperty(value)
	if err != nil {
		return err
	}

	strError := lang.Equals
	if useJson {
		strError = strReplace(strError, "{prop1}", getPropertyJsonName(p))
		strError = strReplace(strError, "{prop2}", getPropertyJsonName(property))
	} else {
		strError = strReplace(strError, "{prop1}", p.Name)
		strError = strReplace(strError, "{prop2}", property.Name)
	}

	switch p.Kind {
	case reflect.String:
		if p.Value.String() != property.Value.String() {
			addErrMsg(p.Name, strError)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if int64(p.Value.Int()) != int64(property.Value.Int()) {
			addErrMsg(p.Name, strError)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if p.Value.Uint() != property.Value.Uint() {
			addErrMsg(p.Name, strError)
		}
	case reflect.Float32, reflect.Float64:
		if float64(p.Value.Float()) != float64(property.Value.Float()) {
			addErrMsg(p.Name, strError)
		}
	}

	return nil
}

func regexp(p *Property, value string) error {
	if p.Kind != reflect.String {
		return errors.New("Regexp validation only handles string, invalid type used")
	}

	r, err := reg.Compile(value)
	if err != nil {
		return err
	}

	if !r.MatchString(p.Value.String()) {
		addErrMsg(p.Name, lang.Regexp)
	}

	return nil
}

func required(p *Property) {
	switch p.Kind {
	case reflect.String:
		if int64(p.Value.Len()) <= 0 {
			addErrMsg(p.Name, lang.Required)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if int64(p.Value.Int()) <= 0 {
			addErrMsg(p.Name, lang.Required)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if p.Value.Uint() <= 0 {
			addErrMsg(p.Name, lang.Required)
		}
	case reflect.Float32, reflect.Float64:
		if float64(p.Value.Float()) <= 0.0 {
			addErrMsg(p.Name, lang.Required)
		}
	}
}
