package utils

import (
	"event_backend/models"
	"fmt"
	"regexp"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

// ValidationErrors ..
var ValidationErrors = []string{}

//ValidateUser => checks for validation errors
func ValidateUser(user models.User, err []string) []string {
	//Email check
	emailVal := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailVal != true {
		err = append(err, "Invalid email")
		fmt.Println(err)
	}

	//Username length check
	if len(user.Username) < 3 || len(user.Username) > 13 {
		err = append(err, "Invalid Username!Should be 3-13 chars")
	}

	//Passoword length check
	if len(user.Password) < 7 {
		err = append(err, "Invalid Password")
	}

	// Validate Gender
	if len(user.Gender) == 0 {
		fmt.Println("Missing Gender")
	}

	return err
}
