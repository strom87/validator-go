package validator

import (
	"strconv"
	"strings"
)

func strToInt(value string) (int64, error) {
	number, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func strToFloat(value string) (float64, error) {
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0, err
	}
	return number, nil
}

func strToUint(value string) (uint64, error) {
	number, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func splitAtt(value string) (string, string) {
	val := strings.Split(value, ",")
	return val[0], val[1]
}

func addErrMsg(property string, message string) {
	if errMsg == nil {
		errMsg = map[string][]string{}
	}

	if errMsg[property] == nil {
		errMsg[property] = []string{message}
	} else {
		errMsg[property] = append(errMsg[property], message)
	}
}
