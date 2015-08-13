package validator

import (
	"reflect"
	"testing"
)

type validationFunc func(p *Property, value string) error

type minTestStruct struct {
	v1       int     `validator:"min:200"`
	v2       string  `validator:"min:4"`
	v3       float64 `validator:"min:10.5"`
	expected int
}

type maxTestStruct struct {
	v1       int     `validator:"max:10"`
	v2       string  `validator:"max:4"`
	v3       float64 `validator:"max:10.5"`
	expected int
}

type betweenTestStruct struct {
	v1       int     `validator:"between:15,20"`
	v2       string  `validator:"between:3,8"`
	v3       float64 `validator:"between:100,200"`
	expected int
}

type requiredTestStruct struct {
	v1       int     `validator:"required"`
	v2       string  `validator:"required"`
	v3       float64 `validator:"required"`
	expected int
}

type equalsTestStruct struct {
	i1       int `validator:"equals:i2"`
	i2       int
	s1       string `validator:"equals:s2"`
	s2       string
	f1       float64 `validator:"equals:f2"`
	f2       float64
	expected int
}

type regexpTestStruct struct {
	v1       string `validator:"regexp:[0-9]{2}.*[A-Za-z]{2}"`
	expected int
}

type emailTestStruct struct {
	v1       string `validator:"email"`
	expected int
}

type alphaTestStruct struct {
	v1       string `validator:"alpha"`
	expected int
}

func resetValuesForTest() {
	errMsg = nil
	properties = nil
}

func TestMinF(t *testing.T) {
	initialize("en")
	values := []minTestStruct{
		{199, "Bob", 10.49, 3},
		{300, "Jhonny", 1042.95, 0},
		{-2, "Brit", 10.5, 1},
	}

	for _, v := range values {
		testRunner(v, min)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func TestMaxF(t *testing.T) {
	initialize("en")
	values := []maxTestStruct{
		{11, "Jonny", 10.511, 3},
		{10, "Bobb", 10.5, 0},
		{-200, "abcdefg", 10.499, 1},
	}

	for _, v := range values {
		testRunner(v, max)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func TestBetweenF(t *testing.T) {
	initialize("en")
	values := []betweenTestStruct{
		{15, "Bob", 100.00, 3},
		{16, "Jhon", 100.00001, 0},
		{20, "Jhonny", 199.999, 1},
	}

	for _, v := range values {
		testRunner(v, between)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func TestRequiredF(t *testing.T) {
	initialize("en")
	values := []requiredTestStruct{
		{0, "", 0.0, 3},
		{1, "a", 0.00001, 0},
		{-1, "abc", -0.0001, 2},
	}

	for _, v := range values {
		testRunner(v, func(p *Property, value string) error {
			required(p)
			return nil
		})
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func TestEqualsF(t *testing.T) {
	initialize("en")
	values := []equalsTestStruct{
		{10, 11, "aBc12E", "AbC12e", 11.123, 11.124, 3},
		{10, 10, "AbC12e", "AbC12e", 11.123, 11.123, 0},
		{99, 99, "QWERTY", "qwerty", 0.0010, 0.0001, 2},
	}

	for _, v := range values {
		testRunner(v, equals)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func TestRegexpF(t *testing.T) {
	initialize("en")
	values := []regexpTestStruct{
		{"12&?$Ab", 0},
		{"12&!=a2", 1},
		{"12ab", 0},
		{"ab12", 1},
	}

	for _, v := range values {
		testRunner(v, regexp)
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func TestEmailF(t *testing.T) {
	initialize("en")
	values := []emailTestStruct{
		{"email.address@test.com", 0},
		{"email.addresstest.org", 1},
		{"email.address@test", 1},
		{"email.address21@some.stuff.nu", 0},
	}

	for _, v := range values {
		testRunner(v, func(p *Property, value string) error {
			return email(p)
		})
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func TestAlphaF(t *testing.T) {
	initialize("en")
	values := []alphaTestStruct{
		{"abcdefgåäö", 0},
		{"XYZPÅlkfäs", 0},
		{"1234567890", 1},
		{"ab2sqwerty", 1},
		{"ab#sq&er!t", 1},
	}

	for _, v := range values {
		testRunner(v, func(p *Property, value string) error {
			return alpha(p)
		})
		if len(errMsg) != v.expected {
			t.Error("Expected", v.expected, "got", len(errMsg))
		}
		resetValuesForTest()
	}
}

func testRunner(obj interface{}, fn validationFunc) {
	data := reflect.ValueOf(obj)
	for i := 0; i < data.NumField(); i++ {
		addProperties(NewProperty(data.Field(i), data.Type().Field(i)))
	}

	for _, p := range properties {
		for _, v := range p.Validations {
			fn(p, v.Value)
		}
	}
}
