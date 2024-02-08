# go-validator


Usage example
 

```go
package main

import (
	"fmt"

	validator "github.com/otyang/go-validator"
)

func main() {
	validator := validator.New(
        // Register custom validation tags here
		validator.WithCustomValidationAlphaSpace(),
	)

	type User struct {
		Name  string `validate:"required,min=4"`
		Email string `validate:"required,email"`
		Age   int    `validate:"required,gte=18"`
	}

	user := User{
		Name:  "John",
		Email: "john.doe@example.com",
		Age:   25,
	}

	err := validator.ValidateStruct(&user)
	if err != nil {
		// Handle validation errors
		fmt.Println(err)
	} else {
		// User data is valid
		fmt.Println("User data is valid")
	}

    // add custom validation errors
	validator.AddError("field", "error message")

	if !validator.Valid() {
		for field, message := range validator.Errors {
			fmt.Println("Error for field", field, ":", message)
		}
	}

	validator.Reset() // resetting errors before preceeding to another validation
}
```
 
See test for more on how to use.