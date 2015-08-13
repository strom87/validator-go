# Validator-GO
An easy and simple way to validate property values in strucs for golang.   
Can be used to validate incoming request data to ensure that all the user data is correct.

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
    ValueFour    string  `validator:"alpha"`
    Email        string  `validator:"email"`
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

### Only for strings
    regexp
        Runs regexp against strings.
        
    email
        Validates that the string is a correct email address.
    
    alpha
        Validates that the string only contains alphabetic characters.
    
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

func ValidateLang(obj interface{}, language string) (bool, map[string][]string, error) {
    // With language parameter to change language on error messages
}

func ValidateJsonLang(obj interface{}, language string) (bool, map[string][]string, error) {
    // Json with language parameter to change language on error messages
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

### Languages
In the folder "lang" all the translation files is held. If support for a new language is supported just add a new json language file in this map. The name of the file is the key to choose the language.  
Say we would like to add a language file and call it "my_language.json", we would just create this file in the lang folder.   
Then just coppy all the keys from the "en.json" language file and translate them.
##### Use new language file
~~~ go
if isValid, errorMessages, err := validator.ValidateLang(&structObj, "my_language"); !isValid {
    // Handle errors
}

// or

if isValid, errorMessages, err := validator.ValidateJsonLang(&structObj, "my_language"); !isValid {
    // Handle errors
}
~~~
