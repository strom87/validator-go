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
	val1, val2 := splitAtt(value)

	switch p.Kind {
	case reflect.String:
		min, err := strToInt(val1)
		if err != nil {
			return err
		}
		max, err := strToInt(val2)
		if err != nil {
			return err
		}
		len := int64(p.Value.Len())
		if len <= min || len >= max {
			addErrMsg(p.Name, "between string")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		min, err := strToInt(val1)
		if err != nil {
			return err
		}
		max, err := strToInt(val2)
		if err != nil {
			return err
		}
		len := int64(p.Value.Int())
		if len <= min || len >= max {
			addErrMsg(p.Name, "between int")
		}
	}

	return nil
}
