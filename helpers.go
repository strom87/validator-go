package validator

import (
	"errors"
	"strconv"
	"strings"
)

func strSplit(text string, sep string) (string, string) {
	length := len(text)
	index := strings.Index(text, sep)

	if index == -1 {
		return text, ""
	}

	return text[0:index], text[index+1 : length]
}

func strReplace(text string, replace string, value string) string {
	return strings.Replace(text, replace, value, 1)
}

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

func addProperties(p *Property) {
	if properties == nil {
		properties = []*Property{p}
	} else {
		properties = append(properties, p)
	}
}

func isValid() bool {
	return errMsg == nil
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

func propertyNamesToJson() {
	if errMsg == nil {
		return
	}

	for _, p := range properties {
		if errMsg[p.Name] == nil || p.NameJson == "" {
			continue
		}

		errMsg[p.NameJson] = errMsg[p.Name]
		delete(errMsg, p.Name)
	}
}

func getProperty(name string) (*Property, error) {
	for _, p := range properties {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, errors.New("No property " + name + "found")
}

func getPropertyJsonName(p *Property) string {
	if p.NameJson == "" {
		return p.Name
	}
	return p.NameJson
}
