package validator

import (
	"reflect"
)

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

func max(p *Property, value string) error {
	switch p.Kind {
	case reflect.String:
		num, err := strToInt(value)
		if err != nil {
			return err
		}
		if int64(p.Value.Len()) > num {
			addErrMsg(p.Name, "max string wrong")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num, err := strToInt(value)
		if err != nil {
			return err
		}
		if int64(p.Value.Int()) > num {
			addErrMsg(p.Name, "max int wrong")
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		num, err := strToUint(value)
		if err != nil {
			return err
		}
		if p.Value.Uint() > num {
			addErrMsg(p.Name, "max uint wrong")
		}
	case reflect.Float32, reflect.Float64:
		num, err := strToFloat(value)
		if err != nil {
			return err
		}
		if float64(p.Value.Float()) > num {
			addErrMsg(p.Name, "max float wrong")
		}
	}

	return nil
}

func between(p *Property, value string) error {
	minStr, maxStr := splitAtt(value)

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
			addErrMsg(p.Name, "between string")
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
			addErrMsg(p.Name, "between int")
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
			addErrMsg(p.Name, "between uint")
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
			addErrMsg(p.Name, "between float")
		}
	}

	return nil
}

func equals(p *Property, value string) error {
	property, err := getProperty(value)
	if err != nil {
		return err
	}

	switch p.Kind {
	case reflect.String:
		if p.Value.String() != property.Value.String() {
			addErrMsg(p.Name, "same string error")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if int64(p.Value.Int()) != int64(property.Value.Int()) {
			addErrMsg(p.Name, "same int error")
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if p.Value.Uint() != property.Value.Uint() {
			addErrMsg(p.Name, "same uint error")
		}
	case reflect.Float32, reflect.Float64:
		if float64(p.Value.Float()) != float64(property.Value.Float()) {
			addErrMsg(p.Name, "same float error")
		}
	}

	return nil
}

func required(p *Property) {
	switch p.Kind {
	case reflect.String:
		if int64(p.Value.Len()) <= 0 {
			addErrMsg(p.Name, "required string wrong")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if int64(p.Value.Int()) <= 0 {
			addErrMsg(p.Name, "required int wrong")
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if p.Value.Uint() <= 0 {
			addErrMsg(p.Name, "required uint wrong")
		}
	case reflect.Float32, reflect.Float64:
		if float64(p.Value.Float()) <= 0.0 {
			addErrMsg(p.Name, "required float wrong")
		}
	}
}
