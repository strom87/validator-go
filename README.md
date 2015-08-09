# Validator-GO
Validates properties for structs.

```sh
$ go get github.com/strom87/validator-go
```

### Get started
~~~ go
import (
    "github.com/strom87/validator-go"
)

type Example struct {
    ValueOne     int     `validator:"between:3,8"`
    ValueTwo     float64 `validator:"min:5|max:10"`
    ValueThree   string  `validator:"regexp:ca[a-z]{2}e"`
    Password     string  `validator:"required|equals:ConfPassword"`
    ConfPassword string  
}

func main() {
    example := Example{8, 10.1, "calle", "password", "pazzw0rd"}
    if isValid, errorMessages, err := validator.Validate(&example); !isValid {
        // Handle errors
    }
}
~~~

### Types
    min
        Minimum length of a string or value of a number.
        
    max
        Maximum length of a string or value of a number.
        
    between
        Length of a string between two numbers or a numberic value between two numbers.
        
    required
        Checks that the length of a string is longer then zero or that a number is greater than zero.
        
    equals
        Validates that two properties have the same value.
        
    regexp
        Runs regexp agains a string, only work with string values.
    
### Multiple validation types
You can have multiple validation types on the same property, just separeate each propery with character | (pipe)   
    Example: `validator:"required|min:2|max:20"`
    
### Functions
~~~ go
func Validate(obj interface{}) (bool, map[string][]string, error) {
    // Returns errorMessages with the names of the properties of the struct
}

func ValidateJson(obj interface{}) (bool, map[string][]string, error) {
    // Returns errorMessages with the names of the json tags of the struct
}
~~~

### Json example
~~~go
import (
    "github.com/strom87/validator-go"
)

type Example struct {
    FirstName string `json:"first_name" validator:"min:5"`
    LastName  string `json:"last_name"  validator:"max:6"`
}

func main() {
    example := Example{"Bob", "Barkinzon"}
    if isValid, errorMessages, err := validator.ValidateJson(&example); !isValid {
        // Handle errors
    }
}
~~~
